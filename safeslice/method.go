package safeslice

import (
	"fmt"
	"math/rand"
)

func (s Slice[T]) Len() int {
	return len(s)
}

func (s *Slice[T]) Copy() *Slice[T] {
	copy := append((*s)[:0:0], *s...)
	return &copy
}

func (s *Slice[T]) Clear() *Slice[T] {
	*s = (*s)[:0]
	return s
}

func (s *Slice[T]) Extend(size int) *Slice[T] {
	if size < 0 {
		*s = (*s)[:max(len(*s)+size, 0)]
	} else {
		*s = append(*s, make([]T, size)...)
	}
	return s
}

func (s *Slice[T]) Shrink(size int) *Slice[T] {
	return s.Extend(-size)
}

func (s *Slice[T]) Set(i int, v T) {
	if len(*s) == 0 {
		return
	}
	handleIndexException(&i, len(*s))
	(*s)[i] = v
}

func (s *Slice[T]) Get(i int) T {
	var v T
	if len(*s) == 0 {
		return v
	}
	handleIndexException(&i, len(*s))
	v = (*s)[i]
	return v
}

func (s *Slice[T]) Slice(i, j int) *Slice[T] {
	handleIndexException(&i, len(*s))
	handleIndexException(&j, len(*s))
	if i > j {
		i, j = j, i
	}
	fmt.Println(i, j)

	*s = (*s)[i:j]
	return s
}

func (s *Slice[T]) Splice(i, j int, v ...T) *Slice[T] {
	//preventIndexException(&i, &j, len(*s))
	*s = append((*s)[:i], append(v, (*s)[j:]...)...)
	return s
}

func (s *Slice[T]) Expand(i int, v ...T) *Slice[T] {
	return s.Splice(i, i, v...)
}

func (s *Slice[T]) Insert(i int, v T) *Slice[T] {
	return s.Expand(i, v)
}

func (s *Slice[T]) Cut(i, j int) *Slice[T] {
	return s.Splice(i, j)
}

func (s *Slice[T]) Delete(i int) *Slice[T] {
	return s.Cut(i, i+1)
}

func (s *Slice[T]) DeleteReplace(i int) *Slice[T] {
	//preventIndexException(&i, nil, len(*s))
	(*s)[i] = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return s
}

func (s *Slice[T]) Push(v T) *Slice[T] {
	return s.Insert(len(*s), v)
}

func (s *Slice[T]) Pop() T {
	var r T
	if len(*s) == 0 {
		return r
	}
	r = (*s)[len(*s)-1]
	s.Delete(len(*s) - 1)
	return r
}

func (s *Slice[T]) UnShift(v T) *Slice[T] {
	return s.Insert(0, v)
}

func (s *Slice[T]) Shift() T {
	var r T
	if len(*s) == 0 {
		return r
	}
	r = (*s)[0]
	s.Delete(0)
	return r
}

func (s *Slice[T]) Reverse() *Slice[T] {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
	return s
}

func (s *Slice[T]) Rotate(n int) *Slice[T] {
	if len(*s) == 0 {
		return s
	}

	if n > 0 {
		index := n % len(*s)
		*s = append((*s)[index:], (*s)[:index]...)
	} else if n < 0 {
		index := -n % len(*s)
		*s = append((*s)[len(*s)-index:], (*s)[:len(*s)-index]...)
	}
	return s
}

func (s *Slice[T]) Shuffle() *Slice[T] {
	rand.Shuffle(len(*s), func(i, j int) {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	})
	return s
}

func (s *Slice[T]) Filter(f func(T) bool) *Slice[T] {
	n := 0
	for _, v := range *s {
		if f(v) {
			(*s)[n] = v
			n++
		}
	}
	*s = (*s)[:n]
	return s
}

func (s *Slice[T]) Repeat(cnt int) *Slice[T] {
	repeat := make(Slice[T], 0, len(*s)*cnt)
	for i := 0; i < cnt; i++ {
		repeat = append(repeat, *s...)
	}
	*s = repeat
	return s
}

func (s Slice[T]) Split(splits ...int) []Slice[T] {
	var r []Slice[T] = make([]Slice[T], 0, len(splits)+1)
	var end int
	for i := 0; i < len(splits); i++ {
		endNext := end + splits[i]
		//preventIndexException(&endNext, nil, len(s))
		if endNext < end {
			r = append(r, s[endNext:end])
		} else {
			r = append(r, s[end:endNext])
		}
		end = endNext
	}
	if end < len(s) {
		r = append(r, s[end:])
	}
	return r
}

func (s Slice[T]) Batch(size int) []Slice[T] {
	if size <= 0 {
		size = 1
	}
	splits := make([]int, (len(s))/size)
	for i := 0; i < len(splits); i++ {
		splits[i] = size
	}
	return s.Split(splits...)
}

func (s Slice[T]) Join(sep string) string {
	var r string
	for i, v := range s {
		if i != 0 {
			r += sep
		}
		r += fmt.Sprintf("%v", v)
	}
	return r
}

func (s Slice[T]) IndexOf(f func(T) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

func (s Slice[T]) Contains(f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func (s Slice[T]) Each(f func(T)) {
	for _, v := range s {
		f(v)
	}
}

func (s Slice[T]) ForEach(f func(int, T)) {
	for i, v := range s {
		f(i, v)
	}
}

func handleIndexException(idx *int, len int) {
	if len <= 0 {
		*idx = 0
		return
	}
	*idx = (*idx%len + len) % len
}
