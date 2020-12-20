package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Video definir atributos de vídeo
type Video struct {
	ID         string    `valid:"uuid"`
	ResourceID string    `valid:"notnull"`
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
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
