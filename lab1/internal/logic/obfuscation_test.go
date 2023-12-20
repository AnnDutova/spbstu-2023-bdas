package logic_test

import (
	"encoding/xml"
	"testing"

	"github.com/AnnDutova/bdas/lab1/internal/logic"
)

func TestObfuscationAndDeobfuscation(t *testing.T) {
	xmlData := `
		<employees>
			<employee id="111">
				<firstName>Lokesh</firstName>
				<lastName>Gupta</lastName>
				<location>India</location>
			</employee>
			<employee id="222">
				<firstName>Alex</firstName>
				<lastName>Gussin</lastName>
				<location>Russia</location>
			</employee>
		</employees>`

	obfuscatedData := logic.Obfuscation(xmlData)
	deobfuscatedData := logic.Deobfuscation(obfuscatedData)

	var original logic.Employees
	err := xml.Unmarshal([]byte(xmlData), &original)
	if err != nil {
		t.Fatal(err)
	}

	var deobfuscated logic.Employees
	err = xml.Unmarshal([]byte(deobfuscatedData), &deobfuscated)
	if err != nil {
		t.Fatal(err)
	}

	if !employeesEqual(original, deobfuscated) {
		t.Error("Deobfuscated data does not match the original data")
	}
}

func employeesEqual(e1, e2 logic.Employees) bool {
	if len(e1.EmployeeList) != len(e2.EmployeeList) {
		return false
	}

	for i := range e1.EmployeeList {
		if e1.EmployeeList[i].ID != e2.EmployeeList[i].ID ||
			e1.EmployeeList[i].FirstName != e2.EmployeeList[i].FirstName ||
			e1.EmployeeList[i].LastName != e2.EmployeeList[i].LastName ||
			e1.EmployeeList[i].Location != e2.EmployeeList[i].Location {
			return false
		}
	}

	return true
}

func TestShiftRunes(t *testing.T) {
	// Тестовые данные
	input := "abcXYZ123"

	// Тестируем shiftRunes с положительным сдвигом
	shifted := logic.ShiftRunes(input, 3)
	expectedShifted := "def[\\]456"
	if shifted != expectedShifted {
		t.Errorf("Expected: %s, Got: %s", expectedShifted, shifted)
	}

	// Тестируем shiftRunes с отрицательным сдвигом
	unshifted := logic.ShiftRunes(shifted, -3)
	if unshifted != input {
		t.Errorf("Expected: %s, Got: %s", input, unshifted)
	}
}
