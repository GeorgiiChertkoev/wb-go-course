package grep

import (
	"go-grep/internal/matcher"
	"go-grep/internal/options"
	"os"
	"strconv"
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
		grepRes, err := grepper.grepFile(filename)
		if err != nil {
			return nil, err
		}
		if opts.CountOnly {
			size := 0
			for i := range grepRes.Groups {
				size += len(grepRes.Groups[i].Lines)
			}
			grepRes = &FileGrepResult{
				Groups: []GreppedGroup{
					{Lines: []string{strconv.Itoa(size)}},
				},
			}
		}
		result = append(result, grepRes)
	}

	return result, nil
}
