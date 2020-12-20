package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// o ponteiro em vídeo é para mudar o dado também na entidade vídeo
type Job struct {
	ID               string    `valid:"uuid"`
	OutputBucketPath string    `valid:"notnull"`
	Status           string    `valid:"notnull"`
	Video            *Video    `valid:"-"`
	VideoID          string    `valid:"-"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `valid:"-"`
	UpdatedAt        time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// cria por padrão os dados de ID
func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{OutputBucketPath: output,
		Status: status,
		Video:  video,
	}
	job.prepare()

	err := job.Validate()
	if err != nil {
		return nil, err
	}
	return &job, nil
}

// só pode ser chamado dentro do pacote domain por estar com letra minúscula
// Cria por padrão os valores de ID e CreatedAt
func (job *Job) prepare() {
	job.ID = uuid.NewV4().String()
	job.CreatedAt = time.Now()

}

func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)
	if err != nil {
		return err
	}
	return nil
}
