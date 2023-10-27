// Copyright 2023 The STMP Authors
// SPDX-License-Identifier: GPL-3.0-or-later

package mpvplayer

type QueueItem struct {
	Id       string
	Uri      string
	Title    string
	Artist   string
	Duration int
}

// StatusData is a player progress report for the UI
type StatusData struct {
	Volume   int64
	Position int64
	Duration int64
}
