package reader

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/pysf/turbo-journey/internal/emr"
	"github.com/pysf/turbo-journey/internal/processor/processoriface"
)

type CSVReadrOption func(*CSVReader)

type CSVReader struct {
	Comma              rune
	Source             io.Reader
	EmissionProcessors []processoriface.EmissionProcessor
}

func NewCSVReader(r io.Reader, op ...CSVReadrOption) *CSVReader {
	csvReader := &CSVReader{
		Comma:  ';',
		Source: r,
	}

	for _, o := range op {
		o(csvReader)
	}
	return csvReader
}

func WithProcessor(emp processoriface.EmissionProcessor) CSVReadrOption {
	return func(c *CSVReader) {
		c.EmissionProcessors = append(c.EmissionProcessors, emp)
	}
}

func WithComma(comma rune) CSVReadrOption {
	return func(c *CSVReader) {
		c.Comma = comma
	}
}

func (c *CSVReader) Start() error {
	r := csv.NewReader(c.Source)
	r.Comma = c.Comma

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("Start: err= %w", err)
		}

		for _, emp := range c.EmissionProcessors {
			emissionRecord, err := emr.NewEmissionRecord(record)

			if err != nil {
				return fmt.Errorf("Start: create emr err=%w", err)
			}

			emp.ProcessEmission(*emissionRecord)
		}
	}

	return nil
}
