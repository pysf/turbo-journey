package processoriface

import "github.com/pysf/turbo-journey/internal/emr"

type EmissionProcessor interface {
	ProcessEmission(emr.EmissionRecord)
	Result() interface{}
}
