package handler

import (
	"github.com/kirillrdy/vidos/fs"
	"github.com/kirillrdy/vidos/path"
	"log"
	"net/http"
	"os"
)

//TODO serve real filename
//Stream actually streams video content, respecting Range header
func Stream(response http.ResponseWriter, request *http.Request) {
	log.Printf("Streaming byte range: %v", request.Header.Get("Range"))
	video := videoFromRequest(request)
	http.ServeFile(response, request, fs.VideosDataDir+string(os.PathSeparator)+video.Filepath)
}

func videoFromRequest(request *http.Request) fs.Video {
	return fs.Video{Filepath: request.FormValue(path.ParamKeys.Filepath)}
}
