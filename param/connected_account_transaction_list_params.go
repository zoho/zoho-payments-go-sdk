package param

import "net/url"

// ConnectedAccountTransactionListParams are the optional filters for listing a connected account's transactions.
type ConnectedAccountTransactionListParams struct {
	SortColumn        *string
	TransactionType   *string
	TransactionID     *string
	FilterBy          *string
	FromDate          *string
	ToDate            *string
	PaymentMethodType *string
	CardBrand         *string
	CardType          *string
	PerPage           *int
	Page              *int
}

func (p *ConnectedAccountTransactionListParams) ToQuery() url.Values {
	query := url.Values{}
	if p == nil {
		return query
	}
	qSetStr(query, "sort_column", p.SortColumn)
	qSetStr(query, "transaction_type", p.TransactionType)
	qSetStr(query, "transaction_id", p.TransactionID)
	qSetStr(query, "filter_by", p.FilterBy)
	qSetStr(query, "from_date", p.FromDate)
	qSetStr(query, "to_date", p.ToDate)
	qSetStr(query, "payment_method_type", p.PaymentMethodType)
	qSetStr(query, "card_brand", p.CardBrand)
	qSetStr(query, "card_type", p.CardType)
	qSetInt(query, "per_page", p.PerPage)
	qSetInt(query, "page", p.Page)
	return query
}
