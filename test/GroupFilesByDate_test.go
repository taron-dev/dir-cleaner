package dir_cleaner

import (
	"dir_cleaner"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func Test_GroupFilesByDate_MultipleFiles_GroupedCorrectly(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2020-01-28")
	t2, _ := time.Parse("2006-01-02", "2020-01-29")

	f1 := FileConstructor(t1, "f1.png", 1, 0755, false)
	f2 := FileConstructor(t1, "f2.png", 1, 0755, false)
	f3 := FileConstructor(t2, "f3.png", 1, 0755, false)

	files := []os.FileInfo{f1, f2, f3}

	actualDatesMap := dir_cleaner.GroupFilesByDate(files)

	if len(actualDatesMap) != 2 {
		t.Errorf("got %d, wanted %d", len(actualDatesMap), 2)
	}

	if len(actualDatesMap[t1]) != 2 {
		t.Errorf("got %d, wanted %d", len(actualDatesMap[t1]), 2)
	}

	if len(actualDatesMap[t2]) != 1 {
		t.Errorf("got %d, wanted %d", len(actualDatesMap[t2]), 1)
	}

	actualF1 := actualDatesMap[t1][0]
	actualF2 := actualDatesMap[t1][1]
	assert.Equal(t, actualF1.Name(), f1.name, "The two words should be the same.")
	assert.Equal(t, actualF2.Name(), f2.name, "The two words should be the same.")

	actualF3 := actualDatesMap[t2][0]
	assert.Equal(t, actualF3.Name(), f3.name, "The two words should be the same.")
}
