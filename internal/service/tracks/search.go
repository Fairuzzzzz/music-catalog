package tracks

import (
	"context"

	"github.com/Fairuzzzzz/music-catalog/internal/models/spotify"
	"github.com/Fairuzzzzz/music-catalog/internal/models/trackactivities"
	spotifyRepo "github.com/Fairuzzzzz/music-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error search track to spotify")
		return nil, err
	}

	// Mendapatkan apakah track tersebut di like atau tidak
	trackIDs := make([]string, len(trackDetails.Tracks.Items))
	for idx, item := range trackDetails.Tracks.Items {
		trackIDs[idx] = item.Id
	}

	trackActivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from database")
		return nil, err
	}

	return modelToResponse(trackDetails, trackActivities), nil
}

func modelToResponse(data *spotifyRepo.SpotifySearchResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {

		artistName := make([]string, len(item.Artists))
		for idx, artist := range item.Artists {
			artistName[idx] = artist.Name
		}

		imagesUrl := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imagesUrl[idx] = image.Url
		}

		items = append(items, spotify.SpotifyTrackObject{
			// album related fields
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesUrl:   imagesUrl,
			AlbumName:        item.Album.Name,
			AlbumReleaseDate: item.Album.ReleaseDate,

			// artists related fields
			ArtistsName: artistName,

			// track related fields
			Explicit: item.Explicit,
			Id:       item.Id,
			Name:     item.Name,
			IsLiked:  mapTrackActivities[item.Id].IsLiked,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Items:  items,
		Total:  data.Tracks.Total,
	}
}
