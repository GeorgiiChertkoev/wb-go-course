package tests

import (
	"os"
	"strings"
	"testing"

	"unix-sort/internal/args"
	"unix-sort/internal/sorter"

	"github.com/google/go-cmp/cmp"
)

// helper: запускает Sort и возвращает строки
func runSort(t *testing.T, opts args.SortOptions) []string {
	t.Helper()

	tmpOut := t.TempDir() + "/out.txt"
	opts.OutputFile = tmpOut

	if err := sorter.Sort(opts); err != nil {
		t.Fatalf("sort failed: %v", err)
	}

	data, err := os.ReadFile(tmpOut)
	if err != nil {
		t.Fatalf("cannot read output: %v", err)
	}

	// разбиваем по строкам, убираем пустой хвост
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
	return lines
}

func TestSort_StringColumn(t *testing.T) {
	opts := args.ParseArgs(strings.Split(`test_data/fruits.csv -t ","`, " "))
	got := runSort(t, opts)

	want := []string{
		"apple,42,Jan,1K",
		"banana,7,Dec,512",
		"cherry,105,Jul,2M",
		"date,7,Feb,10K",
		"elderberry,3,Oct,3G",
		"fig,1000,Mar,200M",
		"grape,42,Aug,999K",
		"honeydew,9999,Sep,5G",
		"kiwi,15,Apr,123M",
		"lemon,500,May,2K",
		"mango,100,Jun,750K",
		"nectarine,75,Nov,8M",
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("string sort mismatch (-want +got):\n%s", diff)
	}
}

func TestSort_NumericColumn(t *testing.T) {
	opts := args.ParseArgs(strings.Split(`-k 2 -n test_data/fruits.csv -t ","`, " "))

	got := runSort(t, opts)

	want := []string{
		"elderberry,3,Oct,3G",
		"banana,7,Dec,512",
		"date,7,Feb,10K",
		"kiwi,15,Apr,123M",
		"apple,42,Jan,1K",
		"grape,42,Aug,999K",
		"nectarine,75,Nov,8M",
		"mango,100,Jun,750K",
		"cherry,105,Jul,2M",
		"lemon,500,May,2K",
		"fig,1000,Mar,200M",
		"honeydew,9999,Sep,5G",
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("numeric sort mismatch (-want +got):\n%s", diff)
	}
}

func TestSort_MonthColumn(t *testing.T) {
	opts := args.ParseArgs(strings.Split(`test_data/fruits.csv -t "," -Mk3`, " "))
	got := runSort(t, opts)

	// Jan, Feb, Mar...
	want := []string{
		"apple,42,Jan,1K",
		"date,7,Feb,10K",
		"fig,1000,Mar,200M",
		"kiwi,15,Apr,123M",
		"lemon,500,May,2K",
		"mango,100,Jun,750K",
		"cherry,105,Jul,2M",
		"grape,42,Aug,999K",
		"honeydew,9999,Sep,5G",
		"elderberry,3,Oct,3G",
		"nectarine,75,Nov,8M",
		"banana,7,Dec,512",
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("month sort mismatch (-want +got):\n%s", diff)
	}
}

func TestSort_HumanNumericColumn(t *testing.T) {
	opts := args.ParseArgs(strings.Split(`test_data/fruits.csv -t "," -k 4 -h`, " "))

	got := runSort(t, opts)

	want := []string{
		"banana,7,Dec,512",
		"apple,42,Jan,1K",
		"lemon,500,May,2K",
		"date,7,Feb,10K",
		"mango,100,Jun,750K",
		"grape,42,Aug,999K",
		"cherry,105,Jul,2M",
		"nectarine,75,Nov,8M",
		"kiwi,15,Apr,123M",
		"fig,1000,Mar,200M",
		"elderberry,3,Oct,3G",
		"honeydew,9999,Sep,5G",
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("human numeric sort mismatch (-want +got):\n%s", diff)
	}
}

func TestSort_OnlyUniques(t *testing.T) {
	opts := args.ParseArgs(strings.Split(`test_data/fruits.csv test_data/fruits2.csv -u -t "," -k 4 -h`, " "))

	got := runSort(t, opts)

	want := []string{
		"banana,7,Dec,512",
		"apple,42,Jan,1K",
		"lemon,500,May,2K",
		"date,7,Feb,10K",
		"mango,100,Jun,750K",
		"grape,42,Aug,999K",
		"cherry,105,Jul,2M",
		"nectarine,75,Nov,8M",
		"kiwi,15,Apr,123M",
		"fig,1000,Mar,200M",
		"elderberry,3,Oct,3G",
		"honeydew,9999,Sep,5G",
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("human numeric sort mismatch (-want +got):\n%s", diff)
	}
}

func TestSort_RemoveTrailingBlanks(t *testing.T) {
	opts := args.ParseArgs(strings.Split(`test_data/blanks.csv --keep-temps -b -h -k 2 -t ","`, " "))
	got := runSort(t, opts)
	want := []string{
		"1,3M ",
		"3,4M  ",
		"2,1G    ",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("no blanks sort mismatch (-want +got):\n%s", diff)
	}

}
