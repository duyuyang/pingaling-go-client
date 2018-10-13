package pingaling

// Iter is channel of interface{}
type Iter chan interface{}

// Predicate is filter function
type Predicate func(interface{}) bool

// Mapper map Iter to a function
type Mapper func(interface{}) interface{}

// New returns a channel of interface
func New(els []interface{}) Iter {
	c := make(Iter)
	go func() {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

// StringToInterface converts list of string to interface
func StringToInterface(s []string) (r []interface{}) {
	r = make([]interface{}, len(s))
	for i, v := range s {
		r[i] = v
	}
	return
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

// MapEx Execute Mapper
func MapEx(fn Mapper, it Iter) {
	for el := range it {
		fn(el)
	}
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
