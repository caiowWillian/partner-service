package encodedError

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SendAndResponse struct {
	send     error
	response int
}

func TestGetCode(t *testing.T) {
	sendAndResponse := []SendAndResponse{
		{BadRequest, 400},
		{InternalServerError, 500},
		{errors.New("test"), 500},
		{NoContent, 204},
	}

	for _, item := range sendAndResponse {
		assert.Equal(t, getCode(item.send), item.response)
	}
}
