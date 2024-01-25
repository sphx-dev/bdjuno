package remote

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authsource "github.com/forbole/bdjuno/v4/modules/auth/source"
	"github.com/forbole/juno/v5/node/remote"
)

var (
	_ authsource.Source = &Source{}
)

// Source implements authsource.Source using a remote node
type Source struct {
	*remote.Source
	authClient authtypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, authClient authtypes.QueryClient) *Source {
	return &Source{
		Source:     source,
		authClient: authClient,
	}
}

// GetParams implements authsource.Source
func (s Source) GetParams(height int64) (authtypes.Params, error) {
	res, err := s.authClient.Params(remote.GetHeightRequestContext(s.Ctx, height), &authtypes.QueryParamsRequest{})
	if err != nil {
		return authtypes.Params{}, err
	}

	return res.Params, nil
}
