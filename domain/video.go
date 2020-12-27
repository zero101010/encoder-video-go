package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Video definir atributos de vídeo
type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"`
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"`
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"-" valid:"-"`
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

// Verifica se os validadores de tipo, notnnull estão obrigatórios antes de criar o objeto
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

// Definir como método da Classe e o ponteiro é para mudar na estrutura principal
func (video *Video) Validate() error {
	_, error := govalidator.ValidateStruct(video)
	if error != nil {
		return error
	}
	return nil
}
