package collection

// Comparable defines customized comparable interface for sorting.
type Comparable func(a, b interface{}) int

// Collection shows generic collection structure.
type Collection interface {

	// Size of elements.
	Count() int

	// Gets rid of all elements.
	Clear()

	// Gets all elements.
	Elements() []interface{}
}

// Map defines generic key-value structure.
type Map interface {

	// PUT
	Put(key interface{}, value interface{})

	// GET
	Get(key interface{}) (value interface{}, found bool)

	// Delete
	Delete(key interface{})

	// Returns all keys.
	Keys() []interface{}

	Collection
}

// List defines generic list structure.
type List interface {

	// GET
	Get(index int) (interface{}, bool)

	// DELETE
	Delete(index int)

	// ADD
	Add(values ...interface{})

	// Contains
	Contains(value interface{}) bool

	// INSERTs
	Insert(value interface{}, index int)

	// Sorts
	Sort(comparator Comparable)

	Collection
}
