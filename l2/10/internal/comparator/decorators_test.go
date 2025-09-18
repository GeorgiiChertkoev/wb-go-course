package comparator

import "testing"

func TestGetCollumns(t *testing.T) {
	tests := []struct {
		S          string
		Separators string
		Id         int
		Res        string
	}{
		{"a b c d", " ", 2, "b"},
		{"123\t1\t2", "\t", 1, "123"},
		{"123\t1\t2", "\t", 15, ""},
		{"a bcqd\t1", "q\t", 2, "d"},
	}

	for _, tt := range tests {
		field := getCollumn(tt.S, tt.Id, tt.Separators)
		if field != tt.Res {
			t.Errorf("Getting %d collumn from %s (separators=%q), got %s, expected %s", tt.Id, tt.S, tt.Separators, field, tt.Res)
		}
	}
}
