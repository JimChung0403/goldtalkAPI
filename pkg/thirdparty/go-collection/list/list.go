package list

//List prototype
type List struct {
	elements []interface{}
	size     int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

// New list
func New() *List {
	return &List{}
}

// Add appends a value at the end of the list
func (list *List) Add(values ...interface{}) {
	list.scaleUp(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

//Get the element via the given index
func (list *List) Get(index int) (interface{}, bool) {

	if index < 0 || index >= list.size {
		return nil, false
	}

	return list.elements[index], true
}

//Insert an element
func (list *List) Insert(index int, values ...interface{}) {

	if index < 0 || index >= list.size {
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.scaleUp(l)
	list.size += l

	for i := list.size - 1; i >= index+l; i-- {
		list.elements[i] = list.elements[i-l]
	}

	for i, value := range values {
		list.elements[index+i] = value
	}
}

//Delete element by the given index
func (list *List) Delete(index int) {

	if index < 0 || index >= list.size {
		return
	}

	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--

	list.scaleDown()
}

//Contains return true if list contains the value
func (list *List) Contains(value interface{}) bool {

	found := false
	for _, element := range list.elements {
		if element == value {
			found = true
			break
		}
	}

	return found
}

// Values returns all elements in the list
func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// Empty returns true if list is empty
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns number of elements
func (list *List) Count() int {
	return list.size
}

// Clear removes all elements
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

func (list *List) scaleUp(n int) {

	capacity := cap(list.elements)
	if list.size+n > capacity {
		newcapacity := int(growthFactor * float32(capacity+n))
		list.resize(newcapacity)
	}
}

func (list *List) scaleDown() {

	if shrinkFactor < 0.01 {
		return
	}

	capacity := cap(list.elements)
	if list.size < int(float32(capacity)*shrinkFactor) {
		list.resize(list.size)
	}
}

func (list *List) resize(n int) {
	newElements := make([]interface{}, n, n)
	copy(newElements, list.elements)
	list.elements = newElements
}
