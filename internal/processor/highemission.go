package processor

import (
	"github.com/pysf/turbo-journey/internal/emr"
	"github.com/pysf/turbo-journey/internal/processor/processoriface"
)

type HighEmissionProcessor struct {
	highestEMR emr.EmissionRecord
}

func NewHighEmissionProcessor() processoriface.EmissionProcessor {
	return &HighEmissionProcessor{}
}

func (f *HighEmissionProcessor) ProcessEmission(emr emr.EmissionRecord) {
	if emr.TotalCO2 > f.highestEMR.TotalCO2 {
		f.highestEMR = emr
	}
}

func (f *HighEmissionProcessor) Result() interface{} {
	return &f.highestEMR
}
