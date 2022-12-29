package remote

import (
	customparamstypes "github.com/CoreumFoundation/coreum/x/customparams/types"
	customparamssource "github.com/forbole/bdjuno/v3/modules/customparams/source"
	"github.com/forbole/juno/v3/node/remote"
)

var (
	_ customparamssource.Source = &Source{}
)

// Source implements customparamssource.Source using a remote node
type Source struct {
	*remote.Source
	customparamsClient customparamstypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, customparamsClient customparamstypes.QueryClient) *Source {
	return &Source{
		Source:             source,
		customparamsClient: customparamsClient,
	}
}

// GetParams implements customparamssource.Source
func (s Source) GetParams(height int64) (customparamstypes.StakingParams, error) {
	res, err := s.customparamsClient.StakingParams(remote.GetHeightRequestContext(s.Ctx, height), &customparamstypes.QueryStakingParamsRequest{})
	if err != nil {
		return customparamstypes.StakingParams{}, err
	}

	return res.Params, nil
}
