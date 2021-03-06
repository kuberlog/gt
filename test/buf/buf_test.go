package test

import (
	"github.com/kuberlog/gt/buf"
	"strings"
	"testing"
)

var str = "this\nis\na buffer\n"

func TestCreateAndView(t *testing.T) {
	b := buf.FromString(str)
	actual := b.View()
	expected := []string{"this", "is", "a buffer"}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Fail()
		}
	}
}

func TestViewWithIndices(t *testing.T) {
	b := buf.FromString(str)
	actual := b.ViewLines(0, 2)
	if len(actual) != 2 || "this" != actual[0] || "is" != actual[1] {
		t.Errorf("actual: %s", actual)
	}
}

func TestMarkAndGetMarkers(t *testing.T) {
	b := buf.FromString(str)
	expected := make([]buf.Marker, 0)
	for index := range strings.Split(str, "\n") {
		expected = append(expected, b.Mark(index, 0))
	}
	actual := b.GetMarkers()
	if len(actual) != len(expected) {
		t.Errorf("len(actual)=%v, len(expected)=%v", len(actual), len(expected))
		return
	}
	for index := range expected {
		if actual[index] != expected[index] {
			t.Errorf("%v != %v", actual[index], expected[index])
		}
	}
}

func TestDeleteMarker(t *testing.T) {
	b := buf.FromString(str)
	b.Mark(0, 0)
	mark := b.Mark(1, 0)
	b.DeleteMarker(mark)
	if len(b.GetMarkers()) != 1 {
		t.Errorf("len(b.GetMarkers) != 1")
	}
}

func TestGetLineByMarker(t *testing.T) {
	b := buf.FromString(str)
	marker := b.Mark(1, 0)
	line := b.GetLineByMarker(marker)
	if line != "is" {
		t.Errorf("%s != %s", "is", line)
	}
}

func TestDeleteLineByMarker_MiddleLine(t *testing.T) {
	b := buf.FromString(str)
	marker := b.Mark(1, 0)

	b.DeleteLineByMarker(marker)
	lines := b.View()
	strSlice := strings.Split(str, "\n")
	expected := append(strSlice[:1], strSlice[2:]...)
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("%s != %s", line, expected[i])
		}
	}
}

func TestDeleteLineByMarker_MarkerRowToBig(t *testing.T) {
	b := buf.FromString(str)
	marker := b.Mark(10, 0)

	b.DeleteLineByMarker(marker)
	lines := b.View()

	// Should do nothing
	strSlice := strings.Split(str, "\n")
	for i, line := range lines {
		if line != strSlice[i] {
			t.Errorf("%s != %s", line, strSlice[i])
		}
	}
}

func TestDeleteLineByMarker_MarkerRowIsZero(t *testing.T) {
	b := buf.FromString(str)
	marker := b.Mark(0, 0)

	b.DeleteLineByMarker(marker)
	lines := b.View()

	// Should do nothing
	strSlice := strings.Split(str, "\n")
	expected := strSlice[1:]
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("%s != %s", line, expected[i])
		}
	}
}
