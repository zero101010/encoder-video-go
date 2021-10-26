package services

import (
	"context"
	"encoder/application/repositories"
	"encoder/domain"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/storage"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

// Baixa vídeo do Storage
func (v *VideoService) Download(bucketName string) error {
	ctx := context.Background()
	// Cria client google cloud storage
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	bkt := client.Bucket(bucketName)
	obj := bkt.Object(v.Video.FilePath)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}
	defer r.Close()
	// Le video baixado
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	// Cria arquivo Mp4
	f, err := os.Create(os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}
	// Escreve o valor baixado no arquivo mp4
	_, err = f.Write(body)
	if err != nil {
		return err
	}
	defer f.Close()
	log.Printf("video %v has been stored", v.Video.ID)

	return nil

}

// Fragmenta video usando o bento4 para fragmentar antes de aplicar a quebra para mpeg-dash

func (v *VideoService) Fragment() error {

	err := os.Mkdir(os.Getenv("localStoragePath")+"/"+v.Video.ID, os.ModePerm)
	if err != nil {
		return err
	}

	source := os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4"
	target := os.Getenv("localStoragePath") + "/" + v.Video.ID + ".frag"

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	fmt.Println("hahahah------")
	fmt.Println(output)
	if err != nil {
		return err
	}

	printOutput(output)

	return nil
}

// func (v *VideoService) Encode() error {
// 	cmdArgs := []strings{}
// 	cmdArgs = append(cmdArgs, os.Getenv("localStoragePath")+"/"+v.Video.ID+"frag")
// 	cmdArgs = append(cmdArgs, "--use-segment-timeline")
// 	cmdArgs = append(cmdArgs, "-o")
// 	cmdArgs = append(cmdArgs, os.Getenv("localStoragePath")+"/"+v.Video.ID)
// 	cmdArgs = append(cmdArgs, "-f")
// 	cmdArgs = append(cmdArgs, "-exec-fir")
// 	cmdArgs = append(cmdArgs, "/opt/bento4/bin")
// 	cmd := exec.Command("mp4dash", cmdArgs...)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return err
// 	}
// 	printOutput(output)
// 	return nil

// }

// Verifica o output do comando mp4fragment
func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("======> Output> %s\n", string(out))
	}
}
