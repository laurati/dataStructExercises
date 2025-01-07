package main

import (
	"encoding/json"
	"fmt"
)

type ModelCar struct {
	Brand   string
	Name    string
	Version int
}

type Car struct {
	Model             ModelCar
	ManufacturingYear string
	EngineNumber      string
}

type Defect struct {
	Model      ModelCar
	YearsIssue []string
	Code       string
}

type DefectProvider interface {
	GetDefects([]ModelCar) []Defect
}

func NewModelCar(brand, name string) *ModelCar {
	return &ModelCar{
		Brand:   brand,
		Name:    name,
		Version: 1,
	}
}
func (m *ModelCar) incrementVersion() {
	m.Version++
}

func (m *ModelCar) showModelCar() string {
	return fmt.Sprintf("%s %s v%d", m.Brand, m.Name, m.Version)
}

func (m *ModelCar) showTextRepresentation() (string, error) {
	mJson, err := json.Marshal(&m)
	if err != nil {
		return "", err
	}
	return string(mJson), nil
}

// func getDefects(cars []Car, defect DefectProvider) []Defect{

// 	// transform the cars to the car model
// 	// ask the remote service using the car models
// 	// find the match between cars and provided car defect

// }

func main() {

	m1 := NewModelCar("BMW", "X1")
	fmt.Println(m1.showModelCar())

	m1.incrementVersion()
	fmt.Println(m1.showModelCar())

	fmt.Println(m1.showTextRepresentation())

}
