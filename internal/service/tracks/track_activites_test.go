package tracks

import (
	"context"
	"fmt"
	"testing"

	"github.com/Fairuzzzzz/music-catalog/internal/models/trackactivities"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_UpsertTrackActivites(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockTrackActivityRepo := NewMocktrackActivitesRepository(ctrlMock)

	isLiked := true
	isLikedFalse := false
	type args struct {
		userID  uint
		request trackactivities.TrackActivitiesRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success: create",
			args: args{
				userID: 1,
				request: trackactivities.TrackActivitiesRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &isLiked,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyID).Return(nil, gorm.ErrRecordNotFound)

				mockTrackActivityRepo.EXPECT().Create(gomock.Any(), trackactivities.TrackActivity{
					UserID:    args.userID,
					SpotifyID: args.request.SpotifyID,
					IsLiked:   args.request.IsLiked,
					CreatedBy: fmt.Sprintf("%d", args.userID),
					UpdatedBy: fmt.Sprintf("%d", args.userID),
				}).Return(nil)
			},
		},
		{
			name: "success: update",
			args: args{
				userID: 1,
				request: trackactivities.TrackActivitiesRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &isLiked,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyID).Return(&trackactivities.TrackActivity{
					IsLiked: &isLikedFalse,
				}, nil)

				mockTrackActivityRepo.EXPECT().Update(gomock.Any(), trackactivities.TrackActivity{
					IsLiked: args.request.IsLiked,
				}).Return(nil)
			},
		},
		{
			name: "failed",
			args: args{
				userID: 1,
				request: trackactivities.TrackActivitiesRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &isLiked,
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyID).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				trackActivitiesRepo: mockTrackActivityRepo,
			}
			if err := s.UpsertTrackActivites(context.Background(), tt.args.userID, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.UpsertTrackActivites() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
