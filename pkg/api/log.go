package pkg

import (
	formatv1 "github.com/celikelozdinc/formatter"
	formatv2 "github.com/celikelozdinc/formatter/v2"
)

// FmtWithv1 invokes the Format method of formatter package
func FmtWithv1(input string) string {
	return formatv1.Format(input)
}

// FmtWithv2 invokes the Format method of formatter package
func FmtWithv2(input string) string {
	return formatv2.Format(input)
}
