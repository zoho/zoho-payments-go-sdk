package param

import "net/url"

// VirtualAccountPaymentListParams are the optional filters for listing virtual account payments.
type VirtualAccountPaymentListParams struct {
	Status     *string
	SortColumn *string
	SortOrder  *string
	PerPage    *int
	Page       *int
}

func (p *VirtualAccountPaymentListParams) ToQuery() url.Values {
	query := url.Values{}
	if p == nil {
		return query
	}
	qSetStr(query, "status", p.Status)
	qSetStr(query, "sort_column", p.SortColumn)
	qSetStr(query, "sort_order", p.SortOrder)
	qSetInt(query, "per_page", p.PerPage)
	qSetInt(query, "page", p.Page)
	return query
}
