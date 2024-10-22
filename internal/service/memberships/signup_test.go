package memberships

import (
	"database/sql"
	"testing"

	"github.com/Fairuzzzzz/music-catalog/internal/configs"
	"github.com/Fairuzzzzz/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_service_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: memberships.SignUpRequest{
					Email:     "test@gmail.com",
					Username:  "testusername",
					Passwword: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, sql.ErrNoRows)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		{
			name: "failed when Get User",
			args: args{
				request: memberships.SignUpRequest{
					Email:     "test@gmail.com",
					Username:  "testusername",
					Passwword: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when Create User",
			args: args{
				request: memberships.SignUpRequest{
					Email:     "test@gmail.com",
					Username:  "testusername",
					Passwword: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, sql.ErrNoRows)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			if err := s.SignUp(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
