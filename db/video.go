package db

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kirillrdy/vidos/ffmpeg"
)

const dataDir = "data"

type Video struct {
	Id       uint64
	Filename string
	Encoded  bool
	Progress string
	Duration string
}

func (video Video) dataDirPath() string {
	return fmt.Sprintf("%v/%v", dataDir, video.Id)
}
func (video Video) FilePath() string {
	return fmt.Sprintf("%v/%v", video.dataDirPath(), video.Filename)
}

func (video Video) IdString() string {
	return fmt.Sprint(video.Id)
}

func (video Video) Mkdir() {
	err := os.MkdirAll(video.dataDirPath(), os.ModePerm)
	if err != nil {
		log.Print(err)
	}
}

func (video Video) Reencode() {
	video.Encoded = false
	result := Session.Save(&video)
	if result.Error != nil {
		log.Print(result.Error)
	}

	video.StartEncoding()
}

func (video Video) StartEncoding() {
	go func() {
		video.Encode()
		video.Encoded = true
		video.Progress = ""
		Session.Save(&video)
	}()
}

func (video Video) Encode() {
	update := func(timeProgress string) {
		video.Progress = timeProgress
		//TODO errors
		Session.Save(&video)
	}
	ffmpeg.Encode(video.FilePath(), video.EncodedPath(), update)
}

func (video *Video) CalculateDuration() {
	video.Duration = ffmpeg.Duration(video.FilePath())
	Session.Save(video)
}

func (video Video) EncodedPath() string {
	//TODO do basename to strip out ext
	return video.FilePath() + ".mp4"
}

//TODO get rid of log.Fatal
func (video Video) Save(reader io.ReadCloser) {

	video.Mkdir()

	destinationFile, err := os.Create(video.FilePath())

	if err != nil {
		log.Fatal(err)
	}

	n, err := io.Copy(destinationFile, reader)
	if n == 0 || err != nil {
		log.Fatal(err)
	}

	defer reader.Close()
	defer destinationFile.Close()
}