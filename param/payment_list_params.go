package param

import "net/url"

// PaymentListParams are the optional filters for listing payments.
type PaymentListParams struct {
	Status            *string
	SearchText        *string
	FilterBy          *string
	FromDate          *string
	ToDate            *string
	PaymentMethodType *string
	PerPage           *int
	Page              *int
}

func (p *PaymentListParams) ToQuery() url.Values {
	query := url.Values{}
	if p == nil {
		return query
	}
	qSetStr(query, "status", p.Status)
	qSetStr(query, "search_text", p.SearchText)
	qSetStr(query, "filter_by", p.FilterBy)
	qSetStr(query, "from_date", p.FromDate)
	qSetStr(query, "to_date", p.ToDate)
	qSetStr(query, "payment_method_type", p.PaymentMethodType)
	qSetInt(query, "per_page", p.PerPage)
	qSetInt(query, "page", p.Page)
	return query
}
