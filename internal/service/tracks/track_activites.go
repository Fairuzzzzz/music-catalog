package tracks

import (
	"context"
	"fmt"

	"github.com/Fairuzzzzz/music-catalog/internal/models/trackactivities"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *service) UpsertTrackActivites(ctx context.Context, userID uint, request trackactivities.TrackActivitiesRequest) error {
	activities, err := s.trackActivitiesRepo.Get(ctx, userID, request.SpotifyID)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error getting record from database")
		return err
	}

	if err == gorm.ErrRecordNotFound || activities == nil {
		// create user activity jika belum ada datanya
		err = s.trackActivitiesRepo.Create(ctx, trackactivities.TrackActivity{
			UserID:    userID,
			SpotifyID: request.SpotifyID,
			IsLiked:   request.IsLiked,
			CreatedBy: fmt.Sprintf("%d", userID),
			UpdatedBy: fmt.Sprintf("%d", userID),
		})
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Error().Err(err).Msg("error create record to databse")
			return err
		}
		return nil
	}
	// Jika sudah ada data nya, maka hanya update untuk isliked saja
	activities.IsLiked = request.IsLiked
	err = s.trackActivitiesRepo.Update(ctx, *activities)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error update record to databse")
		return err
	}
	return nil
}
