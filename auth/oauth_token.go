// Package auth holds OAuth authentication value types for the Zoho Payments SDK.
package auth

// OAuthToken is an OAuth access token returned by the Zoho IAM token endpoint.
type OAuthToken struct {
	AccessToken string
	ExpiresIn   int64
}
