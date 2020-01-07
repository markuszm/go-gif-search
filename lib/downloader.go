package lib

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
)

type Downloader struct {
	Folder string
}

func (d *Downloader) StoreFile(url string, fileName string) (string, error) {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "download failed")
	}
	body := resp.Body
	filePath := fmt.Sprintf("%s/%s.mp4", d.Folder, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", errors.Wrap(err, "cannot create new file")
	}
	defer file.Close()
	_, err = io.Copy(file, body)
	return filePath, err
}
