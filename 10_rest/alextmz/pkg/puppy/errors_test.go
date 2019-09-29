package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	var tests = []struct {
		arg  int
		want Error
	}{
		{arg: 400, want: Error{Message: "400 Bad Request", Code: 400}},
		{arg: Err400BadRequest, want: Error{Message: "400 Bad Request", Code: 400}},
		{arg: 404, want: Error{Message: "404 Not Found", Code: 404}},
		{arg: Err404NotFound, want: Error{Message: "404 Not Found", Code: 404}},
		{arg: 409, want: Error{Message: "409 Conflict", Code: 409}},
		{arg: Err409Conflict, want: Error{Message: "409 Conflict", Code: 409}},
		{arg: 500, want: Error{Message: "500 Internal Server Error", Code: 500}},
		{arg: Err500InternalError, want: Error{Message: "500 Internal Server Error", Code: 500}},
		{arg: 501, want: Error{Message: "501 Not Implemented", Code: 501}},
		{arg: Err501NotImplemented, want: Error{Message: "501 Not Implemented", Code: 501}},
		{arg: 171, want: Error{Message: "501 Not Implemented", Code: 171}},
		{arg: 666, want: Error{Message: "501 Not Implemented", Code: 666}},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.want.Message, func(t *testing.T) {
			got := NewError(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}
