package grep

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func (g *grepper) grepReader(reader io.Reader) (*FileGrepResult, error) {
	scanner := bufio.NewScanner(reader)
	lineNum := 0
	var groups []GreppedGroup
	var curGroup GreppedGroup
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Printf("scanned: %s\n", scanner.Text())
		if g.matcher.Match(line) {
			// push prefs
			if g.opts.LineNum {
				line = fmt.Sprintf("%d:%s", lineNum, line)
			}
			curGroup.Lines = append(curGroup.Lines, line)
		} else {
			if len(curGroup.Lines) != 0 {
				groups = append(groups, curGroup)
				curGroup.Lines = make([]string, 0)
			}
		}
		lineNum++
	}
	if len(curGroup.Lines) != 0 {
		groups = append(groups, curGroup)
	}
	return &FileGrepResult{
		Groups: groups,
	}, nil
}

func (g *grepper) grepFile(filename string) (*FileGrepResult, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return g.grepReader(file)
}
