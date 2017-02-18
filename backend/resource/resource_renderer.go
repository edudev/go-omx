package resource

import (
	"errors"
	"net/http"

	"github.com/manyminds/api2go"

	"github.com/edudev/go-omx/backend/storage"
)

type RendererResource struct {
	RendererStorage *storage.RendererStorage
}

func (s RendererResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	renderers := s.RendererStorage.GetAll()
	return &Response{Res: renderers}, nil
}

func (s RendererResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	renderer, err := s.RendererStorage.GetOne(ID)
	return &Response{Res: renderer}, err
}

func (s RendererResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}

func (s RendererResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}

func (s RendererResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}
