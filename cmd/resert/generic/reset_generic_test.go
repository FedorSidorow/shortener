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
	t.Data = t.Data[:0]
}

type TestSliceStruct struct {
	Items []int
	Count int
}

func (t *TestSliceStruct) Reset() {
	t.Items = t.Items[:0]
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
	item2 := pool.Get()

	if item1 != item2 {
		t.Errorf("Multiple Get() calls returned different items: %v vs %v", item1, item2)
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

	poolItem := pool.Get()
	if poolItem.ID != 200 || poolItem.Name != "put_test" || len(poolItem.Data) != 3 {
		t.Errorf("Pool's internal item was incorrectly modified: %+v", poolItem)
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

	for i := 0; i < 5; i++ {
		item := &TestStruct{
			ID:   i + 10,
			Name: "temp",
			Data: []byte{byte(i)},
		}
		pool.Put(item)

		if item.ID != 0 || item.Name != "" || len(item.Data) != 0 {
			t.Errorf("Put() failed to reset item on iteration %d: %+v", i, item)
		}
	}

	poolItem := pool.Get()
	if poolItem.ID != 1 || poolItem.Name != "multi" || len(poolItem.Data) != 1 {
		t.Errorf("Pool's internal item was modified after multiple Put calls: %+v", poolItem)
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

	// Original pool item should be unchanged
	original := pool.Get()
	if original.ID != 123 || original.Name != "generated" || len(original.Tags) != 3 {
		t.Errorf("Pool's internal MyStruct was modified: %+v", original)
	}
}
