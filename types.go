package anon

import (
	"net/http"
	"time"

	"github.com/cretz/bine/tor"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Debug  bool
	Logger *logrus.Logger
}

// Client is used to make HTTP requests. It adds additional functionality
// like automatic retries to tolerate minor outages.
type Client struct {
	HTTPClient *http.Client   // Internal HTTP client.
	Conf       Config         // Client Config
	Logger     *logrus.Logger // Customer logger instance.

	Tor       *tor.Tor        // Tor process
	Transport *http.Transport // Transport for tor socks5 proxy

	RetryWaitMin time.Duration // Minimum time to wait
	RetryWaitMax time.Duration // Maximum time to wait
	RetryMax     int           // Maximum number of retries

	// RequestLogHook allows a user-supplied function to be called
	// before each retry.
	RequestLogHook RequestLogHook

	// ResponseLogHook allows a user-supplied function to be called
	// with the response from each HTTP request executed.
	ResponseLogHook ResponseLogHook

	// CheckRetry specifies the policy for handling retries, and is called
	// after each request. The default policy is DefaultRetryPolicy.
	CheckRetry CheckRetry

	// Backoff specifies the policy for how long to wait between retries
	Backoff Backoff

	// ErrorHandler specifies the custom error handler to use, if any
	ErrorHandler ErrorHandler
}
