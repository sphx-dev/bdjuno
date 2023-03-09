package remote

import (
	assetnfttypes "github.com/CoreumFoundation/coreum/x/asset/nft/types"
	assetnftsource "github.com/forbole/bdjuno/v3/modules/assetnft/source"
	"github.com/forbole/juno/v3/node/remote"
)

var (
	_ assetnftsource.Source = &Source{}
)

// Source implements assetnftsource.Source using a remote node
type Source struct {
	*remote.Source
	assetnftClient assetnfttypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, assetnftClient assetnfttypes.QueryClient) *Source {
	return &Source{
		Source:         source,
		assetnftClient: assetnftClient,
	}
}

// GetParams implements assetnftsource.Source
func (s Source) GetParams(height int64) (assetnfttypes.Params, error) {
	res, err := s.assetnftClient.Params(remote.GetHeightRequestContext(s.Ctx, height), &assetnfttypes.QueryParamsRequest{})
	if err != nil {
		return assetnfttypes.Params{}, err
	}

	return res.Params, nil
}
