package main

import (
	"fmt"
	"sync"
)

// O código Go fornecido implementa uma interface TempInterface e uma struct temp para armazenar e imprimir números inteiros únicos.

type TempInterface interface {
	Save(int)
	Print()
}

type temp struct {
	mapTemp map[int]struct{}
	arrTemp []int
	mu      sync.Mutex
}

// func NewTemp(m map[int]struct{}) *temp

func NewTemp(m map[int]struct{}) TempInterface {
	return &temp{
		mapTemp: m,
	}
}

func (t *temp) Save(n int) {
	if _, ok := t.mapTemp[n]; !ok {
		t.mu.Lock()
		t.mapTemp[n] = struct{}{}
		t.arrTemp = append(t.arrTemp, n)
		t.mu.Unlock()
	}
}

func (t *temp) Print() {
	t.mu.Lock()
	fmt.Println(t.arrTemp)
	t.mu.Unlock()
}

func main() {

	// estrutura vazia, que ocupa 0 bytes de memória
	m := map[int]struct{}{}

	t := NewTemp(m)
	// var t TempInterface = NewTemp(m)

	arrayEx := []int{4, 8, 9, 4, 7, 8}

	for _, v := range arrayEx {
		t.Save(v)
	}

	t.Print()

}
