package golog

import "strings"

func baseFile(absolutePath string) string {
	i := strings.LastIndex(absolutePath, "/")
	if i == -1 {
		return absolutePath
	}
	return absolutePath[i+1:]
}
