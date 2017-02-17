package model

type Renderer struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
}

func (c Renderer) GetID() string {
    return c.Name + "@" + c.Host
}

func (c *Renderer) SetID(id string) error {
	return nil
}
