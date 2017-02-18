package model

import "testing"

func TestRendererHasID(t *testing.T) {
	r := Renderer{Name: "test", Host: "dog"}
	if r.GetID() != "" {
		return
	}
	t.Error("Unable to get Unique ID of Renderer")
}

func TestRenderersHaveUniqueID(t *testing.T) {
	r1 := Renderer{Name: "test", Host: "dog"}
	r2 := Renderer{Name: "test", Host: "dog"}
	r3 := Renderer{Name: "test2", Host: "dog"}
	r4 := Renderer{Name: "test", Host: "dogy"}

	if r1.GetID() != r2.GetID() || r1.GetID() == r3.GetID() || r1.GetID() == r4.GetID() || r3.GetID() == r4.GetID() {
		t.Error("Renderers don't have a unique ID based on name and host")
	}
}
