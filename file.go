package bot

import (
	"io"
	"log"
	"net/http"
	"os"
)

type FilePath string

type FileID string

type InputFile interface{}

type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FilePath     string `json:"file_path"`
	FileSize     int    `json:"file_size"`
}

func (b *Bot) getFile(fileID string) (*File, error) {
	file := &File{}
	_, err := b.postJSON(MethodGetFile, map[string]string{
		"file_id": fileID,
	}, file)
	return file, err
}

func (b *Bot) DownloadFile(fileId string, pathname string) error {
	file, err := b.getFile(fileId)
	if err != nil {
		return err
	}

	url := b.getFileURL(file.FilePath)
	if b.Debug {
		log.Println("DownloadFile()", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	localFile, err := os.Create(pathname)
	if err != nil {
		return err
	}
	defer localFile.Close()

	_, err = io.Copy(localFile, resp.Body)
	return err
}
