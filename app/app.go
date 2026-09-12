package app

import (
	"context"
	"fmt"
	"net/url"

	"dolsh/client/jellyfin"
	configstore "dolsh/repo/config"
	librarystore "dolsh/repo/library"
)

type App struct {
	jellyfin *jellyfin.Client

	config       config
	
	libraryRepo libraryRepo
	configRepo  configRepo
}

type libraryRepo interface {
	GetTracks() error
}

type configRepo interface {
	Load() (config, error)
	Save(*config) error
}

type config struct {
	URL         string `toml:"url"`
	AccessToken string `toml:"access_token"`
	UserID      string `toml:"user_id"`
	LibraryID   string `toml:"library_id"`
}

// SessionState describes the authentication state of the app at startup.
type SessionState int

const (
	SessionNeedsSetup SessionState = iota
	SessionNeedsLogin
	SessionNeedsLibrary
	SessionReady
)

func New() (*App, error) {
	configRepo, err := configstore.New[config]()
	if err != nil {
		return nil, fmt.Errorf("failed to start config store: %w", err)
	}

	libraryRepo, err := librarystore.New()
	if err != nil {
		return nil, fmt.Errorf("failed to start library store: %w", err)
	}

	cfg, err := configRepo.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	return &App{
		config: cfg,
		configRepo:  configRepo,
		libraryRepo: libraryRepo,
	}, nil
}

func (a *App) SetServerURL(rawURL string) error {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("invalid URL: %q", rawURL)
	}

	a.config.URL = u.String()
	if err := a.configRepo.Save(&a.config); err != nil {
		return fmt.Errorf("save config: %w", err)
	}

	return nil
}

func (a *App) HasServerURL() bool {
	return a.config.URL != ""
}

func (a *App) ServerURL() string {
	return a.config.URL
}

func (a *App) serverURL() (*url.URL, error) {
	if !a.HasServerURL() {
		return nil, ErrNoServerURL
	}
	return url.ParseRequestURI(a.config.URL)
}

func (a *App) Login(ctx context.Context, username, password string) (err error) {
	u, err := a.serverURL()
	if err != nil {
		return err
	}

	a.jellyfin, err = jellyfin.New(u, jellyfin.WithCredentials(ctx, username, password))
	if err != nil {
		return err
	}

	a.config.AccessToken = a.jellyfin.Token()
	a.config.UserID = a.jellyfin.UserID()
	if err := a.configRepo.Save(&a.config); err != nil {
		return fmt.Errorf("save config: %w", err)
	}
	return nil
}

func (a *App) RestoreSession() SessionState {
	if !a.HasServerURL() {
		return SessionNeedsSetup
	}

	if a.config.AccessToken == "" || a.config.UserID == "" {
		return SessionNeedsLogin
	}

	u, err := a.serverURL()
	if err != nil {
		return SessionNeedsSetup
	}

	client, err := jellyfin.New(u, jellyfin.WithToken(a.config.AccessToken, a.config.UserID))
	if err != nil {
		return SessionNeedsSetup
	}

	a.jellyfin = client

	if a.config.LibraryID == "" {
		return SessionNeedsLibrary
	}

	return SessionReady
}

func (a *App) SelectLibrary(id string) error {
	a.config.LibraryID = id
	if err := a.configRepo.Save(&a.config); err != nil {
		return fmt.Errorf("save config: %w", err)
	}
	return nil
}

func (a *App) ListFolders(ctx context.Context) ([]jellyfin.Folder, error) {
	if _, err := a.serverURL(); err != nil {
		return nil, err
	}

	folders, err := a.jellyfin.Library.VirtualFolders(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't get virtual folders: %w", err)
	}

	return folders, nil
}

func (a *App) Albums() ([]album) 
