package grep

import (
	"bufio"
	"container/list"
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
	linesBefore := list.New()
	contextAfter := 0 // counter for lines after context
	for scanner.Scan() {
		line = scanner.Text()
		if g.matcher.Match(line) {
			// pushing prefs
			for linesBefore.Len() > 0 {
				line := linesBefore.Front().Value.(string)
				curGroup.Lines = append(curGroup.Lines, g.formatLine(line, lineNum-linesBefore.Len()))
				linesBefore.Remove(linesBefore.Front())
			}
			line = g.formatLine(line, lineNum)
			curGroup.Lines = append(curGroup.Lines, line)
			contextAfter = g.opts.After
		} else if contextAfter > 0 {
			curGroup.Lines = append(curGroup.Lines, g.formatLine(line, lineNum))
			contextAfter--
		} else {
			if len(curGroup.Lines) != 0 {
				groups = append(groups, curGroup)
				curGroup.Lines = make([]string, 0)
			}
			linesBefore.PushBack(line)
			if linesBefore.Len() > g.opts.Before {
				linesBefore.Remove(linesBefore.Front())
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
	res, err := g.grepReader(file)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *grepper) formatLine(line string, lineNum int) string {
	if !g.opts.LineNum {
		return line
	}
	delimeter := "-"
	if g.matcher.Match(line) {
		delimeter = ":"
	}
	return fmt.Sprintf("%d%s%s", lineNum, delimeter, line)

}
