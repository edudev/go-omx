# Go OMX interface
This binary provides a server with JSON API interface to Raspberry Pi's omxplayer

## Downloading
(Not necessarily on RPi)

You need to have setup your $GOPATH

`
go get github.com/edudev/go-omx/
cd $GOPATH/src/github.com/edudev/go-omx/
git checkout backend
cd -
`

## Building example program
(Not necessarily on RPi)

`
GOOS=linux GOARCH=arm GOARM=7 go build -v github.com/edudev/go-omx/backend/
`

## Running example program
(On RPi, as it is the only platform supporting omxplayer)

You need to copy the the previously built backend file to your RPi

`
./backend <media_files_dir> <media_files_url>
`
The best way to serve the media files is with an http server (like nginx)
Nginx setup is beyond the scope of this README
