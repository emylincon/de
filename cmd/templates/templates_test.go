package templates

import (
	"fmt"
	"testing"
)

func TestPythonTemplate(t *testing.T) {
	if fmt.Sprintf("%T", PythonTemplate()) != fmt.Sprintf("%T", "type") {
		t.Errorf("TestPythonTemplate : types not equal")
	}
}

func TestGolangTemplate(t *testing.T) {
	if fmt.Sprintf("%T", GolangTemplate()) != fmt.Sprintf("%T", "type") {
		t.Errorf("TestPythonTemplate : types not equal")
	}
}
