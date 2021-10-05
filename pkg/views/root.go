package views

import (
	null "gopkg.in/guregu/null.v4"
	"net/http"
)

type Response struct {
	StatusCode uint16        `json:"status_code"`
	Error      null.String   `json:"error" swaggertype:"string"`
	Data       interface{}   `json:"data"`
	Links      ResponseLinks `json:"_links"`
}

type ResponseLinks struct {
	Self string      `json:"_self"`
	Prev null.String `json:"_prev,omitempty" swaggertype:"string"`
	Next null.String `json:"_next,omitempty" swaggertype:"string"`
}

type PaginatedResponse struct {
	StatusCode uint16
	Error      null.String `swaggertype:"string"`
	Data       interface{}
	Count      uint
	Total      uint
	Links      ResponseLinks
}

func Success(value interface{}) Response {
	return Response{
		StatusCode: http.StatusOK,
		Error:      null.String{},
		Data:       value,
	}
}
