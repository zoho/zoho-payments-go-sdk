package param

import "net/url"

// PayoutTransactionListParams are the optional filters for listing the transactions of a payout.
type PayoutTransactionListParams struct {
	PerPage *int
	Page    *int
}

func (p *PayoutTransactionListParams) ToQuery() url.Values {
	query := url.Values{}
	if p == nil {
		return query
	}
	qSetInt(query, "per_page", p.PerPage)
	qSetInt(query, "page", p.Page)
	return query
}
