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
