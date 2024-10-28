package tracks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fairuzzzzz/music-catalog/internal/models/spotify"
	"github.com/Fairuzzzzz/music-catalog/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_GetRecommendation(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       *spotify.RecommendationResponse
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: 200,
			expectedBody: &spotify.RecommendationResponse{
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
					},
				},
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().GetRecommendation(gomock.Any(), uint(1), 10, "trackID").Return(&spotify.RecommendationResponse{
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
						},
					},
				}, nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       nil,
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().GetRecommendation(gomock.Any(), uint(1), 10, "trackID").Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn()
		api := gin.New()
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			// endpoint untuk api nya
			endpoint := `/tracks/recommendation?limit=10&trackID=trackID`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)

			// Token JWT Authorization
			token, err := jwt.CreateToken(1, "", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			// Membuat ServeHTTP
			h.ServeHTTP(w, req)

			// Pengecekan status code
			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := spotify.RecommendationResponse{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}
