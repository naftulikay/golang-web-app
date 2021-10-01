package response

import (
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/views"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"strings"
)

func BadRequest(req interfaces.RequestContext, msg string) *views.Response {
	return &views.Response{
		StatusCode: http.StatusBadRequest,
		Error:      null.StringFrom(msg),
		Data:       nil,
		Links: views.ResponseLinks{
			Self: req.Req().URL.Path,
		},
	}
}

func Forbidden(req interfaces.RequestContext, msg ...string) *views.Response {
	return &views.Response{
		StatusCode: http.StatusForbidden,
		Error:      null.StringFrom(strings.Join(msg, "\n")),
		Data:       nil,
		Links: views.ResponseLinks{
			Self: req.Req().URL.Path,
		},
	}
}
