package views

import "gopkg.in/guregu/null.v4"

type Response struct {
	StatusCode uint16
	Error      null.String
	Data       interface{}
	Links      ResponseLinks `json:"_links"`
}

type ResponseLinks struct {
	Self string      `json:"_self"`
	Prev null.String `json:"_prev,omitempty"`
	Next null.String `json:"_next,omitempty"`
}

type PaginatedResponse struct {
	StatusCode uint16
	Error      null.String
	Data       []interface{}
	Count      uint
	Total      uint
	Links      ResponseLinks
}
