package main

import (
	"fmt"
	"sync"
)

type TempInterface interface {
	Save(int)
	Print()
}

type temp struct {
	mTemp map[int]struct{}
	aTemp []int
	mu    sync.Mutex
}

// func NewTemp(m map[int]struct{}) *temp

func NewTemp(m map[int]struct{}) TempInterface {
	return &temp{
		mTemp: m,
	}
}

func (t *temp) Save(n int) {
	if _, ok := t.mTemp[n]; !ok {
		t.mu.Lock()
		t.mTemp[n] = struct{}{}
		t.aTemp = append(t.aTemp, n)
		t.mu.Unlock()
	}
}

func (t *temp) Print() {
	fmt.Println(t.aTemp)
}

func main() {

	// estrutura vazia, que ocupa 0 bytes de memória
	m := map[int]struct{}{}

	// t := NewTemp(m)
	var t TempInterface = NewTemp(m)

	arrayEx := []int{4, 8, 9, 4, 7, 8}

	for _, v := range arrayEx {
		t.Save(v)
	}

	t.Print()

}
