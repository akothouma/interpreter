package main

type MemoryManager struct {
	Variables map[string]int
	Slices    map[string][]int
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		Variables: make(map[string]int),
		Slices:    make(map[string][]int),
	}
}

func (mm *MemoryManager) GetVariable(name string) (int, bool) {
	val, ok := mm.Variables[name]
	return val, ok
}

func (mm *MemoryManager) SetVariable(name string, value int) {
	mm.Variables[name] = value
}

func (mm *MemoryManager) GetSlice(name string) ([]int, bool) {
	slice, ok := mm.Slices[name]
	return slice, ok
}

func (mm *MemoryManager) SetSlice(name string, slice []int) {
	mm.Slices[name] = slice
}