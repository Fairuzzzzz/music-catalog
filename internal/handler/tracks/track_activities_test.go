package tracks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fairuzzzzz/music-catalog/internal/models/trackactivities"
	"github.com/Fairuzzzzz/music-catalog/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_UpsertTrackActivities(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	isLikedTrue := true
	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().UpsertTrackActivites(gomock.Any(), uint(1), trackactivities.TrackActivitiesRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &isLikedTrue,
				}).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().UpsertTrackActivites(gomock.Any(), uint(1), trackactivities.TrackActivitiesRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &isLikedTrue,
				}).Return(assert.AnError)
			},
			expectedStatusCode: http.StatusBadRequest,
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
			endpoint := "/tracks/track-activity"

			payload := trackactivities.TrackActivitiesRequest{
				SpotifyID: "spotifyID",
				IsLiked:   &isLikedTrue,
			}
			payloadBytes, err := json.Marshal(payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, endpoint, io.NopCloser(bytes.NewBuffer(payloadBytes)))
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			// Token JWT Authorization
			token, err := jwt.CreateToken(1, "", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			// Membuat ServeHTTP
			h.ServeHTTP(w, req)

			// Pengecekan status code
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
