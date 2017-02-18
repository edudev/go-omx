package storage

import (
	"fmt"
	"sort"

	"mime"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/edudev/go-omx/backend/model"
)

type sourceByID []*model.Source

func (c sourceByID) Len() int {
	return len(c)
}

func (c sourceByID) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c sourceByID) Less(i, j int) bool {
	return c[i].GetID() < c[j].GetID()
}

// NewSourceStorage creates a new SourceStorage instance
func NewSourceStorage(rootDir, baseURL string) *SourceStorage {
	return &SourceStorage{rootDir: rootDir, baseURL: baseURL}
}

// SourceStorage holds config data needed to list all the sources
type SourceStorage struct {
	rootDir string
	baseURL string
}

// GetAll returns all the sources as an array
func (s SourceStorage) GetAll() []*model.Source {
	result := []*model.Source{}
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
			key, err := filepath.Rel(s.rootDir, path)
			if err != nil {
				return nil
			}
			u, err := url.Parse(s.baseURL)
			if err != nil {
				return nil
			}
			u.Path = filepath.Join(u.Path, key)
			uri := u.String()
			result = append(result, &model.Source{URI: uri})
		}
		return nil
	})

	sort.Sort(sourceByID(result))
	return result
}

// GetOne returns a single source based on ID
func (s SourceStorage) GetOne(id string) (*model.Source, error) {
	sources := s.GetAll()

	for _, source := range sources {
		if source.GetID() == id {
			return source, nil
		}
	}

	return &model.Source{}, fmt.Errorf("Source for id %s not found", id)
}

// GetByURI returns a single source based on URI
func (s SourceStorage) GetByURI(uri string) *model.Source {
	sources := s.GetAll()

	for _, source := range sources {
		if source.URI == uri {
			return source
		}
	}

	return nil
}
