package storage

import (
	"fmt"
	"os"
	"sort"

	"github.com/edudev/go-omx/backend/model"
	"github.com/edudev/go-omx/omx"
)

type rendererByID []*model.Renderer

func (c rendererByID) Len() int {
	return len(c)
}

func (c rendererByID) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c rendererByID) Less(i, j int) bool {
	return c[i].GetID() < c[j].GetID()
}

// NewRendererStorage creates a new RendererStorage instance
func NewRendererStorage() *RendererStorage {
	return &RendererStorage{}
}

// RendererStorage is a dummy struct, supposed to hold all the Renderers
type RendererStorage struct {
}

// GetAll returns all the renderers as an array
func (s RendererStorage) GetAll() []*model.Renderer {
	result := []*model.Renderer{}

	hostname, err := os.Hostname()
	if err == nil {
		inter, _ := omx.NewInterface()
		result = append(result, &model.Renderer{
			Name:      "omx",
			Host:      hostname,
			Interface: inter})
	}

	sort.Sort(rendererByID(result))
	return result
}

// GetOne returns a single renderer based on ID
func (s RendererStorage) GetOne(id string) (*model.Renderer, error) {
	renderers := s.GetAll()

	for _, renderer := range renderers {
		if renderer.GetID() == id {
			return renderer, nil
		}
	}

	return &model.Renderer{}, fmt.Errorf("Renderer for id %s not found", id)
}
