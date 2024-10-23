package tracks

import (
	"context"

	"github.com/Fairuzzzzz/music-catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type spotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type service struct {
	spotifyOutbound spotifyOutbound
}

func NewService(spotifyOutbound spotifyOutbound) *service {
	return &service{
		spotifyOutbound: spotifyOutbound,
	}
}