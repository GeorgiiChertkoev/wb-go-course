package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

type CutOpts struct {
	InputFiles    []string
	Fields        []int
	Delimeter     string
	OnlySeparated bool
	Output        io.WriteCloser
}

func ParseArgs(args []string) (CutOpts, error) {
	fs := pflag.NewFlagSet("cut", pflag.ContinueOnError)
	var opts CutOpts
	var fields string
	fs.StringVarP(&fields, "fields", "f", "", "column numbers that will be print")
	fs.StringVarP(&opts.Delimeter, "delimiter", "d", "\t", "delimeter used for spliting columns")
	fs.BoolVarP(&opts.OnlySeparated, "separated", "s", false, "only print lines with delimeter")
	fs.Parse(args)
	f, err := parseFields(fields)
	if err != nil {
		return opts, nil
	}
	opts.Fields = f
	opts.InputFiles = fs.Args()
	opts.Output = os.Stdout
	return opts, nil
}

// example: 1,3-5 -> {1, 3, 4, 5}
func parseFields(s string) ([]int, error) {
	// check that string contains only letters or ',' or '-'
	// to safely ignore Atoi errors
	regex := regexp.MustCompile("^[0-9,-]*$")
	if !regex.MatchString(s) {
		return nil, fmt.Errorf("fields value contains forbidden cymbols")
	}

	if s == "" {
		return nil, nil
	}

	groups := strings.Split(s, ",")
	var res []int

	for _, g := range groups {
		if !strings.Contains(g, "-") {
			n, _ := strconv.Atoi(g)
			res = append(res, n)
			continue
		}
		if strings.Count(g, "-") > 1 {
			return nil, fmt.Errorf("fileds value can't have more than 1 '-'")
		}
		before, after, _ := strings.Cut(g, "-")
		first, _ := strconv.Atoi(before)
		last, _ := strconv.Atoi(after)
		if first > last {
			return nil, fmt.Errorf("fields value has invalid range: %d-%d", first, last)
		}
		for i := first; i < last+1; i++ {
			res = append(res, i)
		}
	}
	return res, nil
}
