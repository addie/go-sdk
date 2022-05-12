/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package r2

import (
	"net/http"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestEventOptions(t *testing.T) {
	assert := assert.New(t)

	e := NewEvent(Flag)

	assert.Nil(e.Request)
	OptEventRequest(&http.Request{RequestURI: "abcdef"})(&e)
	assert.NotNil(e.Request)
	assert.Equal("abcdef", e.Request.RequestURI)

	assert.Nil(e.Response)
	OptEventResponse(&http.Response{Proto: "not-http"})(&e)
	assert.NotNil(e.Response)
	assert.Equal("not-http", e.Response.Proto)

	assert.Nil(e.Body)
	OptEventBody([]byte(`example-string`))(&e)
	assert.NotNil(e.Body)
	assert.Equal("example-string", string(e.Body))
}
