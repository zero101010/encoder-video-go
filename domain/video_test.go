package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

// testes com go
func TestValidateIfIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIdIsNotUuid(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "ABC"
	video.ResourceID = "A"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "teste"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)

}
