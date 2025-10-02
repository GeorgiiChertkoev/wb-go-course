package grep

import (
	"fmt"
	"go-grep/internal/matcher"
	"go-grep/internal/options"
	"os"
)

type grepper struct {
	opts    options.GrepOptions
	matcher matcher.Matcher
}

func Grep(opts options.GrepOptions) ([]*FileGrepResult, error) {
	matcher := matcher.BuildMatcher(opts)
	grepper := grepper{
		opts:    opts,
		matcher: matcher,
	}

	if len(opts.Files) == 0 {
		res, err := grepper.grepReader(os.Stdin)
		if err != nil {
			return nil, err
		}
		return []*FileGrepResult{res}, nil
	}
	result := make([]*FileGrepResult, 0, len(opts.Files))
	for _, filename := range opts.Files {
		fmt.Printf("\n\ngonna grep %s\n\n", filename)
		grepRes, err := grepper.grepFile(filename)
		if err != nil {
			return nil, err
		}
		result = append(result, grepRes)
	}

	return result, nil
}
