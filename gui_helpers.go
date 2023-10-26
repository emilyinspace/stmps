package main

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wildeyedskies/stmp/mpvplayer"
)

func makeModal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}

func formatPlayerStatus(volume int64, position float64, duration float64) string {
	if position < 0 {
		position = 0.0
	}

	if duration < 0 {
		duration = 0.0
	}

	positionMin, positionSec := secondsToMinAndSec(position)
	durationMin, durationSec := secondsToMinAndSec(duration)

	return fmt.Sprintf("[%d%%][::b][%02d:%02d/%02d:%02d]", volume,
		positionMin, positionSec, durationMin, durationSec)
}

func formatSongForStatusBar(currentSong *mpvplayer.QueueItem) (text string) {
	if currentSong == nil {
		return
	}
	if currentSong.Title != "" {
		text += "[::-] [white]" + currentSong.Title
	}
	if currentSong.Artist != "" {
		text += " [gray]by [white]" + currentSong.Artist
	}
	return
}
