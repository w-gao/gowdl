package internal

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// GetVersion returns the version of the input WDL document.
func GetVersion(uri string) (string, error) {
	f, err := SmartOpen(uri)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	defer f.Close()

	var line string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line = strings.TrimSpace(sc.Text())
		// find the first non-empty, non-comment line in the WDL file
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			break
		}
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if strings.HasPrefix(line, "version ") {
		return strings.Split(line, " ")[1], nil
	}

	// no version string... infer that it is draft-2
	return "draft-2", nil
}

// IsIn checks if str is inside the slice.
func IsIn(str string, slice []string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}

	return false
}

// SmartOpen smartly opens a file based on the URI. Supported URIs are file://,
// http://, and https://.
//
// Read() of the returned io.ReadCloser may raise an error. For example, when
// the connection terminates when streaming from the Internet.
//
// It is also the caller's responsibility to close the returned io.ReadCloser.
func SmartOpen(uri string) (io.ReadCloser, error) {

	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		resp, err := http.DefaultClient.Get(uri)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return resp.Body, nil
	}

	if strings.HasPrefix(uri, "file://") {
		uri = uri[7:]
	}

	// Everything else falls back to os.Open()
	return os.Open(uri) // f, err
}

func ReadString(uri string) (string, error) {
	f, err := SmartOpen(uri)
	if err != nil {
		return "", err
	}

	builder := new(strings.Builder)
	written, err := io.Copy(builder, f)
	if written < 0 || err != nil {
		return "", err
	}

	return builder.String(), nil
}
