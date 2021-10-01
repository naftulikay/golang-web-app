package response

import (
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/views"
	"gopkg.in/guregu/null.v4"
	"net/http"
)

func Ok(req interfaces.RequestContext, data interface{}) *views.Response {
	return &views.Response{
		StatusCode: http.StatusOK,
		Error:      null.String{},
		Data:       data,
		Links: views.ResponseLinks{
			Self: req.Req().URL.Path,
		},
	}
}
