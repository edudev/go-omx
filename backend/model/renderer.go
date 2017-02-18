package model

type PlayerInterface interface {
	StartPlayer(uri string) error
	ConnectToPlayer() error
	DisconnectPlayer()

	HasPlayer() bool
	PlaybackStatus() (string, error)
	Uri() (string, error)
	Duration() (int64, error)
	Position() (int64, error)
	Volume() (float64, error)

	SetPosition(int64) (int64, error)
	SetVolume(float64) (float64, error)

	Mute() error
	Unmute() error
	Stop() error

	PlayPause() error
	Pause() error
	Play() error

	Seek(int64) (int64, error)
}

type Renderer struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Interface PlayerInterface `json:"-"`
}

func (c Renderer) GetID() string {
	return c.Name + "@" + c.Host
}
