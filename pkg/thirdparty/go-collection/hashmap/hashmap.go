package hashmap

// HashMap keeps the native structure.
type HashMap struct {
	m map[interface{}]interface{}
}

// New creates new one
func New() *HashMap {
	return &HashMap{m: make(map[interface{}]interface{})}
}

// Put key-value
func (m *HashMap) Put(key interface{}, value interface{}) {
	m.m[key] = value
}

// Get value by key
func (m *HashMap) Get(key interface{}) (value interface{}, found bool) {
	value, found = m.m[key]
	return
}

// Delete element
func (m *HashMap) Delete(key interface{}) {
	delete(m.m, key)
}

// Count elements
func (m *HashMap) Count() int {
	return len(m.m)
}

// Keys returns all keys.
func (m *HashMap) Keys() []interface{} {
	keys := make([]interface{}, m.Count())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

// Values returns all values.
func (m *HashMap) Values() []interface{} {
	values := make([]interface{}, m.Count())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

// Clear all elements
func (m *HashMap) Clear() {
	m.m = make(map[interface{}]interface{})
}

// Elements returns element one by one
func (m *HashMap) Elements() []interface{} {

	return m.Values()
}
