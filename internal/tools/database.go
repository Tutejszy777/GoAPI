package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database return
type LoginDetails struct {
	AuthToken string
	Username string
}

type CoinsDetails struct {
	Username string
	Coins int64
}

