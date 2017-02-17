package storage

import (
	"fmt"
	"sort"

	"os"
	"mime"
	"strings"
	"path/filepath"

	"github.com/edudev/go-omx/backend/model"
)

// sorting
type byID []model.Source

func (c byID) Len() int {
	return len(c)
}

func (c byID) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c byID) Less(i, j int) bool {
	return c[i].GetID() < c[j].GetID()
}

func NewSourceStorage(rootDir string) *SourceStorage {
	return &SourceStorage{rootDir: rootDir}
}

type SourceStorage struct {
	rootDir string
}

func (s SourceStorage) GetAll() []model.Source {
	result := []model.Source{}
	filepath.Walk(s.rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		ext := filepath.Ext(path)
		mimeType := mime.TypeByExtension(ext)
		mimeTypeSplit := strings.Split(mimeType, "/")
		topLevelType := mimeTypeSplit[0]

		switch topLevelType {
		case
			"video",
			"audio":
			result = append(result, model.Source{Uri:path})
		}
		return nil
	})

	sort.Sort(byID(result))
	return result
}

func (s SourceStorage) GetOne(id string) (*model.Source, error) {
	sources := s.GetAll()

	for _, source := range sources {
		if source.GetID() == id {
			return &source, nil
		}
	}

	return &model.Source{}, fmt.Errorf("Source for id %s not found", id)
}
