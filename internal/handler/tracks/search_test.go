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

func TestHandler_Search(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       spotify.SearchResponse
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "succes",
			expectedStatusCode: 200,
			expectedBody: spotify.SearchResponse{
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
					},
				},
				Total: 907,
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "bohemian rhapsody", 10, 1, uint(1)).Return(&spotify.SearchResponse{
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
						},
					},
					Total: 907,
				}, nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: 400,
			expectedBody:       spotify.SearchResponse{},
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "bohemian rhapsody", 10, 1, uint(1)).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			// endpoint untuk api nya
			endpoint := `/tracks/search?query=bohemian+rhapsody&pageSize=10&pageIndex=1`

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

				response := spotify.SearchResponse{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
