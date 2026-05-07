package testpkg

// generate:reset
type MyStruct struct {
	ID      int
	Name    string
	Tags    []string
	Options map[string]int
	Ptr     *int
	Embed   Embedded
}

type Embedded struct {
	Value float64
}
