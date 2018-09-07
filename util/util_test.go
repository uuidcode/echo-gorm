package util

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetPageUrl(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/book?name=test&p=7", nil)
	url := GetUrl(req)
	assert.Equal(t, "/book?name=test&p=7", url)
}

func TestGetUrlAndRemovePageParam(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/book?name=test&p=7", nil)
	url := GetUrlAndRemovePageParam(req)
	assert.Equal(t, "/book?name=test&", url)

	req = httptest.NewRequest(echo.GET, "/book?name=test&p=7&title=method", nil)
	url = GetUrlAndRemovePageParam(req)
	assert.Equal(t, "/book?name=test&title=method&", url)

	req = httptest.NewRequest(echo.GET, "/book", nil)
	url = GetUrlAndRemovePageParam(req)
	assert.Equal(t, "/book?", url)

	req = httptest.NewRequest(echo.GET, "/book?p=12", nil)
	url = GetUrlAndRemovePageParam(req)
	assert.Equal(t, "/book?", url)
}
