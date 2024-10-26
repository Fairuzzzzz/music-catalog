package memberships

import (
	"errors"

	"github.com/Fairuzzzzz/music-catalog/internal/models/memberships"
	"github.com/Fairuzzzzz/music-catalog/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(request memberships.LoginRequest) (string, error) {
	userDetail, err := s.repository.GetUser(request.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get user from database")
		return "", err
	}

	if userDetail == nil {
		errors.New("email not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("email and password not match")
	}

	accessToken, err := jwt.CreateToken(userDetail.ID, userDetail.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create jwt token")
		return "", err
	}

	return accessToken, nil
}
