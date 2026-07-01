package param

// BrowserInfo carries browser metadata for 3DS / customer-on-session flows.
type BrowserInfo struct {
	UserAgent      *string
	AcceptHeader   *string
	ScreenHeight   *int
	ScreenWidth    *int
	Language       *string
	TimeZoneOffset *int
	ColorDepth     *int
}

func (b *BrowserInfo) toMap() map[string]any {
	body := map[string]any{}
	putStr(body, "user_agent", b.UserAgent)
	putStr(body, "accept_header", b.AcceptHeader)
	putInt(body, "screen_height", b.ScreenHeight)
	putInt(body, "screen_width", b.ScreenWidth)
	putStr(body, "language", b.Language)
	putInt(body, "time_zone_offset", b.TimeZoneOffset)
	putInt(body, "color_depth", b.ColorDepth)
	return body
}
