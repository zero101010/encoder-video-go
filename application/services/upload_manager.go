package services

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"cloud.google.com/go/storage"
)

type VideoUpload struct {
	Paths        []string
	VideoPath    string
	OutputBucket string
	Errors       []string
}

// Construtor
func NewVideoUpload() *VideoUpload {
	return &VideoUpload{}
}

// fará o upload de um vídeo para o storage
func (vu *VideoUpload) UploadObject(objectPath string, client *storage.Client, ctx context.Context) error {
	// Pego o path do vídeo que será feito o upload
	path := strings.Split(objectPath, os.Getenv("localStoragePath")+"/")
	f, err := os.Open(objectPath)
	if err != nil {
		return err
	}
	defer f.Close()
	// Escrevo o arquivo no bucket
	wc := client.Bucket(vu.OutputBucket).Object(path[1]).NewWriter(ctx)
	// Permissão de leitura para todos os usuários que tem uma role
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	if _, err = io.Copy(wc, f); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

// Pegar a lista de vídeos que serão feito o uploads
func (vu *VideoUpload) loadPaths() error {
	// anda pelos arquivos de um diretório e adicionar em vu.paths
	err := filepath.Walk(vu.VideoPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			vu.Paths = append(vu.Paths, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (vu *VideoUpload) ProcessUpload(concurrency int, doneUpload chan string) error {
	// Crio um ponto de entrada dos nossos canais, possibilitando que possa existir uma comunicacao entre as go routines
	in := make(chan int, runtime.NumCPU())
	returnChannel := make(chan string)
	err := vu.loadPaths()
	if err != nil {
		return err
	}
	uploadClient, ctx, err := getClientUpload()
	if err != nil {
		return err
	}
	// cria go routines de acordo com a quantidade de concurrency para rodar em background
	for process := 0; process < concurrency; process++ {
		go vu.uploadWorker(in, returnChannel, uploadClient, ctx)
	}
	go func() {
		for x := 0; x < len(vu.Paths); x++ {
			in <- x
		}
		close(in)
	}()

	for r := range returnChannel {
		if r != "" {
			doneUpload <- r
			break
		}
	}
	return nil
}

func (vu *VideoUpload) uploadWorker(in chan int, returnChannel chan string, uploadClient *storage.Client, ctx context.Context) {
	for x := range in {
		err := vu.UploadObject(vu.Paths[x], uploadClient, ctx)
		if err != nil {
			vu.Errors = append(vu.Errors, vu.Paths[x])
			log.Printf("error during the upload: %v. Error: %v", vu.Paths[x], err)
			returnChannel <- err.Error()
		}
		returnChannel <- ""
	}
	returnChannel <- "Uploaded completed"
}

// Pegar o client do bucket
func getClientUpload() (*storage.Client, context.Context, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}
	return client, ctx, nil
}
