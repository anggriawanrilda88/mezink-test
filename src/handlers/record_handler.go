package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	errMsg "github.com/mezink-test/src/errors"
	"github.com/mezink-test/src/usecases"
	"github.com/mezink-test/src/utils"
)

type GetRecordsQuery struct {
	StartDate *string `form:"startDate"`
	EndDate   *string `form:"endDate"`
	MinCount  *int    `form:"minCount"`
	MaxCount  *int    `form:"maxCount"`
}

type BodyRecord struct {
	Name  string `form:"name"`
	Marks []int  `form:"marks"`
}

type GetRecordsTransform struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalMarks int64     `json:"totalMarks"`
}

// Get request, handler query params, and show response
func GetRecordsHandler(c *gin.Context) {
	var req GetRecordsQuery

	// get query api and add to struct
	if err := c.ShouldBindQuery(&req); err != nil {
		errData := utils.ParseErrorMessage(errors.New(errMsg.GET_LIST_RECORD_REQ_ERROR))
		c.JSON(errData.ErrStatus, errData.ErrData)
		return
	}

	// validate query api
	err := validation.ValidateStruct(&req,
		// Validate StartDate and EndDate format
		validation.Field(&req.StartDate, validation.Date("2006-01-02")),
		validation.Field(&req.EndDate, validation.Date("2006-01-02")),
		// Validate MinCount and MaxCount values
		validation.Field(&req.MinCount, validation.Min(0)),
		validation.Field(&req.MaxCount, validation.Min(0)),
	)
	if err != nil {
		errData := utils.ParseErrorMessage(errors.New(errMsg.GET_LIST_RECORD_VALIDATION_ERROR))
		errData.ErrData["msg"] = err.Error()
		c.JSON(errData.ErrStatus, errData.ErrData)
		return
	}

	// get usecase list records
	records, err := usecases.GetRecordsUsecase(req.StartDate, req.EndDate, req.MinCount, req.MaxCount)
	if err != nil {
		errData := utils.ParseErrorMessage(err)
		c.JSON(errData.ErrStatus, errData.ErrData)
		return
	}

	// transform data to response needed
	var recordTransform []GetRecordsTransform
	for _, val := range records {
		total := int64(0)
		for _, num := range val.Marks {
			total += num
		}

		recordTransform = append(recordTransform, GetRecordsTransform{
			ID:         val.ID,
			CreatedAt:  val.CreatedAt,
			TotalMarks: total,
		})
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Success get list records", "records": recordTransform})
}

// Get request, handler query params, and show response
func CreateRecordHandler(c *gin.Context) {
	var req BodyRecord

	// get query api and add to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		errData := utils.ParseErrorMessage(errors.New(errMsg.CREATE_RECORD_BODY_ERROR))
		c.JSON(errData.ErrStatus, errData.ErrData)
		return
	}

	// validate query api
	err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Marks, validation.Required),
	)
	if err != nil {
		errData := utils.ParseErrorMessage(errors.New(errMsg.CREATE_RECORD_VALIDATION_ERROR))
		errData.ErrData["msg"] = err.Error()
		c.JSON(errData.ErrStatus, errData.ErrData)
		return
	}

	// get usecase list records
	err = usecases.CreateRecordUsecase(req.Name, req.Marks)
	if err != nil {
		errData := utils.ParseErrorMessage(err)
		c.JSON(errData.ErrStatus, errData.ErrData)
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Success create record"})
}
