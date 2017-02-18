package resource

import (
	"errors"
	"net/http"

	"github.com/manyminds/api2go"

	"github.com/edudev/go-omx/backend/storage"
)

// RendererResource is a resource management struct for Renderers
type RendererResource struct {
	RendererStorage *storage.RendererStorage
}

// FindAll takes an HTTP request and returns all Renderers in a response
func (s RendererResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	renderers := s.RendererStorage.GetAll()
	return &Response{Res: renderers}, nil
}

// FindOne takes an HTTP request and returns a single Renderer in a response
func (s RendererResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	renderer, err := s.RendererStorage.GetOne(ID)
	return &Response{Res: renderer}, err
}

// Create is not implemented for renderers
func (s RendererResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}

// Delete is not implemented for renderers
func (s RendererResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}

// Update is not implemented for renderers
func (s RendererResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}
