package net

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"time"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64; rv:127.0) Gecko/20100101 Firefox/127.0",
}

type StealthTransport struct {
	Base http.RoundTripper
}

func (s *StealthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())
	r.Header.Del("X-Forwarded-For")
	r.Header.Del("Via")
	r.Header.Del("X-Real-IP")

	if r.Header.Get("User-Agent") == "" {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Header.Set("User-Agent", userAgents[rnd.Intn(len(userAgents))])
	}

	tr := s.Base
	if tr == nil {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}
	}
	return tr.RoundTrip(r)
}

func NewStealthClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &StealthTransport{
			Base: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
			},
		},
	}
}
