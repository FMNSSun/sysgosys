package fs

import "strings"

func SplitRoot(fpath string) []string {
	parts := strings.Split(fpath, "\\")

	if len(parts) != 2 {
		return nil
	}

	return parts
}

func SplitPath(fpath string) []string {
	parts := strings.Split(fpath, ":")

	return parts
}
