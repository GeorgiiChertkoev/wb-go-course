package matcher

import "strings"

type SubstringMatcher struct {
	substr string
}

func NewSubstringMatcher(s string) *SubstringMatcher {
	return &SubstringMatcher{substr: s}
}

func (matcher *SubstringMatcher) Match(s string) bool {
	return strings.ContainsAny(s, matcher.substr)
}
