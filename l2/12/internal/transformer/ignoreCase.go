package transformer

import "strings"

type LowerCaseTransformer struct{}

func (i LowerCaseTransformer) Transform(s string) string {
	return strings.ToLower(s)
}
