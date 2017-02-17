package resource

import (
    "errors"
    "net/http"

	"github.com/manyminds/api2go"

	"github.com/edudev/go-omx/backend/storage"
)

type SourceResource struct {
	SourceStorage *storage.SourceStorage
}

func (s SourceResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	sources := s.SourceStorage.GetAll()
	return &Response{Res: sources}, nil
}

func (s SourceResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	source, err := s.SourceStorage.GetOne(ID)
	return &Response{Res: source}, err
}

func (s SourceResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
    return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}

func (s SourceResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
    return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}

func (s SourceResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
    return &Response{}, api2go.NewHTTPError(errors.New("Not implemented"), "Not implemented", http.StatusBadRequest)
}
