package matcher

import "regexp"

type RegexMatcher struct {
	regex *regexp.Regexp
}

func NewRegexMatcher(pattern string) *RegexMatcher {
	return &RegexMatcher{
		regex: regexp.MustCompile(pattern),
	}
}

func (r *RegexMatcher) Match(s string) bool {
	return r.regex.MatchString(s)
}
