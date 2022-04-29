package services

import (
	"fmt"
	"strings"
)

type Url struct {
	Base string
	Path string
	Full string
}

func Parse(base string, path string) Url {
	if !strings.Contains(base, "http") {
		return Url{
			fmt.Sprintf("https://%s", base),
			path,
			fmt.Sprintf("https://%s%s", base, path),
		}
	}

	return Url{
		base,
		path,
		fmt.Sprintf("%s%s", base, path),
	}
}
