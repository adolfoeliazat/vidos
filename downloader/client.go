package downloader

import (
	"github.com/anacrolix/torrent"
	"log"
)

const FileDir = "files/"

var Client *torrent.Client

func init() {

	config := torrent.Config{DataDir: FileDir}
	var err error
	Client, err = torrent.NewClient(&config)

	//TODO pehaps allow client to be nil for the whole app
	if err != nil {
		log.Panic(err)
	}

	//TODO implement proper shutdown, and call Close as part of it
	//defer Client.Close()

	//TODO also possibly part of Client.Close()
	//Client.WaitAll()
}
