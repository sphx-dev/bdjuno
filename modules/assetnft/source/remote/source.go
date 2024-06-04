package remote

import (
	assetnftsource "github.com/forbole/bdjuno/v4/modules/assetnft/source"
	"github.com/forbole/juno/v5/node/remote"

	assetnfttypes "github.com/CoreumFoundation/coreum/v4/x/asset/nft/types"
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
