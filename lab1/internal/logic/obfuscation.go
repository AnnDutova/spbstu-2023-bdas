package logic

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type Employee struct {
	ID        string `xml:"id,attr"`
	FirstName string `xml:"firstName"`
	LastName  string `xml:"lastName"`
	Location  string `xml:"location"`
}

type Employees struct {
	XMLName      xml.Name   `xml:"employees"`
	EmployeeList []Employee `xml:"employee"`
}

const (
	shift = 4
)

func Obfuscation(content string) string {
	return obfuscate(string(content), shift)
}

func Deobfuscation(content string) string {
	return obfuscate(string(content), -shift)
}

func obfuscate(data string, shift int) string {
	fmt.Println(data)
	var text Employees
	err := xml.Unmarshal([]byte(data), &text)
	if err != nil {
		fmt.Printf("error: %v", err)
		return ""
	}
	fmt.Println(text)

	result := make([]Employee, 0)
	for _, employe := range text.EmployeeList {
		result = append(result, Employee{
			ID:        employe.ID,
			FirstName: ShiftRunes(employe.FirstName, shift),
			LastName:  ShiftRunes(employe.LastName, shift),
			Location:  ShiftRunes(employe.Location, shift),
		})
	}

	newEmployees := Employees{
		EmployeeList: result,
	}

	xmlText, err := xml.MarshalIndent(newEmployees, " ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return ""
	}
	return string(xmlText)
}

func ShiftRunes(input string, shift int) string {
	var buf bytes.Buffer
	for _, w := range input {
		buf.WriteRune(rune(int(w) + shift))
	}
	return buf.String()
}
