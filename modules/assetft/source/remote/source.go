package remote

import (
	assetftsource "github.com/forbole/bdjuno/v4/modules/assetft/source"
	"github.com/forbole/juno/v5/node/remote"

	assetfttypes "github.com/CoreumFoundation/coreum/v4/x/asset/ft/types"
)

var (
	_ assetftsource.Source = &Source{}
)

// Source implements assetftsource.Source using a remote node
type Source struct {
	*remote.Source
	assetftClient assetfttypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, assetftClient assetfttypes.QueryClient) *Source {
	return &Source{
		Source:        source,
		assetftClient: assetftClient,
	}
}

// GetParams implements assetftsource.Source
func (s Source) GetParams(height int64) (assetfttypes.Params, error) {
	res, err := s.assetftClient.Params(remote.GetHeightRequestContext(s.Ctx, height), &assetfttypes.QueryParamsRequest{})
	if err != nil {
		return assetfttypes.Params{}, err
	}

	return res.Params, nil
}
