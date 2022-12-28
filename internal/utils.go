package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetVersion(filename string) (string, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var line string
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

	// no version string... infer that it is draft-2
	if !strings.HasPrefix(line, "version") {
		return "draft-2", nil
	}

	return strings.Split(line, "version ")[1], nil
}
