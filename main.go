package IntSet

import (
	"bytes"
	"fmt"
)

const x86 = 32 << (uint(0) >> 63)

type IntSet struct{ words []uint }

// Has указывает, содержит ли множество неотрицательное значение х.
func (s *IntSet) Has(x int) bool {
	word, bit := x/x86, uint(x%x86)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество.
func (s *IntSet) Add(x int) {
	word, bit := x/x86, uint(x%x86)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll добавляет неотрицательные значения ~nums~ в множество.
func (s *IntSet) AddAll(nums ...int) {
	for _, x := range nums {
		s.Add(x)
	}
}

// UnionWith делает множество s равным объединению множеств s и t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith делает множество s равным пересечению множеств  s и t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// SymmetricDifference делает множество s равным семетричной разности множеств s и t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith делает множество s равным разности множеств s и t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// String возвращает множество как строку вида "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < x86; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", x86*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len возвращает колличество элементов в множестве s.
func (s *IntSet) Len() int {
	var count int
	for _, x := range s.words {
		for x != 0 {
			x = x & (x - 1)
			count++
		}
	}
	return count
}

// Remove удаляет число из множества s, возвращает true если число было удалено.
func (s *IntSet) Remove(x int) bool {
	word, bit := x/x86, uint(x%x86)
	if s.words[word]&(1<<bit) != 0 {
		s.words[word] ^= 1 << bit
		return true
	}
	return false
}

// Clear отчищает множество s.
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy возвращает копию множества s
func (s *IntSet) Copy() *IntSet {
	words := make([]uint, len(s.words))
	copy(words, s.words)
	return &IntSet{words: words}
}

// Elems возвращает срез состоящий из элементов множества s.
func (s *IntSet) Elems() []int {
	var res []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < x86; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, i*x86+j)
			}
		}
	}
	return res
}
