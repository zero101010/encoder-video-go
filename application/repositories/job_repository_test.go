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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()
	// Criando Objeto Vídeo e instanciando valores
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	// seto qual banco de dados eu devo utilizar no meu repository
	repo := repositories.VideoRepositoryDb{Db: db}
	// insiro o vídeo no banco de acordo com a implementação do Insert
	repo.Insert(video)
	// inicializa o objeto Job com os valores
	job, err := domain.NewJob("output_path", "Pending", video)
	// Verifico se não ocorreu nenhum erro nessa criação
	require.Nil(t, err)
	// Seto qual banco de dados usaremos
	repoJob := repositories.JobRepositoryDb{Db: db}
	// insiro o objeto job no banco
	repoJob.Insert(job)
	// Faço um find para garantir que o Job foi inserido
	j, err := repoJob.Find(job.ID)
	// Verifico os casos de teste
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()
	// Criando Objeto Vídeo e instanciando valores
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	// seto qual banco de dados eu devo utilizar no meu repository
	repo := repositories.VideoRepositoryDb{Db: db}
	// insiro o vídeo no banco de acordo com a implementação do Insert
	repo.Insert(video)
	// inicializa o objeto Job com os valores
	job, err := domain.NewJob("output_path", "Pending", video)
	// Verifico se não ocorreu nenhum erro nessa criação
	require.Nil(t, err)
	// Seto qual banco de dados usaremos
	repoJob := repositories.JobRepositoryDb{Db: db}
	// insiro o objeto job no banco
	repoJob.Insert(job)
	// Update Status
	job.Status = "Completed"
	repoJob.Update(job)
	// Faço um find para garantir que o Job foi inserido
	j, err := repoJob.Find(job.ID)
	// Verifico os casos de teste
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
