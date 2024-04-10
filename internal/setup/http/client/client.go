package client

import (
	"errors"
	"net/http"
	"time"
)

var (
	errRedirectNotAllowed = errors.New("redirect not allowed")
)

type HttpClient struct {
	*http.Client
}

// NewHttpClient - .
func NewHttpClient(cfg *Config) *HttpClient {
	httpClient := getHttpClient(cfg)

	return &HttpClient{
		httpClient,
	}
}

// disallowFollowRedirectFunc - redirect wrapper
func disallowFollowRedirectFunc(req *http.Request, via []*http.Request) error {
	return errRedirectNotAllowed
}

// getHttpClient - Returns *http.Client
func getHttpClient(cfg *Config) *http.Client {
	var transport *http.Transport
	if _, ok := http.DefaultTransport.(*http.Transport); !ok {
		panic(ok)
	} else {
		transport = http.DefaultTransport.(*http.Transport).Clone()
	}

	transport.MaxIdleConns = cfg.MaxIdleConnections
	transport.MaxConnsPerHost = cfg.MaxConnectionsPerHost
	transport.MaxIdleConnsPerHost = cfg.MaxIdleConnections

	var redirectFunc func(req *http.Request, via []*http.Request) error

	if cfg.AllowFollowRedirect {
		redirectFunc = nil
	} else {
		redirectFunc = disallowFollowRedirectFunc
	}

	return &http.Client{
		Timeout:       time.Duration(cfg.Timeout) * time.Second,
		Transport:     transport,
		CheckRedirect: redirectFunc,
	}
}
