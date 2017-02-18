package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"

	"github.com/edudev/go-omx/backend/model"
	"github.com/edudev/go-omx/backend/resource"
	"github.com/edudev/go-omx/backend/storage"
)

func main() {
	port := 8000
	api := api2go.NewAPI("v1")

    if len(os.Args) < 3 {
        fmt.Println("Usage: ./backend <media_files_dir> <media_files_url>")
        return
    }
	sourceStorage := storage.NewSourceStorage(os.Args[1], os.Args[2])
	rendererStorage := storage.NewRendererStorage()
	attachedSourceStorage := storage.NewAttachedSourceStorage()

	api.AddResource(model.Source{}, resource.SourceResource{SourceStorage: sourceStorage})
	api.AddResource(model.Renderer{}, resource.RendererResource{RendererStorage: rendererStorage})
	api.AddResource(model.AttachedSource{}, resource.AttachedSourceResource{
		SourceStorage:         sourceStorage,
		RendererStorage:       rendererStorage,
		AttachedSourceStorage: attachedSourceStorage})

	fmt.Printf("Listening on :%d", port)
	handler := api.Handler().(*httprouter.Router)

	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
