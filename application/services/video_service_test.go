package services_test

import (
	"encoder/application/repositories"
	"encoder/application/services"
	"encoder/domain"
	"encoder/framework/database"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
func prepare() (*domain.Video, repositories.VideoRepository) {
	db := database.NewDbTest()
	defer db.Close()
	// Criando Objeto Vídeo e instanciando valores
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "danca.mp4"
	video.CreatedAt = time.Now()
	// seto qual banco de daods eu devo utilizar no meu repository
	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {

	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("video-go")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	err = videoService.Finish()
	require.Nil(t, err)
}
