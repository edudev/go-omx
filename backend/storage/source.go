package storage

import (
	"fmt"
	"sort"

	"os"
	"mime"
	"strings"
	"path/filepath"
	"net/url"

	"github.com/edudev/go-omx/backend/model"
)

type SourceByID []*model.Source

func (c SourceByID) Len() int {
	return len(c)
}

func (c SourceByID) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c SourceByID) Less(i, j int) bool {
	return c[i].GetID() < c[j].GetID()
}

func NewSourceStorage(rootDir, baseURL string) *SourceStorage {
	return &SourceStorage{rootDir: rootDir, baseURL: baseURL}
}

type SourceStorage struct {
	rootDir string
	baseURL string
}

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
				result = append(result, &model.Source{Uri:uri})
		}
		return nil
	})

	sort.Sort(SourceByID(result))
	return result
}

func (s SourceStorage) GetOne(id string) (*model.Source, error) {
	sources := s.GetAll()

	for _, source := range sources {
		if source.GetID() == id {
			return source, nil
		}
	}

	return &model.Source{}, fmt.Errorf("Source for id %s not found", id)
}
