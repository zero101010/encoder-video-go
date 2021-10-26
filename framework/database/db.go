package database

import (
	"encoder/domain"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Database struct {
	Db          *gorm.DB
	Dsn         string
	DsnTest     string
	DbType      string
	DbTypeTeste string
	Debug       bool
	AutoMigrate bool
	Env         string
}

func NewDb() *Database {
	return &Database{}
}

// seta os valores do objeto quando está rodando o banco de teste e cria a conecção
func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "Test"
	dbInstance.DbTypeTeste = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrate = true
	dbInstance.Debug = true
	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}
	return connection
}

// Cria a conexão com o banco de dados e gerencia o que vai ser criado de acordo com cada env
func (d *Database) Connect() (*gorm.DB, error) {
	var err error
	// Verifica se deve abrir a conexão com o banco de test ou em produção
	if d.Env != "Test" {
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTeste, d.DsnTest)
	}
	// Verifica se a conexão ocorreu da forma como esperamos
	if err != nil {
		return nil, err
	}
	// verifica se o modo debug está ativo ou não
	if d.Debug {
		d.Db.LogMode(true)
	}
	if d.AutoMigrate {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
		d.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos (id)", "CASCADE", "CASCADE")
	}
	return d.Db, nil
}
