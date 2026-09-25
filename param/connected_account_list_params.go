package param

import "net/url"

// ConnectedAccountListParams are the optional filters for listing connected accounts.
type ConnectedAccountListParams struct {
	ConnectedAccountID *string
	FilterBy           *string
	FromDate           *string
	ToDate             *string
	PerPage            *int
	Page               *int
}

func (p *ConnectedAccountListParams) ToQuery() url.Values {
	query := url.Values{}
	if p == nil {
		return query
	}
	qSetStr(query, "connected_account_id", p.ConnectedAccountID)
	qSetStr(query, "filter_by", p.FilterBy)
	qSetStr(query, "from_date", p.FromDate)
	qSetStr(query, "to_date", p.ToDate)
	qSetInt(query, "per_page", p.PerPage)
	qSetInt(query, "page", p.Page)
	return query
}
