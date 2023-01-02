package main

import (
	"fmt"
	"io"
	"strings"
)

type Options interface{}

type Handler interface {
	Handle(uri string, options Options) (io.ReadCloser, error)
}

type HTTPHandler struct {
	Handler
}

type HTTPOptions struct {
	Headers map[string]string
}

type S3Options struct {
}

func (H *HTTPHandler) Handle(uri string, options Options) (io.ReadCloser, error) {
	httpOpts, ok := options.(HTTPOptions)
	if !ok {
		return nil, fmt.Errorf("Options mismatch")
	}

	fmt.Printf("%v; %v\n", uri, httpOpts)
	return nil, fmt.Errorf("HTTP handler not implemented")
}

var registeredHandlers map[string]Handler = map[string]Handler{
	"http":  &HTTPHandler{},
	"https": &HTTPHandler{},
}

func SmartOpen(uri string, options Options) (io.ReadCloser, error) {
	proto, _, found := strings.Cut(uri, "://")
	if !found {
		// If there's no scheme, treat the input as local path.
		proto = "file"
	}

	handler, ok := registeredHandlers[proto]
	if !ok {
		// TODO: just fallback to os.Open()?
		return nil, fmt.Errorf("No handlers found for URI: %v", uri)
	}

	return handler.Handle(uri, options)
}

func main() {
	// version, err := internal.GetVersion(os.Args[1])
	// if err != nil {
	// 	fmt.Printf("error: %s\n", err)
	// 	return
	// }

	// fmt.Printf("WDL version: '%s'\n", version)

	// cmd.Execute()

	options := HTTPOptions{
		Headers: map[string]string{"Test": "abc"},
	}

	_, err := SmartOpen("https://wlgao.com", options)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
