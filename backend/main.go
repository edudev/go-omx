package main

import (
    "os"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"

	"github.com/edudev/go-omx/backend/model"
	"github.com/edudev/go-omx/backend/resource"
	"github.com/edudev/go-omx/backend/storage"
)

func main() {
	port := 8000
	api := api2go.NewAPI("v0")

	sourceStorage := storage.NewSourceStorage(os.Args[1])
	rendererStorage := storage.NewRendererStorage()

	api.AddResource(model.Source{}, resource.SourceResource{SourceStorage: sourceStorage})
	api.AddResource(model.Renderer{}, resource.RendererResource{RendererStorage: rendererStorage})

	fmt.Printf("Listening on :%d", port)
	handler := api.Handler().(*httprouter.Router)

	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
