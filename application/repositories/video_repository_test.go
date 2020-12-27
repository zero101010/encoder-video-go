package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewVideoRepositoryInsert(t *testing.T) {
	// Definindo o banco que vou usar
	db := database.NewDbTest()
	defer db.Close()
	// Criando Objeto Vídeo e instanciando valores
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	// seto qual banco de daods eu devo utilizar no meu repository
	repo := repositories.VideoRepositoryDb{Db: db}
	// insiro o vídeo no banco de acordo com a implementação do Insert
	repo.Insert(video)
	// dÁ UM FIND no banco buscando pelo valor que foi inserido para ver se foi criado realmente
	v, err := repo.Find(video.ID)
	// verifica se o valor do find id não está vazio
	require.NotEmpty(t, v.ID)
	//verifica se o retorno do err no find é Nil
	require.Nil(t, err)
	// Verifica se o o find e o create possuem o mesmo valor de ID
	require.Equal(t, v.ID, video.ID)

}
