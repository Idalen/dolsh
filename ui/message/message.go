// Package message defines messages that screens send to the root model.
package message

import "dolsh/app"

// Error reports an error to the current screen.
type Error struct {
	Err error
}

type Folders struct {
	Folders []app.Folder
}

type Albums struct {
	Albums []app.Album
}

type Tracks struct {
	AlbumID string
	Tracks  []app.Track
}

type Next struct{}
