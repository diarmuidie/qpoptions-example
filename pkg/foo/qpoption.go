package foo

import (
	"errors"
	"net/http"

	"go.uber.org/zap"
	"knative.dev/serving/pkg/queue/sharedmain"
)

type FooQPOption struct {
	logger               *zap.SugaredLogger
	defaults             *sharedmain.Defaults
	originalRoundTripper http.RoundTripper
}

func NewFooQPOption() *FooQPOption {
	return new(FooQPOption)
}

// RoundTrip implements net/http Roundtrip and modifies the response to add a HTTP header
// NB the request is immutable so cannot be modified, but can be cloned and replaced
func (f *FooQPOption) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	// Catch any panics
	defer func() {
		if r := recover(); r != nil {
			f.logger.Warnf("Recovered from RoundTrip panic: %v\n", r)
			err = errors.New("panic during RoundTrip")
			resp = nil
		}
	}()

	// Call the original round tripper to do the HTTP request
	resp, err = f.originalRoundTripper.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	// Modify the response
	resp.Header["x-foo"] = []string{"bar"}
	f.logger.Info("Foo header added!")

	return resp, nil
}

// Setup sets up the QPOption and overrides the default transport
func (f *FooQPOption) Setup(defaults *sharedmain.Defaults) {
	defer func() {
		if r := recover(); r != nil {
			f.logger.Warnf("Recovered from Setup panic: %v", r)
		}
	}()

	f.defaults = defaults
	f.logger = defaults.Logger
	f.originalRoundTripper = defaults.Transport

	// Replace the default transport with our own round tripper
	defaults.Transport = f
}

// Shutdown performs any shutdown operations
func (f *FooQPOption) Shutdown() {
	defer func() {
		if r := recover(); r != nil {
			f.logger.Warnf("Recovered from Shutdown panic: %v", r)
		}
		f.logger.Sync()
	}()
	// Do any shutdown logic here
}
