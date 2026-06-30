package param

import "net/url"

// CustomerListParams are the optional filters for listing customers (US only).
type CustomerListParams struct {
	FilterBy *string
	FromDate *string
	ToDate   *string
	PerPage  *int
	Page     *int
}

func (p *CustomerListParams) ToQuery() url.Values {
	query := url.Values{}
	if p == nil {
		return query
	}
	qSetStr(query, "filter_by", p.FilterBy)
	qSetStr(query, "from_date", p.FromDate)
	qSetStr(query, "to_date", p.ToDate)
	qSetInt(query, "per_page", p.PerPage)
	qSetInt(query, "page", p.Page)
	return query
}
