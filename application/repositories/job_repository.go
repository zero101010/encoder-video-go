package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
)

// INSERINDO interface
type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(id string) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

// Inserir algo no banco de dados do Vídeo
func (repo JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Create(job).Error

	if err != nil {
		return nil, err
	}
	return job, nil
}

func (repo JobRepositoryDb) Find(id string) (*domain.Job, error) {
	// Crio o objeto que vai recever os dados do Find
	var job domain.Job
	// faço a query e salvo por referência
	repo.Db.Preload("Video").First(&job, "id =?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("job doesn't find")
	}
	return &job, nil
}

func (repo JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Save(&job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}
