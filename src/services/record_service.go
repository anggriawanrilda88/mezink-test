package services

import (
	"errors"
	"log"
	"time"

	"github.com/lib/pq"
	errMsg "github.com/mezink-test/src/errors"
	"github.com/mezink-test/src/services/db"
)

type FilterFindRecord struct {
	StartDate *string
	EndDate   *string
	MinCount  *int
	MaxCount  *int
}

type DataCreateRecord struct {
	Name  string
	Marks []int
}

type Record struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Marks     pq.Int64Array `gorm:"type:integer[]"`
}

// Service to get list record data
func Find(filter FilterFindRecord) ([]Record, error) {
	var models []Record
	q := db.DB.Order("id desc")

	// filter min max count marks
	if filter.MinCount != nil && filter.MaxCount == nil {
		q = q.Where("(SELECT SUM(s) FROM UNNEST(marks) s) >= ?", filter.MinCount)
	} else if filter.MinCount == nil && filter.MaxCount != nil {
		q = q.Where("(SELECT SUM(s) FROM UNNEST(marks) s) <= ?", filter.MaxCount)
	} else if filter.MinCount != nil && filter.MaxCount != nil {
		q = q.Where("(SELECT SUM(s) FROM UNNEST(marks) s) BETWEEN ? AND ?", filter.MinCount, filter.MaxCount)
	}

	// filter date created_at
	if filter.StartDate != nil && filter.EndDate == nil {
		date, _ := time.Parse("2006-01-02", *filter.StartDate)
		q = q.Where("created_at >= ?", date)
	} else if filter.StartDate == nil && filter.EndDate != nil {
		date, _ := time.Parse("2006-01-02", *filter.EndDate)
		oneDayAfter := date.AddDate(0, 0, 1).Add(-time.Second)
		q = q.Where("created_at <= ?", oneDayAfter)
	} else if filter.StartDate != nil && filter.EndDate != nil {
		startDate, _ := time.Parse("2006-01-02", *filter.StartDate)
		endDateParse, _ := time.Parse("2006-01-02", *filter.EndDate)
		endDate := endDateParse.AddDate(0, 0, 1).Add(-time.Second)
		q = q.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	q = q.Find(&models)

	if q.Error != nil {
		log.Print(q.Error)
		err := errors.New(errMsg.FIND_SERVICE_RECORD_ERROR)
		return nil, err
	}

	return models, nil
}

// Service to create record data
func Create(data DataCreateRecord) error {
	var model Record

	var marks []int64
	for _, val := range data.Marks {
		marks = append(marks, int64(val))
	}
	model = Record{
		Name:  data.Name,
		Marks: pq.Int64Array(marks),
	}

	q := db.DB.Save(&model)

	if q.Error != nil {
		log.Print(q.Error)
		err := errors.New(errMsg.CREATE_SERVICE_RECORD_ERROR)
		return err
	}

	return nil
}
