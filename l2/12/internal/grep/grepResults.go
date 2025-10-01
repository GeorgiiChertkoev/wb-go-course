package grep

import (
	"fmt"
	"io"
)

// хранит в себе строки которые рядом
// например подряд идущие подходящие
// или подходящее + контекст
type GreppedGroup struct {
	Lines []string
}

type FileGrepResult struct {
	Groups []GreppedGroup
}

func (gg *GreppedGroup) Print(writer io.Writer) {
	for _, l := range gg.Lines {
		fmt.Fprintln(writer, l)
	}
}

func (fg *FileGrepResult) Print(writer io.Writer, groupDelimeter string) {
	for i := range fg.Groups {
		fg.Groups[i].Print(writer)
		if i != len(fg.Groups)-1 {
			fmt.Fprintln(writer, groupDelimeter)
		}
	}
}
