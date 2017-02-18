package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/edudev/go-omx/omx"
)

func executeCommand(o *omx.Interface, cmd string) bool {
	split := strings.Split(cmd, " ")
	if len(split) == 0 {
		return false
	}

	switch split[0] {
	case "exit":
		return true
	case "startPlayer":
		if len(split) == 1 {
			fmt.Println("not enough args")
			break
		}
		err := o.StartPlayer(split[1])
		fmt.Println(err)
	case "connect":
		err := o.ConnectToPlayer()
		if err != nil {
			fmt.Println(err)
		}
	case "hasPlayer":
		hasPlayer := o.HasPlayer()
		fmt.Println(hasPlayer)
	case "playbackStatus":
		status, err := o.PlaybackStatus()
		fmt.Println(status, err)
	case "uri":
		uri, err := o.URI()
		fmt.Println(uri, err)
	case "duration":
		duration, err := o.Duration()
		fmt.Println(duration, err)
	case "position":
		if len(split) == 1 {
			position, err := o.Position()
			fmt.Println(position, err)
		} else {
			position, err := strconv.ParseInt(split[1], 10, 64)
			if err != nil {
				fmt.Println(err)
				break
			}
			position, err = o.SetPosition(position)
			fmt.Println(position, err)
		}
	case "volume":
		if len(split) == 1 {
			volume, err := o.Volume()
			fmt.Println(volume, err)
		} else {
			volume, err := strconv.ParseFloat(split[1], 64)
			if err != nil {
				fmt.Println(err)
				break
			}
			volume, err = o.SetVolume(volume)
			fmt.Println(volume, err)
		}
	case "mute":
		err := o.Mute()
		fmt.Println(err)
	case "unmute":
		err := o.Unmute()
		fmt.Println(err)
	case "stop":
		err := o.Stop()
		fmt.Println(err)
	case "playPause":
		err := o.PlayPause()
		fmt.Println(err)
	case "pause":
		err := o.Pause()
		fmt.Println(err)
	case "play":
		err := o.Play()
		fmt.Println(err)
	case "seek":
		if len(split) == 1 {
			fmt.Println("not enough args")
			break
		}

		seekBy, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		seekBy, err = o.Seek(seekBy)
		fmt.Println(seekBy, err)
	default:
		fmt.Printf("Invalid command %q\n", cmd)
	}
	return false
}

func main() {
	obj, warn := omx.NewInterface()
	if warn != nil {
		fmt.Println(warn)
	}
	defer obj.DisconnectPlayer()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		cmd = strings.TrimSpace(cmd)
		needExit := executeCommand(obj, cmd)
		if needExit {
			break
		}
	}
}
