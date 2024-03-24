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

func NewVideoUpload() *VideoUpload {
	return &VideoUpload{}
}

func (vu *VideoUpload) UploadObject(objectPath string, client *storage.Client, ctx context.Context) error {
	path := strings.Split(objectPath, os.Getenv("localStoragePath")+"/")

	f, err := os.Open(objectPath)
	if err != nil {
		return err
	}

	defer f.Close()

	wc := client.Bucket(vu.OutputBucket).Object(path[1]).NewWriter(ctx)
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	if _, err = io.Copy(wc, f); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

func (vu *VideoUpload) loadPaths() error {
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
	in := make(chan int, runtime.NumCPU()) // qual o arquivo baseado na posicao do slice Paths
	returnChannel := make(chan string)     // retorno do canal caso tenha gerado erro

	err := vu.loadPaths() // carrega todos os paths
	if err != nil {
		return err
	}

	uploadClient, ctx, err := getClientUpload() // instancia o client do upload
	if err != nil {
		return err
	}

	for process := 0; process < concurrency; process++ { //inicia o numero de workers desejado de acordo com a concurrency passada como parametro
		go vu.uploadWorker(in, returnChannel, uploadClient, ctx)
	}

	go func() { //goroutine que alimenta o canal in para que seja lido e esvaziado na funçao uploadWorker
		for x := 0; x < len(vu.Paths); x++ {
			in <- x
		}
		close(in)
	}()

	for r := range returnChannel { //for para caso retorne um erro em returnChan parar a execução das goroutines
		if r != "" {
			doneUpload <- r
			break
		}
	}

	return nil
}

func (vu *VideoUpload) uploadWorker(in chan int, returnChan chan string, uploadClient *storage.Client, ctx context.Context) {

	for x := range in { // for para ler o conteudo do canal in e esvaziar fazendo o upload por meio da funçao UploadObject, uma vez esvaziado o canal a goroutine da funçao acima enche novamente
		err := vu.UploadObject(vu.Paths[x], uploadClient, ctx)
		if err != nil { //caso UploadObject retorne erro é passado para o returnChan que executa um break nas goroutines
			vu.Errors = append(vu.Errors, vu.Paths[x])
			log.Printf("error during the upload: %v. Error: %v", vu.Paths[x], err)
			returnChan <- err.Error()
		}

		returnChan <- ""
	}
	returnChan <- "upload completed"
}

func getClientUpload() (*storage.Client, context.Context, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}

	return client, ctx, nil
}
