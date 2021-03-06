package view

import (
	"bytes"
	"fmt"
	"github.com/anacrolix/torrent"
	humanise "github.com/dustin/go-humanize"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/web/html"
)

// TorrentsList is the page that lists all the torrents
func TorrentsList(torrents []*torrent.Torrent) html.Node {
	return application.NewPage("Torrent", path.Torrents).ToHTML(TorrentsTable(torrents))
}

// TorrentsStatus displayes torrent status page
func TorrentStatus(buffer *bytes.Buffer) html.Node {
	return application.NewPage("Torrent Status", path.TorrentStatus).ToHTML(html.Pre().Text(buffer.String()))
}

//TorrentsTable should be selfexplanitory
func TorrentsTable(torrents []*torrent.Torrent) html.Node {

	if len(torrents) == 0 {
		return centerByBoxes(html.H4().Text("No torrents have been added"))
	}

	page := html.Div().Children(
		html.Table().Class(tableClass).Children(
			html.Thead().Children(
				html.Tr().Children(
					html.Th().Text("Name"),
					html.Th().Text("Bytes Completed"),
					html.Th().Text("Length"),
					html.Th().Text("%"),
					html.Th().Text("Finished"),
				),
			),

			html.Tbody().Children(
				torrentTrs(torrents)...,
			),
		),
	)

	return page
}

func torrentTrs(torrents []*torrent.Torrent) []html.Node {
	var nodes []html.Node
	for index := range torrents {
		nodes = append(nodes, torrentTr(torrents[index]))
	}
	return nodes
}

func torrentTr(torrent *torrent.Torrent) html.Node {
	percent := 100 * torrent.BytesCompleted() / torrent.Length()
	return html.Tr().Children(
		html.Td().Text(torrent.Name()),
		html.Td().Text(humanise.Bytes(uint64(torrent.BytesCompleted()))),
		html.Td().Text(humanise.Bytes(uint64(torrent.Length()))),
		html.Td().Text(fmt.Sprint(percent)),
		html.Td().Text(fmt.Sprint(torrent.Length() == torrent.BytesCompleted())),
	)

}
