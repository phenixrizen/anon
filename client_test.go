package anon

import (
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetGeoIPData(t *testing.T) {
	require := require.New(t)
	log := logrus.New()
	anonHttpClient, err := NewClient(Config{Logger: log, Debug: true})
	require.NoError(err)
	req, err := NewRequest("GET", "http://www.nathanrockhold.com", nil)
	require.NoError(err)
	resp, err := anonHttpClient.Do(req)
	require.NoError(err)
	require.Equal(http.StatusOK, resp.StatusCode)
	err = anonHttpClient.Shutdown()
	require.NoError(err)
}
