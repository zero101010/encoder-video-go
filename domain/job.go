package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// init é um método padrão do go que executa o que está dentro da entidade antes de qualquer outro, nesse caso adicionamos um validator para os fields
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// o ponteiro em vídeo é para mudar o dado também na entidade vídeo
type Job struct {
	ID               string    `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OutputBucketPath string    `json:"output_bucket_path" valid:"notnull"`
	Status           string    `json:"status" valid:"notnull"`
	Video            *Video    `json:"video" valid:"-"`
	VideoID          string    `json:"-" valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `json:"created_at" valid:"-"`
	UpdatedAt        time.Time `json:"updated_at" valid:"-"`
}

// Cria um novo job com os dados do output, status e vídeo
func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
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
