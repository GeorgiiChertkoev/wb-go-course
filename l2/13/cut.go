package main

import (
	"bufio"
	"os"
	"sync"
)

func Cut(opts CutOpts) error {
	input := make(chan string, pipelineChanSize)
	filteredFields := buildPipeline(opts, input)

	var wg sync.WaitGroup
	wg.Go(func() {
		PrintFields(filteredFields, opts.Delimeter, opts.Output)
	})

	if len(opts.InputFiles) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input <- scanner.Text()
		}
	}

	for _, filename := range opts.InputFiles {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			input <- scanner.Text()
		}
		f.Close()
	}
	close(input)
	wg.Wait()
	return nil
}

func buildPipeline(opts CutOpts, input <-chan string) <-chan []string {
	var filtered <-chan string
	if opts.OnlySeparated {
		filtered = Filter(input, opts.Delimeter)
	} else {
		filtered = input
	}

	splited := SplitIntoFields(filtered, opts.Delimeter)
	filteredFields := FilterFields(splited, opts.Fields)
	return filteredFields
}
