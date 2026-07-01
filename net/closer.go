package net

// Closer is an optional interface an HTTPClient may implement to release transport-level resources.
type Closer interface {
	Close() error
}
