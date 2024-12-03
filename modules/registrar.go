package modules

import (
	"github.com/forbole/bdjuno/v4/modules/actions"
	"github.com/forbole/bdjuno/v4/modules/addresses"
	"github.com/forbole/bdjuno/v4/modules/marginacc"
	"github.com/forbole/bdjuno/v4/modules/marginaccwithdraw"
	"github.com/forbole/bdjuno/v4/modules/markets"
	"github.com/forbole/bdjuno/v4/modules/orders"
	"github.com/forbole/bdjuno/v4/modules/positions"
	"github.com/forbole/bdjuno/v4/modules/types"

	"github.com/forbole/juno/v5/modules/pruning"
	"github.com/forbole/juno/v5/modules/telemetry"

	"github.com/forbole/bdjuno/v4/modules/slashing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	jmodules "github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/modules/messages"
	"github.com/forbole/juno/v5/modules/registrar"

	"github.com/forbole/bdjuno/v4/utils"

	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/bdjuno/v4/modules/assetft"
	"github.com/forbole/bdjuno/v4/modules/assetnft"
	"github.com/forbole/bdjuno/v4/modules/auth"
	"github.com/forbole/bdjuno/v4/modules/bank"
	"github.com/forbole/bdjuno/v4/modules/consensus"
	"github.com/forbole/bdjuno/v4/modules/customparams"
	dailyrefetch "github.com/forbole/bdjuno/v4/modules/daily_refetch"
	"github.com/forbole/bdjuno/v4/modules/distribution"
	"github.com/forbole/bdjuno/v4/modules/feegrant"
	"github.com/forbole/bdjuno/v4/modules/feemodel"
	"github.com/forbole/bdjuno/v4/modules/gov"
	"github.com/forbole/bdjuno/v4/modules/mint"
	"github.com/forbole/bdjuno/v4/modules/modules"
	"github.com/forbole/bdjuno/v4/modules/pricefeed"
	"github.com/forbole/bdjuno/v4/modules/staking"
	"github.com/forbole/bdjuno/v4/modules/upgrade"
)

// UniqueAddressesParser returns a wrapper around the given parser that removes all duplicated addresses
func UniqueAddressesParser(parser messages.MessageAddressesParser) messages.MessageAddressesParser {
	return func(cdc codec.Codec, msg sdk.Msg) ([]string, error) {
		addresses, err := parser(cdc, msg)
		if err != nil {
			return nil, err
		}

		return utils.RemoveDuplicateValues(addresses), nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ registrar.Registrar = &Registrar{}
)

// Registrar represents the modules.Registrar that allows to register all modules that are supported by BigDipper
type Registrar struct {
	parser messages.MessageAddressesParser
}

// NewRegistrar allows to build a new Registrar instance
func NewRegistrar(parser messages.MessageAddressesParser) *Registrar {
	return &Registrar{
		parser: UniqueAddressesParser(parser),
	}
}

// BuildModules implements modules.Registrar
func (r *Registrar) BuildModules(ctx registrar.Context) jmodules.Modules {
	cdc := ctx.EncodingConfig.Codec
	db := database.Cast(ctx.Database)

	sources, err := types.BuildSources(ctx.JunoConfig.Node, ctx.EncodingConfig)
	if err != nil {
		panic(err)
	}

	actionsModule := actions.NewModule(ctx.JunoConfig, ctx.EncodingConfig)
	authModule := auth.NewModule(sources.AuthSource, r.parser, cdc, db)
	bankModule := bank.NewModule(r.parser, sources.BankSource, cdc, db, ctx.JunoConfig.Chain.Bech32Prefix)
	consensusModule := consensus.NewModule(db)
	dailyRefetchModule := dailyrefetch.NewModule(ctx.Proxy, db)
	distrModule := distribution.NewModule(sources.DistrSource, cdc, db)
	feegrantModule := feegrant.NewModule(cdc, db)
	mintModule := mint.NewModule(sources.MintSource, cdc, db)
	slashingModule := slashing.NewModule(sources.SlashingSource, cdc, db)
	stakingModule := staking.NewModule(sources.StakingSource, cdc, db)
	feeModelModule := feemodel.NewModule(sources.FeeModelSource, cdc, db)
	customParamsModule := customparams.NewModule(sources.CustomParamsSource, cdc, db)
	assetFTModule := assetft.NewModule(sources.AssetFTSource, cdc, db)
	assetNFTModule := assetnft.NewModule(sources.AssetNFTSource, cdc, db)

	govModule := gov.NewModule(
		sources.GovSource,
		authModule,
		distrModule,
		mintModule,
		slashingModule,
		stakingModule,
		feeModelModule,
		customParamsModule,
		assetFTModule,
		assetNFTModule,
		cdc,
		db,
	)
	upgradeModule := upgrade.NewModule(db, stakingModule)

	marketsModule := markets.NewModule(cdc, db)
	marginaccModule := marginacc.NewModule(cdc, db)
	marginaccwithdrawModule := marginaccwithdraw.NewModule(cdc, db)

	ordersModule := orders.NewModule(cdc, db)
	positionsModule := positions.NewModule(cdc, db)

	return []jmodules.Module{
		messages.NewModule(r.parser, cdc, ctx.Database),
		telemetry.NewModule(ctx.JunoConfig),
		pruning.NewModule(ctx.JunoConfig, db, ctx.Logger),

		actionsModule,
		authModule,
		bankModule,
		consensusModule,
		dailyRefetchModule,
		distrModule,
		feegrantModule,
		govModule,
		mintModule,
		modules.NewModule(ctx.JunoConfig.Chain, db),
		pricefeed.NewModule(ctx.JunoConfig, cdc, db),
		slashingModule,
		stakingModule,
		upgradeModule,
		feeModelModule,
		customParamsModule,
		assetFTModule,
		assetNFTModule,
		marginaccModule,
		marginaccwithdrawModule,
		marketsModule,
		ordersModule,
		positionsModule,
		// This must be the last item.
		addresses.NewModule(r.parser, cdc, db),
	}
}
