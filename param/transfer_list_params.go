package param

import "net/url"

// TransferListParams are the optional filters for listing transfers.
type TransferListParams struct {
	Status             *string
	FilterBy           *string
	PaymentID          *string
	ConnectedAccountID *string
	FromDate           *string
	ToDate             *string
	SortColumn         *string
	SortOrder          *string
	SearchText         *string
	PerPage            *int
	Page               *int
}

func (p *TransferListParams) ToQuery() url.Values {
	query := url.Values{}
	if p == nil {
		return query
	}
	qSetStr(query, "status", p.Status)
	qSetStr(query, "filter_by", p.FilterBy)
	qSetStr(query, "payment_id", p.PaymentID)
	qSetStr(query, "connected_account_id", p.ConnectedAccountID)
	qSetStr(query, "from_date", p.FromDate)
	qSetStr(query, "to_date", p.ToDate)
	qSetStr(query, "sort_column", p.SortColumn)
	qSetStr(query, "sort_order", p.SortOrder)
	qSetStr(query, "search_text", p.SearchText)
	qSetInt(query, "per_page", p.PerPage)
	qSetInt(query, "page", p.Page)
	return query
}
