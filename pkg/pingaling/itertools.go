// Copyright © 2018 The Pingaling Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pingaling

// Iter is channel of interface{}
type Iter chan interface{}

// Predicate is filter function
type Predicate func(interface{}) bool

// Mapper map Iter to a function
type Mapper func(interface{}) interface{}

// New returns a channel of interface
func New(els ...interface{}) Iter {
	c := make(Iter)
	go func() {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

// List convert Iter to Interface{}
func List(it Iter) []interface{} {
	arr := make([]interface{}, 0, 1)
	for el := range it {
		arr = append(arr, el)
	}
	return arr
}

// ListIter conver interface{} to Iter
func ListIter(i []interface{}) Iter {
	c := make(Iter)
	go func() {
		for _, el := range i {
			c <- el
		}
		close(c)
	}()
	return c
}

// StrIter returns channels of string
func StrIter(els []string) Iter {
	c := make(Iter)
	go func() {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

// Map an iterator to fn(el) for el in it
func Map(fn Mapper, it Iter) Iter {
	c := make(Iter)
	go func() {
		for el := range it {
			c <- fn(el)
		}
		close(c)
	}()
	return c
}

// Filter out any elements where pred(el) == false
func Filter(pred Predicate, it Iter) Iter {
	c := make(Iter)
	go func() {
		for el := range it {
			if keep := pred(el); keep {
				c <- el
			}
		}
		close(c)
	}()
	return c
}
