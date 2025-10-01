package matcher

import "go-grep/internal/transformer"

type matcherWithPipeline struct {
	transformer transformer.Transformer
	baseMatcher Matcher
	postMatcher func(string, bool) bool
}

func (m matcherWithPipeline) Match(s string) bool {
	if m.transformer != nil {
		s = m.transformer.Transform(s)
	}
	mathcingResult := m.baseMatcher.Match(s)
	if m.postMatcher != nil {
		return m.postMatcher(s, mathcingResult)
	}
	return mathcingResult
}
