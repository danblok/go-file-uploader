package upload

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type Uploader struct {
	host string
	url  string
}

func New(host string) *Uploader {
	return &Uploader{
		host: host,
		url:  host + "/upload",
	}
}

func (u *Uploader) UploadFiles(filenames []string) error {
	fmt.Println("Sending files...")
	return u.uploadFiles(filenames)
}

func (u *Uploader) uploadFiles(filenames []string) error {
	body := &bytes.Buffer{}
	mpw := multipart.NewWriter(body)

	err := u.attachFiles(mpw, filenames)
	if err != nil {
		return err
	}

	err = u.sendRequest(mpw, body)
	if err != nil {
		return err
	}

	return nil
}

// Attaches files and ensures writer is closed when done or an error occured
func (u *Uploader) attachFiles(mpw *multipart.Writer, filenames []string) error {
	defer mpw.Close()
	var err error
	for i := range filenames {
		err = u.attachFile(mpw, filenames[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// Attaches a file to the multipart body request
func (u *Uploader) attachFile(mpw *multipart.Writer, filename string) error {
	fw, err := mpw.CreateFormFile("files", filename)
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}

	return nil
}

// Dispatches a request with attached files
func (u *Uploader) sendRequest(mpw *multipart.Writer, body *bytes.Buffer) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodPost, u.url, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", mpw.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with response code: %d", res.StatusCode)
	}
	return nil
}
