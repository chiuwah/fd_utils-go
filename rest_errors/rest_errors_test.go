package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("This is the message.", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "This is the message.", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t,"database error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("This is the message.")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "This is the message.", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {

}

func TestNewError(t *testing.T) {

}
