package gocache

import "time"

var (
	DefaultExpiration = time.Minute * 30
)

type GttStringFunc func() string
type GttInterfaceFunc func() interface{}
