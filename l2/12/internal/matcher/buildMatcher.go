package matcher

import (
	"go-grep/internal/options"
	"go-grep/internal/transformer"
	"strings"
)

func BuildMatcher(opts options.GrepOptions) Matcher {
	var matcher matcherWithPipeline
	if opts.IgnoreCase {
		opts.Pattern = strings.ToLower(opts.Pattern)
		matcher.transformer = transformer.LowerCaseTransformer{}
	}
	if opts.Fixed {
		matcher.baseMatcher = NewSubstringMatcher(opts.Pattern)
	} else {
		matcher.baseMatcher = NewRegexMatcher(opts.Pattern)
	}

	if opts.Invert {
		matcher.postMatcher = func(s string, matchingResult bool) bool {
			return !matchingResult
		}
	}

	return matcher
}
