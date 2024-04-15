package usecases

import (
	"github.com/mezink-test/src/services"
)

// get list data records by params
func GetRecordsUsecase(startDate *string, endDate *string, minCount *int, maxCount *int) ([]services.Record, error) {
	// call record find service
	users, err := services.Find(services.FilterFindRecord{
		StartDate: startDate,
		EndDate:   endDate,
		MinCount:  minCount,
		MaxCount:  maxCount,
	})
	if err != nil {
		return nil, err
	}

	return users, err
}

// create data record
func CreateRecordUsecase(name string, marks []int) error {
	// call record create service
	err := services.Create(services.DataCreateRecord{
		Name:  name,
		Marks: marks,
	})
	if err != nil {
		return err
	}

	return err
}
