package omx

import (
	"github.com/godbus/dbus"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// NoDbusAddress is an Error type
type NoDbusAddress string

// Error returns the error as a string
func (e NoDbusAddress) Error() string {
	return string(e)
}

// Interface has all the methods to manipulate the omxplayer
type Interface struct {
	conn   *dbus.Conn
	busObj dbus.BusObject
}

// NewInterface returns a new Interface
func NewInterface() (obj *Interface, warn error) {
	obj = &Interface{}
	warn = obj.ConnectToPlayer()
	return
}

// StartPlayer boots up omxplayer (if possible)
func (o *Interface) StartPlayer(uri string) (err error) {
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
		go cmd.Wait()
	}
	return
}

// ConnectToPlayer attaches to the player (if possible)
func (o *Interface) ConnectToPlayer() (err error) {
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

// DisconnectPlayer dettaches from the player (if possible)
func (o *Interface) DisconnectPlayer() {
	if o.conn != nil {
		o.conn.Close()
	}
	o.conn = nil
	o.busObj = nil
}

// HasPlayer determines whether there is a running player
func (o *Interface) HasPlayer() bool {
	if o.busObj == nil {
		return false
	}
	_, err := o.PlaybackStatus()
	return err == nil
}

// PlaybackStatus gets whether or the player is playing, paused or other
func (o *Interface) PlaybackStatus() (s string, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.PlaybackStatus", 0)
	err = call.Store(&s)
	if err == nil {
		s = strings.ToLower(s)
	}
	return
}

// URI gets the currently loaded uri to the player
func (o *Interface) URI() (s string, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.GetSource", 0)
	err = call.Store(&s)
	return
}

// Duration gets the currently loaded media's length
func (o *Interface) Duration() (i int64, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Duration", 0)
	err = call.Store(&i)
	return
}

// Position gets the currently loaded media's playback position
func (o *Interface) Position() (i int64, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Position", 0)
	err = call.Store(&i)
	return
}

// Volume gets the player's volume setting
func (o *Interface) Volume() (d float64, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Volume", 0)
	err = call.Store(&d)
	return
}

// SetPosition moves the player playback position to the specified integer
func (o *Interface) SetPosition(i int64) (out int64, err error) {
	call := o.busObj.Call("org.mpris.MediaPlayer2.Player.SetPosition", 0,
		dbus.ObjectPath("/not/used"), i)
	err = call.Store(&out)
	return
}

// SetVolume is used to control the player's volume setting
func (o *Interface) SetVolume(d float64) (out float64, err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Volume", 0, d)
	err = call.Store(&out)
	return
}

// Mute is used to control the player's volume setting (mute all sound)
func (o *Interface) Mute() (err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Mute", 0)
	err = call.Err
	return
}

// Unmute is used to control the player's volume setting (unmute all sound)
func (o *Interface) Unmute() (err error) {
	call := o.busObj.Call("org.freedesktop.DBus.Properties.Unmute", 0)
	err = call.Err
	return
}

// Stop instructs the player to stop playing and exit
func (o *Interface) Stop() (err error) {
	call := o.busObj.Call("org.mpris.MediaPlayer2.Player.Stop", 0)
	err = call.Err
	return
}

// PlayPause is used to toggle between playback states
func (o *Interface) PlayPause() (err error) {
	call := o.busObj.Call("org.mpris.MediaPlayer2.Player.PlayPause", 0)
	err = call.Err
	return
}

// Pause is used to pause the player's playback
func (o *Interface) Pause() (err error) {
	playbackStatus, err := o.PlaybackStatus()
	if err != nil {
		return
	}
	if playbackStatus == "playing" {
		err = o.PlayPause()
	}
	return
}

// Play is used to resume the player's playback
func (o *Interface) Play() (err error) {
	playbackStatus, err := o.PlaybackStatus()
	if err != nil {
		return
	}
	if playbackStatus != "playing" {
		err = o.PlayPause()
	}
	return
}

// Seek is used for relative seeking on the player's position
func (o *Interface) Seek(i int64) (out int64, err error) {
	call := o.busObj.Call("org.mpris.MediaPlayer2.Player.Seek", 0, i)
	err = call.Store(&out)
	return
}
