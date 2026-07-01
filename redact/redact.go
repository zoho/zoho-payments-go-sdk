// Package redact provides the placeholder and helper used to mask the OAuth access token when SDK types that carry it are formatted with %v / %+v / %#v.
package redact

const Mask = "[REDACTED]"

func Token(secret string) string {
	if secret == "" {
		return ""
	}
	return Mask
}
