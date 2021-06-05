package list

import "testing"

func TestListAdd(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Count(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListRemove(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	list.Delete(2)
	if actualValue, ok := list.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Delete(1)
	list.Delete(0)
	list.Delete(0) // no effect
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Count(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListGet(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(3); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Delete(0)
	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
}
