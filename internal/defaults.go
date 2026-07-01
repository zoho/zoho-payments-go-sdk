// Package internal holds the SDK's non-exported machinery and is not part of the public API.
package internal

import "time"

const (
	DefaultConnectTimeout = 30 * time.Second
	DefaultRequestTimeout = 60 * time.Second
)
