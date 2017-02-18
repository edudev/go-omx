package resource

import (
	"strconv"

	"errors"
	"net/http"

	"github.com/manyminds/api2go"

	"github.com/edudev/go-omx/backend/model"
	"github.com/edudev/go-omx/backend/storage"
)

type AttachedSourceResource struct {
	AttachedSourceStorage *storage.AttachedSourceStorage
	SourceStorage         *storage.SourceStorage
	RendererStorage       *storage.RendererStorage
}

func (as AttachedSourceResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*model.AttachedSource
	attachedSources := as.AttachedSourceStorage.GetAll()

	for _, attachedSource := range attachedSources {
		source, err := as.SourceStorage.GetOne(attachedSource.SourceID)
		if err != nil {
			return &Response{}, err
		}

		renderer, err := as.RendererStorage.GetOne(attachedSource.RendererID)
		if err != nil {
			return &Response{}, err
		}

		attachedSource.Source = source
		attachedSource.Renderer = renderer

		result = append(result, attachedSource)
	}

	return &Response{Res: result}, nil
}

func (as AttachedSourceResource) FindOne(IDs string, r api2go.Request) (api2go.Responder, error) {
	ID, err := strconv.Atoi(IDs)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid id given"), "Invalid id given", http.StatusBadRequest)
	}

	attachedSource, err := as.AttachedSourceStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	source, err := as.SourceStorage.GetOne(attachedSource.SourceID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	renderer, err := as.RendererStorage.GetOne(attachedSource.RendererID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	attachedSource.Source = source
	attachedSource.Renderer = renderer

	return &Response{Res: attachedSource}, nil
}

func (as AttachedSourceResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	attachedSource, ok := obj.(model.AttachedSource)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	source, err := as.SourceStorage.GetOne(attachedSource.SourceID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	renderer, err := as.RendererStorage.GetOne(attachedSource.RendererID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	attachedSource.Source = source
	attachedSource.Renderer = renderer

	id := as.AttachedSourceStorage.Insert(attachedSource)
	attachedSource.ID = id

	return &Response{Res: attachedSource, Code: http.StatusCreated}, nil
}

func (as AttachedSourceResource) Delete(ids string, r api2go.Request) (api2go.Responder, error) {
	id, err := strconv.Atoi(ids)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid id given"), "Invalid id given", http.StatusBadRequest)
	}
	err = as.AttachedSourceStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

func (as AttachedSourceResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	attachedSource, ok := obj.(model.AttachedSource)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := as.AttachedSourceStorage.Update(attachedSource)
	return &Response{Res: attachedSource, Code: http.StatusNoContent}, err
}
