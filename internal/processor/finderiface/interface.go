package finderiface

import "github.com/pysf/turbo-journey/internal/emr"

type EmissionProcessor interface {
	ProcessEmission(emr emr.EmissionRecord)
	Result() interface{}
}
