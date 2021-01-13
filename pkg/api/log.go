package pkg

import (
	format "github.com/celikelozdinc/formatter"
)

// Fmt invokes the Format method of formatter package
func Fmt(input string) string {
	return format.Format(input)
}
