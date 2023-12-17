package app

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const MAX_MEMORY = 32 << 20 // 32 Mb

type App struct {
	mux      *http.ServeMux
	filesDir string
	port     int
}

func New(port int, filesDir string) *App {
	app := &App{}
	app.setupMux()
	app.port = port
	app.filesDir = filesDir

	return app
}

func (a *App) setupMux() {
	a.mux = http.NewServeMux()
	a.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("MethodNotAllowed")
			http.Error(w, "MethodNotAllowed", http.StatusMethodNotAllowed)
			return
		}

		page, err := os.ReadFile("./web/index.html")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(page)
	})

	fs := http.FileServer(http.Dir(a.filesDir))
	a.mux.Handle("/files/", http.StripPrefix("/files", fs))

	a.mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("MethodNotAllowed")
			http.Error(w, "MethodNotAllowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseMultipartForm(MAX_MEMORY)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, fh := range r.MultipartForm.File["files"] {
			err = a.saveFile(fh)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(200)
	})
}

func (a *App) saveFile(fh *multipart.FileHeader) error {
	file, err := fh.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	newFile, err := os.Create(a.filesDir + "/" + fh.Filename)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.CopyN(newFile, file, fh.Size)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.mux)
}
