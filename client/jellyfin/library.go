package jellyfin

import (
	"context"
	"dolsh/client/jellyfin/transport"
	"fmt"
	"log"
)

type LibraryService struct {
	*Client
}

type Folder struct {
	Name           string   `json:"Name"`
	CollectionType string   `json:"CollectionType"`
	Locations      []string `json:"Locations"`
	ID             string   `json:"ItemId"`
}

type Album struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

type albumsResponse struct {
	Items []Album `json:"Items"`
}

func (s *LibraryService) VirtualFolders(ctx context.Context) ([]Folder, error) {
	res, err := s.transport.Request[struct{}, []Folder](ctx, "GET", "/Library/VirtualFolders", struct{}{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *LibraryService) Albums(ctx context.Context, folderID string) ([]Album, error) {
	path := fmt.Sprintf("Users/%s/Items", s.UserID())

	res, err := s.transport.Request[struct{}, albumsResponse](
		ctx,
		"GET",
		path,
		struct{}{},
		transport.WithMusicAlbumType(),
		transport.WithParentID(folderID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch albums from jellyfin: %w", err)
	}

	log.Printf("Albums response size: %d", len(res.Items))

	return res.Items, nil
}
