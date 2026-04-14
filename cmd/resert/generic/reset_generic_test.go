package generic

import (
	"testing"

	"github.com/FedorSidorow/shortener/cmd/resert/testpkg"
)

type TestStruct struct {
	ID   int
	Name string
	Data []byte
}

func (t *TestStruct) Reset() {
	t.ID = 0
	t.Name = ""
	if t.Data != nil {
		t.Data = t.Data[:0]
	}
}

type TestSliceStruct struct {
	Items []int
	Count int
}

func (t *TestSliceStruct) Reset() {
	if t.Items != nil {
		t.Items = t.Items[:0]
	}
	t.Count = 0
}

type TestArrayStruct struct {
	Values [3]int
	Tag    string
}

func (t *TestArrayStruct) Reset() {
	t.Values = [3]int{}
	t.Tag = ""
}

func TestNewPool(t *testing.T) {

	ts := &TestStruct{
		ID:   42,
		Name: "test",
		Data: []byte{1, 2, 3},
	}

	pool := New(ts)

	if pool == nil {
		t.Fatal("New() returned nil pool")
	}

	item := pool.Get()
	if item != ts {
		t.Errorf("Get() returned different item, got %v, want %v", item, ts)
	}

	tss := &TestSliceStruct{
		Items: []int{1, 2, 3},
		Count: 3,
	}

	pool2 := New(tss)
	if pool2 == nil {
		t.Fatal("New() returned nil pool for slice struct")
	}

	item2 := pool2.Get()
	if item2 != tss {
		t.Errorf("Get() returned different item for slice struct, got %v, want %v", item2, tss)
	}

	tas := &TestArrayStruct{
		Values: [3]int{4, 5, 6},
		Tag:    "array",
	}

	pool3 := New(tas)
	if pool3 == nil {
		t.Fatal("New() returned nil pool for array struct")
	}

	item3 := pool3.Get()
	if item3 != tas {
		t.Errorf("Get() returned different item for array struct, got %v, want %v", item3, tas)
	}
}

func TestPoolGet(t *testing.T) {
	ts := &TestStruct{
		ID:   100,
		Name: "get_test",
		Data: []byte{4, 5, 6},
	}

	pool := New(ts)

	item1 := pool.Get()
	// After first Get, pool should be empty
	var zero *TestStruct
	item2 := pool.Get()
	if item2 != zero {
		t.Errorf("Second Get() should return zero value when pool is empty, got %v", item2)
	}

	// item1 should be the original item
	if item1 != ts {
		t.Errorf("Get() returned different item, got %v, want %v", item1, ts)
	}

	if item1.ID != 100 || item1.Name != "get_test" || len(item1.Data) != 3 {
		t.Errorf("Get() returned item with incorrect values: %+v", item1)
	}
}

func TestPoolPut(t *testing.T) {
	ts := &TestStruct{
		ID:   200,
		Name: "put_test",
		Data: []byte{7, 8, 9},
	}

	pool := New(ts)

	newItem := &TestStruct{
		ID:   300,
		Name: "new_item",
		Data: []byte{10, 11, 12},
	}

	pool.Put(newItem)

	if newItem.ID != 0 || newItem.Name != "" || len(newItem.Data) != 0 {
		t.Errorf("Put() did not call Reset() properly, got: ID=%d, Name=%s, Data=%v",
			newItem.ID, newItem.Name, newItem.Data)
	}

	// Get should return the most recently put item (LIFO)
	poolItem := pool.Get()
	if poolItem != newItem {
		t.Errorf("Get() should return the put item, got %v, want %v", poolItem, newItem)
	}

	// The put item was reset, so its fields should be zero
	if poolItem.ID != 0 || poolItem.Name != "" || len(poolItem.Data) != 0 {
		t.Errorf("Put item should be reset when retrieved, got: %+v", poolItem)
	}

	// Next Get should return the original item
	poolItem2 := pool.Get()
	if poolItem2 != ts {
		t.Errorf("Second Get() should return original item, got %v, want %v", poolItem2, ts)
	}
}

func TestPoolPutWithSliceStruct(t *testing.T) {
	tss := &TestSliceStruct{
		Items: []int{1, 2, 3, 4, 5},
		Count: 5,
	}

	pool := New(tss)

	newSlice := &TestSliceStruct{
		Items: []int{10, 20, 30},
		Count: 3,
	}

	pool.Put(newSlice)

	if newSlice.Count != 0 || len(newSlice.Items) != 0 {
		t.Errorf("Put() did not call Reset() properly on slice struct, got: Count=%d, Items=%v",
			newSlice.Count, newSlice.Items)
	}

	if cap(newSlice.Items) != 3 {
		t.Errorf("Reset() should preserve slice capacity, got cap=%d, want 3", cap(newSlice.Items))
	}
}

func TestPoolPutWithArrayStruct(t *testing.T) {
	tas := &TestArrayStruct{
		Values: [3]int{100, 200, 300},
		Tag:    "original",
	}

	pool := New(tas)

	newArray := &TestArrayStruct{
		Values: [3]int{400, 500, 600},
		Tag:    "new",
	}

	pool.Put(newArray)

	if newArray.Values != [3]int{} || newArray.Tag != "" {
		t.Errorf("Put() did not call Reset() properly on array struct, got: Values=%v, Tag=%s",
			newArray.Values, newArray.Tag)
	}
}

func TestPoolWithMultipleCalls(t *testing.T) {
	ts := &TestStruct{
		ID:   1,
		Name: "multi",
		Data: []byte{1},
	}

	pool := New(ts)

	// Put 5 items
	putItems := make([]*TestStruct, 5)
	for i := 0; i < 5; i++ {
		item := &TestStruct{
			ID:   i + 10,
			Name: "temp",
			Data: []byte{byte(i)},
		}
		putItems[i] = item
		pool.Put(item)

		if item.ID != 0 || item.Name != "" || len(item.Data) != 0 {
			t.Errorf("Put() failed to reset item on iteration %d: %+v", i, item)
		}
	}

	// Get should return items in LIFO order (last put first)
	for i := 4; i >= 0; i-- {
		poolItem := pool.Get()
		if poolItem != putItems[i] {
			t.Errorf("Get() at iteration %d returned wrong item, got %v, want %v", 4-i, poolItem, putItems[i])
		}
	}

	// Finally, the original item should still be in the pool
	poolItem := pool.Get()
	if poolItem != ts {
		t.Errorf("Original item should still be in pool, got %v, want %v", poolItem, ts)
	}
}

func TestPoolNilSafety(t *testing.T) {

	ts := &TestStruct{
		ID:   999,
		Name: "nil_test",
		Data: nil,
	}

	pool := New(ts)

	item := pool.Get()
	if item.Data != nil {
		t.Errorf("Expected nil Data field, got %v", item.Data)
	}

	nilItem := &TestStruct{
		ID:   111,
		Name: "nil_item",
		Data: nil,
	}

	pool.Put(nilItem)

	if nilItem.ID != 0 || nilItem.Name != "" || nilItem.Data != nil {
		t.Errorf("Put() should reset even with nil slice, got: %+v", nilItem)
	}
}

func TestPoolWithGeneratedReset(t *testing.T) {
	ms := &testpkg.MyStruct{
		ID:   123,
		Name: "generated",
		Tags: []string{"a", "b", "c"},
		Options: map[string]int{
			"x": 1,
			"y": 2,
		},
		Ptr: new(int),
	}
	*ms.Ptr = 999

	pool := New(ms)

	item := pool.Get()
	if item != ms {
		t.Errorf("Get() returned different MyStruct item")
	}

	// Put the original item back
	pool.Put(item)

	ms2 := &testpkg.MyStruct{
		ID:   456,
		Name: "to_reset",
		Tags: []string{"d", "e"},
		Options: map[string]int{
			"z": 3,
		},
		Ptr: new(int),
	}
	*ms2.Ptr = 777

	pool.Put(ms2)

	if ms2.ID != 0 || ms2.Name != "" || len(ms2.Tags) != 0 || len(ms2.Options) != 0 || *ms2.Ptr != 0 {
		t.Errorf("Put() did not properly reset MyStruct: %+v", ms2)
	}

	// Get should return the most recently put item (ms2) - LIFO
	item2 := pool.Get()
	if item2 != ms2 {
		t.Errorf("Get() should return the put item, got %v, want %v", item2, ms2)
	}

	// Next Get should return the original item (ms)
	item3 := pool.Get()
	if item3 != ms {
		t.Errorf("Second Get() should return original item, got %v, want %v", item3, ms)
	}
}
