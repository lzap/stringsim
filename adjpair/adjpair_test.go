package adjpair

import "testing"
import "math"
import "regexp"
import "fmt"

var examplesSplitFilepath = []struct {
	path   string
	result []string
}{
	{"", []string{}},
	{"test", []string{"test"}},
	{"test/foo", []string{"test", "foo"}},
	{"/test/foo", []string{"", "test", "foo"}},
	{"/test/foo/", []string{"", "test", "foo", ""}},
	{"/test foo/", []string{"", "test", "foo", ""}},
	{"/test  foo/", []string{"", "test", "", "foo", ""}},
}

func TestSplitFilepath(t *testing.T) {
	for i, tt := range examplesSplitFilepath {
		result := splitFilepath(tt.path)
		str_result := fmt.Sprintf("%v", result)
		str_expected := fmt.Sprintf("%v", tt.result)
		if str_result != str_expected {
			t.Errorf("%d. path '%s' returned %s, but expected %s", i, tt.path, str_result, str_expected)
		}
	}
}

func TestDetectLength(t *testing.T) {
	e := detectLength([]string{"abc", "def", "ghi"})
	if e != 6 {
		t.Fail()
	}
}

func TestNewPairsEmpty(t *testing.T) {
	p := NewPairsFromArray([]string{})
	if len(p) != 0 {
		t.Fail()
	}
}

func TestNewPairsEmptyString(t *testing.T) {
	p := NewPairsFromArray([]string{""})
	if len(p) != 0 {
		t.Fail()
	}
}

func TestNewPairsEmptyStrings(t *testing.T) {
	p := NewPairsFromArray([]string{"", ""})
	if len(p) != 0 {
		t.Fail()
	}
}

func TestNewPairs6(t *testing.T) {
	p := NewPairsFromArray([]string{"abc", "def", "ghi"})
	if len(p) != 6 {
		t.Fail()
	}
}

var examplesMatchStrings = []struct {
	a   string
	b   string
	exp float64
}{
	{"", "", 1.0},
	{"test", "test", 1.0},
	{"", "test", 0.0},
	{"test", "testa", 0.8571428},
	{"testa", "test", 0.8571428},
	{"test", "atest", 0.8571428},
	{"test", "est", 0.8},
	{"test", "tst", 0.4},
	{"test", "taex", 0.0},
}

func TestMatchStrings(t *testing.T) {
	for i, tt := range examplesMatchStrings {
		result := MatchStrings(tt.a, tt.b)
		if math.Abs(result-tt.exp) > 0.00005 {
			t.Errorf("%d. match '%s' and '%s' returned %f, but expected %f", i, tt.a, tt.b, result, tt.exp)
		}
	}
}

var examplesMatchStringsTokens = []struct {
	a   string
	b   string
	exp float64
}{
	{"", "", 1.0},
	{"foo,bar,baz", "foo,bar,baz", 1.0},
	{"foo,bar,baz", "foo", 0.5},
	{"foo,bar,baz", "foo,bar", 0.8},
	{"foo,bar,baz", "foo,baz,bar", 1.0},
}

func TestMatchStringTokens(t *testing.T) {
	re, _ := regexp.Compile(",")
	for i, tt := range examplesMatchStringsTokens {
		result := MatchStringsTokens(tt.a, tt.b, re)
		if math.Abs(result-tt.exp) > 0.00005 {
			t.Errorf("%d. match '%s' and '%s' returned %f, but expected %f", i, tt.a, tt.b, result, tt.exp)
		}
	}
}
