package main

import (
	"fmt"
	"io"
	"os"
	
	"github.com/kkdai/youtube/v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <URL_DO_VIDEO>")
		return
	}
	
	videoURL := os.Args[1]
	client := youtube.Client{}
	
	fmt.Println("Buscando informações do vídeo...")
	
	video, err := client.GetVideo(videoURL)
	if err != nil {
		fmt.Println("Erro ao obter vídeo:", err)
		return
	}
	
	fmt.Printf("Baixando: %s\n", video.Title)
	
	formats := video.Formats.WithAudioChannels()
	if len(formats) == 0 {
		fmt.Println("Nenhum formato com áudio encontrado")
		return
	}
	
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		fmt.Println("Erro ao obter o stream:", err)
		return
	}
	
	defer stream.Close()
	
	file, err := os.Create("youtubego.mp4")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	
	defer file.Close()
	
	_, err = io.Copy(file, stream)
	if err != nil {
		fmt.Println("Erro ao salvar o vídeo:", err)
		return
	}
	
	fmt.Println("Download concluído com sucesso!")
}