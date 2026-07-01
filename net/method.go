// Package net is the pluggable HTTP transport layer for the Zoho Payments SDK.
package net

// Method is an HTTP request method.
type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)
