package storage

import (
    "os"
	"fmt"
	"sort"

	"github.com/edudev/go-omx/omx"
	"github.com/edudev/go-omx/backend/model"
)

type RendererByID []*model.Renderer

func (c RendererByID) Len() int {
	return len(c)
}

func (c RendererByID) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c RendererByID) Less(i, j int) bool {
	return c[i].GetID() < c[j].GetID()
}

func NewRendererStorage() *RendererStorage {
	return &RendererStorage{}
}

type RendererStorage struct {
}

func (s RendererStorage) GetAll() []*model.Renderer {
	result := []*model.Renderer{}

    hostname, err := os.Hostname()
    if err == nil {
		inter, _ := omx.NewOmxInterface()
        result = append(result, &model.Renderer{
			Name: "omx",
			Host: hostname,
			Interface: inter})
    }

	sort.Sort(RendererByID(result))
	return result
}

func (s RendererStorage) GetOne(id string) (*model.Renderer, error) {
	renderers := s.GetAll()

	for _, renderer := range renderers {
		if renderer.GetID() == id {
			return renderer, nil
		}
	}

	return &model.Renderer{}, fmt.Errorf("Renderer for id %s not found", id)
}
