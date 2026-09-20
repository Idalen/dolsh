package app

import "errors"

var ErrNoServerURL = errors.New("no server URL found")

var ErrNotAuthenticated = errors.New("not authenticated")
