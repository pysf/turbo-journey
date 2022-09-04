package processor

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/pysf/turbo-journey/internal/emr"
)

func TestHighEmissionFinder(t *testing.T) {

	finder := NewHighEmissionProcessor()
	emrs, err := RandomEMRs(10)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range *emrs {
		finder.ProcessEmission(v)
	}

	got := finder.Result()

	want := HighestEmission(emrs)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Finder(): got %v, want: %v", got, want)
	}
}

func RandomEMRs(count int) (*[]emr.EmissionRecord, error) {

	result := make([]emr.EmissionRecord, 0, count)

	start := time.Now()
	for i := 0; i < count; i++ {
		end := start.Add(time.Duration(time.Duration(rand.Intn(72)+1) * time.Hour))
		co2 := rand.Int63n(5) + 1
		emr, err := emr.NewEmissionRecord([]string{
			fmt.Sprintf("%v", start.Unix()),
			fmt.Sprintf("%v", end.Unix()),
			fmt.Sprintf("%v", co2),
		})

		if err != nil {
			return nil, fmt.Errorf("RandomEMR: err=%w", err)
		}
		result = append(result, *emr)

	}

	return &result, nil
}

func HighestEmission(emrs *[]emr.EmissionRecord) *emr.EmissionRecord {

	var maxCO2 emr.EmissionRecord

	for _, v := range *emrs {
		if maxCO2.TotalCO2 < v.TotalCO2 {
			maxCO2 = v
		}
	}

	return &maxCO2
}

func BenchmarkHighEmissionFinder(b *testing.B) {

	for i := 0; i < b.N; i++ {

		finder := NewHighEmissionProcessor()
		emrs, err := RandomEMRs(1000000)
		if err != nil {
			b.Fatal(err)
		}

		for _, v := range *emrs {
			finder.ProcessEmission(v)
		}

		finder.Result()
	}
}
