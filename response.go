package response

import (
	"net/http"

	"github.com/go-chi/render"
)

// Response defines http response for the client
type Response struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Error      Error       `json:"error"`
	StatusCode int         `json:"-"`
}

// Error defines the error
type Error struct {
	Status bool   `json:"status"` // true if we have error
	Msg    string `json:"msg"`    // error message
	Code   int    `json:"code"`   // error code from affiliate, it is not http status code
}

// DefaultSuccessResponse is default response for
// inform user that proccess is success
type DefaultSuccessResponse struct {
	Success bool `json:"success"`
}

// SetError set the response to return the given error.
// code is http status code, http.StatusInternalServerError is the default value
func (res *Response) SetError(err error, code ...int) {
	if len(code) > 0 {
		res.StatusCode = code[0]
	} else {
		res.StatusCode = http.StatusInternalServerError
	}

	if err != nil {
		res.Error = Error{
			Msg:    err.Error(),
			Status: true,
		}
	}

}

// SetSuccess function will set default success response as data
func (res *Response) SetSuccess() {
	res.Data = DefaultSuccessResponse {
		Success:true,
	}
}

// Render writes the http response to the client
func (res *Response) Render(w http.ResponseWriter, r *http.Request) {
	if res.StatusCode == 0 {
		res.StatusCode = http.StatusOK
	}

	render.Status(r, res.StatusCode)
	render.JSON(w, r, res)
}
