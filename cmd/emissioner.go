package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/pysf/turbo-journey/internal/emr"
	"github.com/pysf/turbo-journey/internal/processor"
	"github.com/pysf/turbo-journey/internal/reader"
)

func Execute() {

	if len(os.Args) < 2 {
		log.Println("Err: file path is required e.g: emissioner test.csv")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Err: %v \n", err)
		return
	}
	defer file.Close()

	highEmissionProcessor := processor.NewHighEmissionProcessor()
	csvReader := reader.NewCSVReader(file, reader.WithComma(';'), reader.WithProcessor(highEmissionProcessor))
	if err := csvReader.Start(); err != nil {
		log.Printf("Emissioner() err=%v", err)
		panic(err)
	}

	result := highEmissionProcessor.Result()
	emission := result.(*emr.EmissionRecord)
	fmt.Printf("Start: %v End: %v Emission: %vkg/s \n", emission.Start, emission.End, emission.CO2)

}
