package jellyfin

import (
	"context"
	"fmt"
)

type LibraryService struct {
	*Client
}

type Folder struct {
	Name           string   `json:"Name"`
	CollectionType string   `json:"CollectionType"`
	Locations      []string `json:"Locations"`
	ID             string   `json:"Id"`
}

type album struct {
	ID string `json:"Id"`
	Name string `json:"Name"`
}

func (s *LibraryService) VirtualFolders(ctx context.Context) ([]Folder, error) {
	res, err := s.transport.Request[struct{}, []Folder](ctx, "GET", "/Library/VirtualFolders", struct{}{})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *LibraryService) Albums(ctx context.Context) ([]album, error) {
	path := fmt.Sprintf("Users/%s/Items", s.UserID())
	
	albums, err := s.transport.Request[struct{}, []album](ctx, "GET", path, struct{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch albums from jellyfin: %w", err)
	}
	
	return albums, nil
}
