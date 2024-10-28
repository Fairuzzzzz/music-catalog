package tracks

import (
	"context"
	"reflect"
	"testing"

	"github.com/Fairuzzzzz/music-catalog/internal/models/spotify"
	"github.com/Fairuzzzzz/music-catalog/internal/models/trackactivities"
	spotifyRepo "github.com/Fairuzzzzz/music-catalog/internal/repository/spotify"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_service_GetRecommendation(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSpotifyOutbound := NewMockspotifyOutbound(ctrlMock)

	mockTrackActivitiesRepo := NewMocktrackActivitesRepository(ctrlMock)

	isLiked := true

	type args struct {
		userID  uint
		limit   int
		trackID string
	}
	tests := []struct {
		name    string
		args    args
		want    *spotify.RecommendationResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				userID:  1,
				limit:   10,
				trackID: "trackID",
			},
			want: &spotify.RecommendationResponse{
				Items: []spotify.SpotifyTrackObject{
					{
						AlbumType:        "album",
						AlbumTotalTracks: 22,
						AlbumImagesUrl:   []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
						AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
						AlbumReleaseDate: "2018-10-19",

						ArtistsName: []string{"Queen"},

						Explicit: false,
						Id:       "3z8h0TU7ReDPLIbEnYhWZb",
						Name:     "Bohemian Rhapsody",
						IsLiked:  &isLiked,
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().GetRecommendation(gomock.Any(), 10, "trackID").Return(&spotifyRepo.SpotifyRecommendationResponse{
					Tracks: []spotifyRepo.SpotifyTrackObject{
						{
							Album: spotifyRepo.SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 22,
								Images: []spotifyRepo.SpotifyAlbumImagesObject{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
									},
								},
								Name:                 "Bohemian Rhapsody (The Original Soundtrack)",
								ReleaseDate:          "2018-10-19",
								ReleaseDatePrecision: "day",
							},
							Artists: []spotifyRepo.SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Id:   "1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
									Type: "artist",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
							Id:       "3z8h0TU7ReDPLIbEnYhWZb",
							Name:     "Bohemian Rhapsody",
						},
					},
				}, nil)

				mockTrackActivitiesRepo.EXPECT().GetBulkSpotifyIDs(gomock.Any(), uint(1), []string{"3z8h0TU7ReDPLIbEnYhWZb"}).Return(map[string]trackactivities.TrackActivity{
					"3z8h0TU7ReDPLIbEnYhWZb": {
						IsLiked: &isLiked,
					},
				}, nil)
			},
		},
		{
			name: "failed: went get bulk spotifyIDs",
			args: args{
				userID:  1,
				limit:   10,
				trackID: "trackID",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().GetRecommendation(gomock.Any(), 10, "trackID").Return(&spotifyRepo.SpotifyRecommendationResponse{
					Tracks: []spotifyRepo.SpotifyTrackObject{
						{
							Album: spotifyRepo.SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 22,
								Images: []spotifyRepo.SpotifyAlbumImagesObject{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
									},
								},
								Name:                 "Bohemian Rhapsody (The Original Soundtrack)",
								ReleaseDate:          "2018-10-19",
								ReleaseDatePrecision: "day",
							},
							Artists: []spotifyRepo.SpotifyArtistsObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Id:   "1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
									Type: "artist",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
							Id:       "3z8h0TU7ReDPLIbEnYhWZb",
							Name:     "Bohemian Rhapsody",
						},
					},
				}, nil)
				mockTrackActivitiesRepo.EXPECT().GetBulkSpotifyIDs(gomock.Any(), uint(1), []string{"3z8h0TU7ReDPLIbEnYhWZb"}).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed: went get recommendations",
			args: args{
				userID:  1,
				limit:   10,
				trackID: "trackID",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().GetRecommendation(gomock.Any(), 10, "trackID").Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				spotifyOutbound:     mockSpotifyOutbound,
				trackActivitiesRepo: mockTrackActivitiesRepo,
			}
			got, err := s.GetRecommendation(context.Background(), tt.args.userID, tt.args.limit, tt.args.trackID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetRecommendation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetRecommendation() = %v, want %v", got, tt.want)
			}
		})
	}
}
