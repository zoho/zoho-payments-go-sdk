package internal

import "net/url"

func EncodePathSegment(s string) string {
	return url.PathEscape(s)
}
