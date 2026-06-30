package model

type ListResponse[T any] struct {
	Data        []T
	PageContext PageContext
}
