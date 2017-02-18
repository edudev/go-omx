package resource

import (
	"fmt"
	"math"
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
	as.RefreshAttachedSources()

	var result []*model.AttachedSource
	attachedSources := as.AttachedSourceStorage.GetAll()

	for _, attachedSource := range attachedSources {
		result = append(result, attachedSource)
	}

	return &Response{Res: result}, nil
}

func (as AttachedSourceResource) FindOne(IDs string, r api2go.Request) (api2go.Responder, error) {
	as.RefreshAttachedSources()

	ID, err := strconv.Atoi(IDs)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid id given"), "Invalid id given", http.StatusBadRequest)
	}

	attachedSource, err := as.AttachedSourceStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: attachedSource}, nil
}

func (as AttachedSourceResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	attachedSource, ok := obj.(model.AttachedSource)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	as.AttachedSourceStorage.RemoveRendererID(attachedSource.RendererID)

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

	id := as.AttachedSourceStorage.Insert(&attachedSource)
	attachedSource.ID = id

	err = renderer.Interface.StartPlayer(source.Uri)
	if err != nil {
		fmt.Println("Unable to start player: ", err)
	}
	return &Response{Res: attachedSource, Code: http.StatusCreated}, nil
}

func (as AttachedSourceResource) Delete(ids string, r api2go.Request) (api2go.Responder, error) {
	as.RefreshAttachedSources()

	id, err := strconv.Atoi(ids)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid id given"), "Invalid id given", http.StatusBadRequest)
	}
	err = as.AttachedSourceStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

func (as AttachedSourceResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {

	attachedSourceP, ok := obj.(*model.AttachedSource)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	attachedSource := *attachedSourceP

	as.RefreshAttachedSources()
	old, err := as.AttachedSourceStorage.Update(attachedSource)
	if old.PlaybackStatus != attachedSource.PlaybackStatus {
		if attachedSource.PlaybackStatus == "playing" {
			attachedSource.Renderer.Interface.Play()
		} else if attachedSource.PlaybackStatus == "paused" {
			attachedSource.Renderer.Interface.Pause()
		}
	}

	if math.Abs(old.PlaybackPosition-attachedSource.PlaybackPosition) > 2000.0 {
		attachedSource.Renderer.Interface.SetPosition(
			int64(attachedSource.PlaybackPosition * 1000))
	}

	return &Response{Res: attachedSource, Code: http.StatusNoContent}, err
}

func (as *AttachedSourceResource) RefreshAttachedSources() {
	for _, renderer := range as.RendererStorage.GetAll() {
		rid := renderer.GetID()
		if !renderer.Interface.HasPlayer() {
			//as.AttachedSourceStorage.RemoveRendererID(rid)
			continue
		}

		uri, err := renderer.Interface.Uri()
		if err != nil {
			as.AttachedSourceStorage.RemoveRendererID(rid)
			continue
		}

		source := as.SourceStorage.GetByUri(uri)
		if source == nil {
			continue
		}

		attachedSource := as.AttachedSourceStorage.GetByRenderSourceID(rid, source.GetID())
		if attachedSource == nil {
			attachedSource = &model.AttachedSource{
				SourceID:   source.GetID(),
				Source:     source,
				RendererID: rid,
				Renderer:   renderer,
			}

			asid := as.AttachedSourceStorage.Insert(attachedSource)
			attachedSource, err = as.AttachedSourceStorage.GetOne(asid)
		}

		asid, _ := strconv.Atoi(attachedSource.GetID())

		playbackStatus, err := renderer.Interface.PlaybackStatus()
		if err != nil {
			as.AttachedSourceStorage.Delete(asid)
			continue
		}
		attachedSource.PlaybackStatus = playbackStatus

		position, err := renderer.Interface.Position()
		if err != nil {
			as.AttachedSourceStorage.Delete(asid)
			continue
		}
		if position >= 0 {
			attachedSource.PlaybackPosition = float64(position) / 1000
		}

		volume, err := renderer.Interface.Volume()
		if err != nil {
			as.AttachedSourceStorage.Delete(asid)
			continue
		}
		attachedSource.Volume = volume
	}
}
