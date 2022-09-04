package reader

import (
	"strings"
	"testing"

	"github.com/pysf/turbo-journey/internal/emr"
)

func TestCSVProcessor(t *testing.T) {

	var data strings.Builder

	data.WriteString("1642700100;1642700200;4")
	data.WriteString("\n")
	data.WriteString("1642700000;1642700100;3")
	data.WriteString("\n")
	data.WriteString("1642700050;1642700150;5")
	r := strings.NewReader(data.String())

	mockEMR := &MockEMRProcessor{}
	csvPrecossor := NewCSVReader(r, WithProcessor(mockEMR), WithComma(';'))

	if err := csvPrecossor.Start(); err != nil {
		t.Fatalf("CSVProcessor() err= %v", err)
	}

	wantedFnCalls := 3
	if mockEMR.Calls != wantedFnCalls {
		t.Fatalf("CSVProcessor() got = %d function calls, want %d function calls", mockEMR.Calls, wantedFnCalls)
	}

}

type MockEMRProcessor struct {
	Calls int
}

func (m *MockEMRProcessor) ProcessEmission(emr emr.EmissionRecord) {
	m.Calls = m.Calls + 1
}

func (m *MockEMRProcessor) Result() interface{} {
	return nil
}
