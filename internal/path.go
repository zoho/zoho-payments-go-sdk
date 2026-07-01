package internal

import "net/url"

func EncodePathSegment(segment string) string {
	return url.PathEscape(segment)
}
