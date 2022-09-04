package emr

import (
	"strconv"
	"time"
)

type EmissionRecord struct {
	Start    string
	End      string
	CO2      string
	TotalCO2 float64
}

func NewEmissionRecord(d []string) (*EmissionRecord, error) {

	emr := &EmissionRecord{
		Start: d[0],
		End:   d[1],
		CO2:   d[2],
	}
	totalCO2, err := emr.CalculateTotalCO2()
	if err != nil {
		return nil, err
	}

	emr.TotalCO2 = *totalCO2
	return emr, nil

}

func (d *EmissionRecord) StartTime() (*time.Time, error) {
	startDate, err := parseDate(d.Start)
	if err != nil {
		return nil, err
	}
	return startDate, nil
}

func (d *EmissionRecord) EndTime() (*time.Time, error) {
	endDate, err := parseDate(d.End)
	if err != nil {
		return nil, err
	}
	return endDate, nil
}

func (d *EmissionRecord) CO2Emission() (*int64, error) {
	co2, err := strconv.ParseInt(d.CO2, 10, 64)
	if err != nil {
		return nil, err
	}

	return &co2, nil
}

func (d *EmissionRecord) CalculateTotalCO2() (*float64, error) {

	start, err := d.StartTime()
	if err != nil {
		return nil, err
	}

	end, err := d.EndTime()
	if err != nil {
		return nil, err
	}

	co2, err := d.CO2Emission()
	if err != nil {
		return nil, err
	}

	exclusiveEnd := end.Truncate(time.Second * 1)
	diff := exclusiveEnd.Sub(*start)
	d.TotalCO2 = float64(*co2) * diff.Seconds()

	return &d.TotalCO2, nil

}

func parseDate(dateStr string) (*time.Time, error) {
	epoch, err := strconv.ParseInt(dateStr, 10, 64)
	if err != nil {
		return nil, err
	}

	date := time.Unix(epoch, 0)
	return &date, nil
}
