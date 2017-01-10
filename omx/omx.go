package omx

import (
	"github.com/godbus/dbus"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type NoDbusAddress string

func (e NoDbusAddress) Error() string {
	return string(e)
}

type OmxInterface struct {
	conn   *dbus.Conn
	busObj dbus.BusObject
}

func NewOmxInterface() (obj *OmxInterface, warn error) {
	obj = &OmxInterface{}
	warn = obj.ConnectToPlayer()
	return
}

func (o *OmxInterface) StartPlayer(uri string) (err error) {
	cmd := exec.Command("omxplayer", "-g", "--no-keys", "-o", "local", uri)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	cmd.Dir = "/tmp/"
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	err = cmd.Start()
	if err == nil {
		cmd.Process.Release()
	}
	return
}

func (o *OmxInterface) ConnectToPlayer() (err error) {
	if o.conn != nil {
		o.conn.Close()
	}
	o.conn = nil
	o.busObj = nil

	username := os.Getenv("USER")
	contents, err := ioutil.ReadFile("/tmp/omxplayerdbus." + username)
	if err != nil {
		return
	}
	address := strings.TrimSpace(string(contents))
	if address == "" || address == "autolaunch:" {
		err = NoDbusAddress("no address specified")
		return
	}

	conn, err := dbus.Dial(address)
	if err != nil {
		return
	}

	if err = conn.Auth(nil); err != nil {
		conn.Close()
		return
	}

	if err = conn.Hello(); err != nil {
		conn.Close()
		return
	}

	busObject := conn.Object("org.mpris.MediaPlayer2.omxplayer",
		"/org/mpris/MediaPlayer2")
	o.conn = conn
	o.busObj = busObject
	return
}

func (o *OmxInterface) DisconnectPlayer() {
	if o.conn != nil {
		o.conn.Close()
	}
	o.conn = nil
	o.busObj = nil
}

func (o *OmxInterface) HasPlayer() bool {
	_, err := o.PlaybackStatus()
	return err == nil
}

func (o *OmxInterface) PlaybackStatus() (s string, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.PlaybackStatus", 0)
	err = call.Store(&s)
	if err == nil {
		s = strings.ToLower(s)
	}
	return
}

func (o *OmxInterface) Uri() (s string, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.GetSource", 0)
	err = call.Store(&s)
	return
}

func (o *OmxInterface) Duration() (i int64, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Duration", 0)
	err = call.Store(&i)
	return
}

func (o *OmxInterface) Position() (i int64, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Position", 0)
	err = call.Store(&i)
	return
}

func (o *OmxInterface) Volume() (d float64, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Volume", 0)
	err = call.Store(&d)
	return
}

func (o *OmxInterface) SetPosition(i int64) (out int64, err error) {
	call := o.busObj.Call("org.mpris.MediaPlayer2.Player.SetPosition", 0,
		dbus.ObjectPath("/not/used"), i)
	err = call.Store(&out)
	return
}

func (o *OmxInterface) SetVolume(d float64) (out float64, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Volume", 0, d)
	err = call.Store(&out)
	return
}

func (o *OmxInterface) Mute() (err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Mute", 0)
	err = call.Err
	return
}

func (o *OmxInterface) Unmute() (err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Unmute", 0)
	err = call.Err
	return
}

func (o *OmxInterface) Stop() (err error) {
	call := o.busObj.Call("org.mpris.MediaPlayer2.Player.Stop", 0)
	err = call.Err
	return
}

func (o *OmxInterface) PlayPause() (err error) {
	call := o.busObj.Call("org.mpris.MediaPlayer2.Player.PlayPause", 0)
	err = call.Err
	return
}

func (o *OmxInterface) Pause() (err error) {
	playbackStatus, err := o.PlaybackStatus()
	if err != nil {
		return
	}
	if playbackStatus == "playing" {
		err = o.PlayPause()
	}
	return
}

func (o *OmxInterface) Play() (err error) {
	playbackStatus, err := o.PlaybackStatus()
	if err != nil {
		return
	}
	if playbackStatus != "playing" {
		err = o.PlayPause()
	}
	return
}

func (o *OmxInterface) Seek(i int64) (out int64, err error) {
	call := o.busObj.Call("org.mpris.MediaPlayer2.Player.Seek", 0, i)
	err = call.Store(&out)
	return
}
