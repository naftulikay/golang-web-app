package middleware

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestExtractTokenFromRequest(t *testing.T) {
	r, err := http.NewRequest("GET", "http://localhost:8080/some/arbitrary/path",
		bytes.NewReader(make([]byte, 0)))

	if err != nil {
		t.Errorf("unable to create http request: %s", err)
	}

	// if there is no authorization header, it should return nil
	assert.Nil(t, extractTokenFromRequest(r))

	// if there is a non-bearer value in the header, it should return nil
	r.Header.Set("Authorization", "Basic abc")
	assert.Nil(t, extractTokenFromRequest(r))

	// if there's no keyword, it should also return nil
	r.Header.Set("Authorization", "abcdefg")
	assert.Nil(t, extractTokenFromRequest(r))

	// if there's a badly formed request with no payload, it should also return nil
	r.Header.Set("Authorization", "Bearer   ")
	assert.Nil(t, extractTokenFromRequest(r))

	// if it works as expected, should return the payload
	r.Header.Set("Authorization", "Bearer abc")
	assert.NotNil(t, extractTokenFromRequest(r))
	assert.Equal(t, "abc", *extractTokenFromRequest(r))

	// test that surrounding whitespace does not matter
	r.Header.Set("Authorization", "Bearer    abc    ")
	assert.NotNil(t, extractTokenFromRequest(r))
	assert.Equal(t, "abc", *extractTokenFromRequest(r))
}
