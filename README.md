# Go OMX interface
This library provides an API interface to Raspberry Pi's omxplayer

## Downloading
(Not necessarily on RPi)

You need to have setup your $GOPATH

`
go get github.com/edudev/go-omx/
`

## Building example program
(Not necessarily on RPi)

`
GOOS=linux GOARCH=arm GOARM=7 go build -v github.com/edudev/go-omx/
`

## Running example program
(On RPi, as it is the only platform supporting omxplayer)

You need to copy the the previously built go-omx file to your RPi

`
./go-omx
`
