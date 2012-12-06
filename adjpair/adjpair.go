package adjpair

import "bytes"
import "fmt"
import "regexp"
import "strings"

const PAIR_ACTIVE = 1
const PAIR_INACTIVE = 2

// maximum elements in a filepath
const FILEPATH_MAX = 64

type Pair struct {
  fst byte
  snd byte
  status byte
  __align byte
}

type Pairs []Pair

func (pair *Pair) String() string {
  return fmt.Sprintf("%c%c|%d ", pair.fst, pair.snd, pair.status)
}

func detectLength(tokens []string) int {
  result := 0
  for i := 0; i < len(tokens); i++ {
    l := len(tokens[i]) - 1
    if l > 0 {
      result += l
    }
  }
  return result
}

// This helper func is missing in Go 1.0 (String type)
func splitWithRegexp(s string, re *regexp.Regexp) []string {
  if len(re.String()) > 0 && len(s) == 0 {
    return []string{""}
  }
  matches := re.FindAllStringIndex(s, -1)
  strings := make([]string, 0, len(matches))
  beg := 0
  end := 0
  for _, match := range matches {
    end = match[0]
    if match[1] != 0 {
      strings = append(strings, s[beg:end])
    }
    beg = match[1]
  }
  if end != len(s) {
    strings = append(strings, s[beg:])
  }
  return strings
}

func splitFilepath(s string) []string {
  if len(s) == 0 {
    return []string{}
  }
  strs := make([]string, 0, FILEPATH_MAX)
  slc := s
  ix := 0
  for {
    ix = strings.IndexRune(slc, '/')
    if ix == -1 {
      break
    }
    strs = append(strs, slc[0:ix])
    slc = slc[ix + 1:]
  }
  strs = append(strs, slc[ix + 1:])
  return strs
}

func NewPairsFromArray(tokens []string) Pairs {
  dl := detectLength(tokens)
  pairs := make(Pairs, dl)

  k := 0
  for i := 0; i < len(tokens); i++ {
    t := tokens[i]
    for j := 0; j < len(t) - 1; j++ {
      pairs[k].fst = t[j];
      pairs[k].snd = t[j + 1];
      pairs[k].status = PAIR_ACTIVE;
      k++;
    }
  }
  return pairs
}

func NewPairsFromString(str string) Pairs {
  return NewPairsFromArray([]string{str})
}

func NewPairsFromStringTokens(str string, re regexp.Regexp) Pairs {
  ss := splitWithRegexp(str, &re)
  return NewPairsFromArray(ss)
}

func NewPairsFromFilepath(str string) Pairs {
  ss := splitFilepath(str)
  return NewPairsFromArray(ss)
}

func (self Pairs) String() string {
  buff := bytes.NewBufferString("")
  for _, p := range self {
    buff.WriteString(p.String())
  }
  return buff.String()
}

func (self Pairs) Reactivate() {
  for _, p := range self {
    p.status = PAIR_ACTIVE
  }
}

func (a Pair) Equal(b Pair) bool {
  return (a.fst == b.fst && a.snd == b.snd && 
    (a.status & b.status & PAIR_ACTIVE) == PAIR_ACTIVE)
}

func (self Pairs) Match(other Pairs) float64 {
  matches := 0
  len_self := len(self)
  len_other := len(other)
  sum := len_self + len_other
  if sum == 0 {
    return 1.0
  }
  for i := 0; i < len_self; i++ {
    for j := 0; j < len_other; j++ {
      if self[i].Equal(other[j]) {
        matches++
        other[j].status = PAIR_INACTIVE
        break
      }
    }
  }
  return float64(2 * matches) / float64(sum)
}

func MatchStrings(stra, strb string) float64 {
  a := NewPairsFromString(stra)
  b := NewPairsFromString(strb)
  return a.Match(b)
}

func MatchStringsTokens(stra, strb string, re *regexp.Regexp) float64 {
  a := NewPairsFromStringTokens(stra, *re)
  b := NewPairsFromStringTokens(strb, *re)
  return a.Match(b)
}

