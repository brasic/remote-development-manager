package service

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type installState int

//go:generate stringer -type=installState
const (
	unknown installState = iota
	NotInstalled
	PlistPresentButNotLoaded
	Installed
)

// InstallState indicates whether our service is installed.
func (s *Service) InstallState() (state installState, err error) {
	plist, err := s.plistPath()
	if err != nil {
		return
	}

	if _, err = os.Stat(plist); errors.Is(err, os.ErrNotExist) {
		return NotInstalled, nil
	} else if err != nil {
		return
	}

	_, err = s.launchctlPrint()
	if err != nil {
		return PlistPresentButNotLoaded, nil
	}

	return Installed, nil
}

// Construct the path to the user's LaunchAgents dir and validate that it exists.
func launchAgentsDir() (dir string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	dir = filepath.Join(home, "Library", "LaunchAgents")

	stat, err := os.Stat(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return "", fmt.Errorf("Unexpected missing directory %s (%v)", dir, err)
	}

	if !stat.IsDir() {
		return "", fmt.Errorf("Uh, %s exists but is not a directory somehow?", dir)
	}
	return
}