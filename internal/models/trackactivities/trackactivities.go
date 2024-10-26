package trackactivities

import "gorm.io/gorm"

type (
	TrackActivity struct {
		gorm.Model
		UserID    uint   `gorm:"not null"`
		SpotifyID string `gorm:"not null"`
		IsLiked   *bool
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
	}
)

type (
	TrackActivitiesRequest struct {
		SpotifyID string `json:"spotifyID"`
		IsLiked   *bool  `json:"isLiked"` // true = like, false = dislike, null = neutral
	}
)
