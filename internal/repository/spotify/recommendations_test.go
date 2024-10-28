package spotify

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/Fairuzzzzz/music-catalog/internal/configs"
	"github.com/Fairuzzzzz/music-catalog/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_outbound_GetRecommendation(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockHTTPClient := httpclient.NewMockHTTPClient(ctrlMock)
	type args struct {
		limit   int
		trackID string
	}
	tests := []struct {
		name    string
		args    args
		want    *SpotifyRecommendationResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				limit:   10,
				trackID: "trackID",
			},
			want: &SpotifyRecommendationResponse{
				Tracks: []SpotifyTrackObject{
					{
						Album: SpotifyAlbumObject{
							AlbumType:   "album",
							TotalTracks: 22,
							Images: []SpotifyAlbumImagesObject{
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
						Artists: []SpotifyArtistsObject{
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
			wantErr: false,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("market", "ID")
				params.Set("seed_tracks", args.trackID)

				basePath := `https://api.spotify.com/v1/recommendations`
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(recommendationResponse)),
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				limit:   10,
				trackID: "trackID",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("market", "ID")
				params.Set("seed_tracks", args.trackID)

				basePath := `https://api.spotify.com/v1/recommendations`
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				mockHTTPClient.EXPECT().Do(req).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed: 500",
			args: args{
				limit:   10,
				trackID: "trackID",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("market", "ID")
				params.Set("seed_tracks", args.trackID)

				basePath := `https://api.spotify.com/v1/recommendations`
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 500,
					Body:       io.NopCloser(bytes.NewBufferString(`Internal Server Error`)),
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			o := &outbound{
				cfg:         &configs.Config{},
				client:      mockHTTPClient,
				AccessToken: "accessToken",
				TokenType:   "Bearer",
				ExpiredAt:   time.Now().Add(1 * time.Hour),
			}
			got, err := o.GetRecommendation(context.Background(), tt.args.limit, tt.args.trackID)
			if (err != nil) != tt.wantErr {
				t.Errorf("outbound.GetRecommendation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outbound.GetRecommendation() = %v, want %v", got, tt.want)
			}
		})
	}
}