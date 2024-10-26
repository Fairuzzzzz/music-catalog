package main

import (
	"log"
	"net/http"

	"github.com/Fairuzzzzz/music-catalog/internal/configs"
	membershipsHandler "github.com/Fairuzzzzz/music-catalog/internal/handler/memberships"
	tracksHandler "github.com/Fairuzzzzz/music-catalog/internal/handler/tracks"
	"github.com/Fairuzzzzz/music-catalog/internal/models/memberships"
	"github.com/Fairuzzzzz/music-catalog/internal/models/trackactivities"
	membershipsRepo "github.com/Fairuzzzzz/music-catalog/internal/repository/memberships"
	"github.com/Fairuzzzzz/music-catalog/internal/repository/spotify"
	trackactivitiesRepo "github.com/Fairuzzzzz/music-catalog/internal/repository/trackactivities"
	membershipsSvc "github.com/Fairuzzzzz/music-catalog/internal/service/memberships"
	"github.com/Fairuzzzzz/music-catalog/internal/service/tracks"
	"github.com/Fairuzzzzz/music-catalog/pkg/httpclient"
	"github.com/Fairuzzzzz/music-catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisialisasi config", err)
	}

	cfg = configs.Get()
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database, err: %+v\n", err)
	}

	db.AutoMigrate(&memberships.User{})
	db.AutoMigrate(&trackactivities.TrackActivity{})

	httpClient := httpclient.NewClient(&http.Client{})

	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, httpClient)

	membershipRepo := membershipsRepo.NewRepository(db)

	trackActivitiesRepo := trackactivitiesRepo.NewRepository(db)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)

	tracksSvc := tracks.NewService(spotifyOutbound, trackActivitiesRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	trackHandler := tracksHandler.NewHandler(r, tracksSvc)
	trackHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
