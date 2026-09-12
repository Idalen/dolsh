package jellyfin

import (
	"context"
	"fmt"
	"log"
)

// AuthService owns everything related to authentication and authorization:
// credentials, the access token, the user ID, and the authorization header.
type AuthService struct {
	*Client

	token  string
	userID string
}

type credentials struct {
	Username string `json:"Username"`
	Password string `json:"Pw"`
}

type authorization struct {
	User struct {
		ID string `json:"Id"`
	} `json:"User"`
	AccessToken string `json:"AccessToken"`
}

func (s *AuthService) Login(ctx context.Context, username, password string) error {
	var (
		err error
		req credentials
		res authorization
	)

	log.Printf("authenticating user %q against jellyfin", username)

	req = credentials{
		Username: username,
		Password: password,
	}

	s.setAuthorization()

	res, err = s.transport.Request[credentials, authorization](ctx, "POST", "/Users/authenticatebyname", req)
	if err != nil {
		log.Printf("jellyfin authentication failed: %v", err)
		return fmt.Errorf("authentication failed: %w", err)
	}

	s.setToken(res.AccessToken)
	s.userID = res.User.ID
	
	log.Println("jellyfin authentication succeeded, token stored")
	
	return nil
}

// Token returns the access token for the authenticated session, if any.
func (s *AuthService) Token() string {
	return s.token
}

// UserID returns the authenticated user's ID, if any.
func (s *AuthService) UserID() string {
	return s.userID
}

func (s *AuthService) setToken(token string) {
	s.token = token
	s.setAuthorization()
}

func (s *AuthService) setAuthorization() {
	s.transport.SetHeader("X-Emby-Authorization", s.buildAuthorization())
}

func (s *AuthService) buildAuthorization() string {
	auth := fmt.Sprintf(
		`MediaBrowser Client="%s", Device="%s", DeviceId="%s", Version="%s"`,
		clientName, deviceName, deviceID, version,
	)
	if s.token != "" {
		auth += fmt.Sprintf(`, Token="%s"`, s.token)
	}
	return auth
}
