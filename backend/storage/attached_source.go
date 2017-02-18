package storage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/edudev/go-omx/backend/model"
	"github.com/manyminds/api2go"
)

func NewAttachedSourceStorage() *AttachedSourceStorage {
	return &AttachedSourceStorage{make(map[int]*model.AttachedSource), 1}
}

type AttachedSourceStorage struct {
	attachedSources map[int]*model.AttachedSource
	idCount         int
}

func (as AttachedSourceStorage) GetAll() map[int]*model.AttachedSource {
	return as.attachedSources
}

func (as AttachedSourceStorage) GetOne(id int) (model.AttachedSource, error) {
	attachedSource, ok := as.attachedSources[id]
	if ok {
		return *attachedSource, nil
	}
	errMessage := fmt.Sprintf("AttachedSource for id %d not found", id)
	return model.AttachedSource{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (as *AttachedSourceStorage) Insert(attachedSource model.AttachedSource) int {
	id := as.idCount
	attachedSource.ID = id
	as.attachedSources[id] = &attachedSource
	as.idCount++
	return id
}

func (as *AttachedSourceStorage) Delete(id int) error {
	_, exists := as.attachedSources[id]
	if !exists {
		return fmt.Errorf("AttachedSource with id %d does not exist", id)
	}
	delete(as.attachedSources, id)

	return nil
}

func (as *AttachedSourceStorage) Update(attachedSource model.AttachedSource) error {
	id := attachedSource.ID
	_, exists := as.attachedSources[id]
	if !exists {
		return fmt.Errorf("AttachedSource with id %d does not exist", id)
	}
	as.attachedSources[id] = &attachedSource

	return nil
}
