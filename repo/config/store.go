// Package config provides local configuration management for dolsh application
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Store[T any] struct {
	filePath string
}

func New[T any]() (*Store[T], error) {
	filePath, err := configPath()
	if err != nil {
		return nil, err
	}

	return &Store[T]{
		filePath: filePath,
	}, nil
}

func (s *Store[T]) Load() (cfg T, err error) {
	var data []byte	
	data, err = os.ReadFile(s.filePath)
	if errors.Is(err, os.ErrNotExist) {
		err = nil
		return
	}
	if err != nil {
		return
	}

	if _, err = toml.Decode(string(data), &cfg); err != nil {
		return
	}

	return
}

func (s *Store[T]) Save(config *T) error {
	data, err := toml.Marshal(config)
	if err != nil {
		return fmt.Errorf("encode config: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(s.filePath), 0o700); err != nil {
		return err
	}

	if err := os.WriteFile(s.filePath, data, 0o600); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}

func configPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "dolsh", "config.toml"), nil
}
