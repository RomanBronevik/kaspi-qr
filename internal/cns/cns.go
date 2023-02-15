package cns

import "time"

const (
	MaxHeaderBytes = 1 << 20
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
)
