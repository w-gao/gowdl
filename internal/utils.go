package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetVersion gets the version of the input WDL document.
func GetVersion(uri string) (string, error) {
	f, err := os.Open(uri)
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
