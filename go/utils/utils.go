package utils

import "strings"

// Scanner SplitFunc
func SplitBlankFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

// Returs the maximum of two ints
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// IntSet
type IntSet struct {
	m map[int]bool
}

func NewIntSet(items ...int) *IntSet {
	s := &IntSet{}
	s.m = make(map[int]bool)
	s.Add(items...)
	return s
}

func (s *IntSet) Add(items ...int) {
	for _, item := range items {
		s.m[item] = true
	}
}

func (s *IntSet) Remove(value int) {
	delete(s.m, value)
}

func (s *IntSet) Contains(value int) bool {
	_, c := s.m[value]
	return c
}

func (s *IntSet) Find(f func(int) bool) (int, bool) {
	for item := range s.m {
		if f(item) {
			return item, true
		}
	}

	return 0, false
}

func (s *IntSet) Size() int {
	return len(s.m)
}

// StringSet
type StringSet struct {
	m map[string]bool
}

func NewStringSet(items ...string) *StringSet {
	s := &StringSet{}
	s.m = make(map[string]bool)
	s.Add(items...)
	return s
}

func (s *StringSet) Add(items ...string) {
	for _, item := range items {
		s.m[item] = true
	}
}

func (s *StringSet) Remove(value string) {
	delete(s.m, value)
}

func (s *StringSet) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *StringSet) ContainsAll(items ...string) bool {
	for _, item := range items {
		if _, found := s.m[item]; !found {
			return false
		}
	}
	return true
}

func (s *StringSet) Find(f func(string) bool) (string, bool) {
	for item := range s.m {
		if f(item) {
			return item, true
		}
	}

	return "", false
}

func (s *StringSet) Size() int {
	return len(s.m)
}

// CharSet
type ByteSet struct {
	m map[byte]bool
}

func NewByteSet(items ...byte) *ByteSet {
	s := &ByteSet{}
	s.m = make(map[byte]bool)
	s.Add(items...)
	return s
}

func (s *ByteSet) Add(items ...byte) {
	for _, item := range items {
		s.m[item] = true
	}
}

func (s *ByteSet) Remove(value byte) {
	delete(s.m, value)
}

func (s *ByteSet) Contains(value byte) bool {
	_, c := s.m[value]
	return c
}

func (s *ByteSet) ContainsAll(items ...byte) bool {
	for _, item := range items {
		if _, found := s.m[item]; !found {
			return false
		}
	}
	return true
}

func (s *ByteSet) Find(f func(byte) bool) (byte, bool) {
	for item := range s.m {
		if f(item) {
			return item, true
		}
	}

	return byte(0), false
}

func (s *ByteSet) Size() int {
	return len(s.m)
}
