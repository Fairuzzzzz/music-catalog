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

func Test_service_Search(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSpotifyOutbound := NewMockspotifyOutbound(ctrlMock)

	mockTrackActivityRepo := NewMocktrackActivitesRepository(ctrlMock)

	next := "https://api.spotify.com/v1/search?query=bohemian+rhapsody&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.5&offset=20&limit=20"

	isLiked := true
	type args struct {
		query     string
		pageSize  int
		pageIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    *spotify.SearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:     "bohemian rhapsody",
				pageSize:  10,
				pageIndex: 1,
			},
			want: &spotify.SearchResponse{
				Limit:  20,
				Offset: 0,
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
				Total: 907,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(&spotifyRepo.SpotifySearchResponse{
					Tracks: spotifyRepo.SpotifyTracks{
						Href:   "https://api.spotify.com/v1/search?query=bohemian+rhapsody&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.5&offset=0&limit=20",
						Limit:  20,
						Next:   &next,
						Offset: 0,
						Total:  907,
						Items: []spotifyRepo.SpotifyTrackObject{
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
					},
				}, nil)

				mockTrackActivityRepo.EXPECT().GetBulkSpotifyIDs(gomock.Any(), uint(1), []string{"3z8h0TU7ReDPLIbEnYhWZb"}).Return(map[string]trackactivities.TrackActivity{
					"3z8h0TU7ReDPLIbEnYhWZb": {
						IsLiked: &isLiked,
					},
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				query:     "bohemian rhapsody",
				pageSize:  10,
				pageIndex: 1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				spotifyOutbound:     mockSpotifyOutbound,
				trackActivitiesRepo: mockTrackActivityRepo,
			}
			got, err := s.Search(context.Background(), tt.args.query, tt.args.pageSize, tt.args.pageIndex, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
