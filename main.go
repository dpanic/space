package main

import (
	_ "embed"
	"os"
	"space/backend"
	"space/backend/hooks"
	"time"
)

var (
	Version string
)

func main() {
	hooks.Version = Version
	if hooks.Version == "" {
		raw, err := os.ReadFile(".git/refs/heads/main")
		if err != nil {
			hooks.Version = err.Error()
		} else {
			hooks.Version = string(raw)[0:8]
		}
	}

	hooks.BuiltOn = time.Now().UTC().String()
	backend.Start()
}
