package storage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/edudev/go-omx/backend/model"
	"github.com/manyminds/api2go"
)

// NewAttachedSourceStorage creates a new AttachedSourceStorage instance
func NewAttachedSourceStorage() *AttachedSourceStorage {
	return &AttachedSourceStorage{make(map[int]*model.AttachedSource), 1}
}

// AttachedSourceStorage holds all the attached sources in memory
type AttachedSourceStorage struct {
	attachedSources map[int]*model.AttachedSource
	idCount         int
}

// GetAll returns a map of all the attached sources by ID
func (as AttachedSourceStorage) GetAll() map[int]*model.AttachedSource {
	return as.attachedSources
}

// GetOne returns a single attached source by ID
func (as AttachedSourceStorage) GetOne(id int) (*model.AttachedSource, error) {
	attachedSource, ok := as.attachedSources[id]
	if ok {
		return attachedSource, nil
	}
	errMessage := fmt.Sprintf("AttachedSource for id %d not found", id)
	return nil, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert adds an AttachedSource to the memory DB and returns it's ID
func (as *AttachedSourceStorage) Insert(attachedSource *model.AttachedSource) int {
	id := as.idCount
	attachedSource.ID = id
	as.attachedSources[id] = attachedSource
	as.idCount++
	return id
}

// Delete removes an AttachedSource from the memory DB
func (as *AttachedSourceStorage) Delete(id int) error {
	_, exists := as.attachedSources[id]
	if !exists {
		return fmt.Errorf("AttachedSource with id %d does not exist", id)
	}
	delete(as.attachedSources, id)

	return nil
}

// Update swaps an old AttachedSource with a new one, returning the old one
func (as *AttachedSourceStorage) Update(attachedSource model.AttachedSource) (*model.AttachedSource, error) {
	id := attachedSource.ID
	_, exists := as.attachedSources[id]
	if !exists {
		return nil, fmt.Errorf("AttachedSource with id %d does not exist", id)
	}
	old := as.attachedSources[id]
	as.attachedSources[id] = &attachedSource

	return old, nil
}

// RemoveRendererID removes all attached sources associated with a particular Render
func (as *AttachedSourceStorage) RemoveRendererID(rendererID string) {
	for attachedSourceID, attachedSource := range as.attachedSources {
		if attachedSource.RendererID == rendererID {
			delete(as.attachedSources, attachedSourceID)
		}
	}
}

// GetByRenderSourceID returns an attached source matching both Renderer and Source
func (as *AttachedSourceStorage) GetByRenderSourceID(rendererID, sourceID string) *model.AttachedSource {
	for _, attachedSource := range as.attachedSources {
		if attachedSource.RendererID == rendererID && attachedSource.SourceID == sourceID {
			return attachedSource
		}
	}

	return nil
}
