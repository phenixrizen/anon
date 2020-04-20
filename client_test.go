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
	conf := Config{
		Logger: log,
		Debug:  true,
		Region: "ru",
	}
	anonHttpClient, err := NewClient(conf)
	require.NoError(err)
	req, err := NewRequest("GET", "http://www.nathanrockhold.com", nil)
	require.NoError(err)
	resp, err := anonHttpClient.Do(req)
	require.NoError(err)
	require.Equal(http.StatusOK, resp.StatusCode)
	require.Equal("Russia", anonHttpClient.GeoIPData.Country)
	err = anonHttpClient.Shutdown()
	require.NoError(err)

	conf2 := Config{
		Logger: log,
		Debug:  true,
		Region: "ua",
	}
	anonHttpClient2, err := NewClient(conf2)
	require.NoError(err)
	req2, err := NewRequest("GET", "http://www.nathanrockhold.com", nil)
	require.NoError(err)
	resp2, err := anonHttpClient2.Do(req2)
	require.NoError(err)
	require.Equal(http.StatusOK, resp2.StatusCode)
	require.Equal("Ukraine", anonHttpClient2.GeoIPData.Country)
	err = anonHttpClient2.Shutdown()
	require.NoError(err)
}
