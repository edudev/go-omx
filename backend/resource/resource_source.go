package resource

import (
	"errors"
	"net/http"

	"github.com/manyminds/api2go"

	"github.com/edudev/go-omx/backend/storage"
)

// SourceResource is a resource management struct for Sources
type SourceResource struct {
	SourceStorage *storage.SourceStorage
}

// FindAll takes an HTTP request and returns all Sources in a response
func (s SourceResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	sources := s.SourceStorage.GetAll()
	return &Response{Res: sources}, nil
}

// FindOne takes an HTTP request and returns a single Source in a response
func (s SourceResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	source, err := s.SourceStorage.GetOne(ID)
	return &Response{Res: source}, err
}

// Create is not implemented for sources
func (s SourceResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}

// Delete is not implemented for sources
func (s SourceResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}

// Update is not implemented for sources
func (s SourceResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}
