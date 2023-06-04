package api

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

func NewThrottledDefaultClient(period time.Duration, requests int) *http.Client {
	transport := NewThrottledTransport(period, requests, http.DefaultTransport)
	return &http.Client{
		Transport: transport,
	}
}

type ThrottledTransport struct {
	rt http.RoundTripper
	rl *rate.Limiter
}

func NewThrottledTransport(period time.Duration, requests int, rt http.RoundTripper) http.RoundTripper {
	return &ThrottledTransport{
		rt: rt,
		rl: rate.NewLimiter(rate.Every(period), requests),
	}
}

func (tt *ThrottledTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	err := tt.rl.Wait(r.Context())

	if err != nil {
		return nil, err
	}

	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")

	return tt.rt.RoundTrip(r)
}
