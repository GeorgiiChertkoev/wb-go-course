package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

type GrepOptions struct {
	After      int
	Before     int
	CountOnly  bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Pattern    string
	Files      []string
}

func ParseArgs(args []string) (*GrepOptions, error) {
	fs := pflag.NewFlagSet("grep", pflag.ContinueOnError)

	var opts GrepOptions
	fs.IntVarP(&opts.After, "A", "A", 0, "amount of lines to print after match")
	fs.IntVarP(&opts.Before, "B", "B", 0, "amount of lines to print before match")
	contextSize := fs.IntP("C", "C", 0, "amount of lines to print before and after match")
	fs.BoolVarP(&opts.CountOnly, "c", "c", false, "if set will only print number of matches")
	fs.BoolVarP(&opts.IgnoreCase, "i", "i", false, "ignore case when matching")
	fs.BoolVarP(&opts.Fixed, "F", "F", false, "interpret pattern as string not as regex")
	fs.BoolVarP(&opts.LineNum, "n", "n", false, "print number of line")
	fs.BoolVarP(&opts.Invert, "v", "v", false, "invert seatch")

	fs.Parse(args)
	if len(fs.Args()) == 0 {
		return nil, fmt.Errorf("missing pattern for matching")
	}
	opts.Pattern = fs.Args()[0]
	if len(fs.Args()) >= 2 {
		opts.Files = fs.Args()[1:]
	}

	if *contextSize != 0 {
		opts.Before = *contextSize
		opts.After = *contextSize
	}

	return &opts, nil
}
