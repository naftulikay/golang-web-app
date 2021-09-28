package views

import "gopkg.in/guregu/null.v4"

type Response struct {
	StatusCode uint16
	Error      null.String
	Data       interface{}
	Metadata   ResponseMeta
}

type ResponseMeta struct {
	Self string
}

type PaginatedResponse struct {
	StatusCode uint16
	Error      null.String
	Data       []interface{}
	Count      uint
	Total      uint
	Metadata   PaginatedResponseMeta
}

type PaginatedResponseMeta struct {
	Prev null.String
	Next null.String
	Self string
}
