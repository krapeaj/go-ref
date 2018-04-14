package rest

import (
	"fmt"
	"errors"

	"github.com/jinzhu/gorm"

	dto "github.com/1ambda/go-ref/service-gateway/pkg/generated/swagger/rest_model"
	"github.com/1ambda/go-ref/service-gateway/pkg/generated/swagger/rest_server/rest_api/browser_history"
	"github.com/1ambda/go-ref/service-gateway/internal/distributed"
	"github.com/1ambda/go-ref/service-gateway/internal/config"
	"github.com/1ambda/go-ref/service-gateway/internal/model"
)

func addOneBrowserHistory(params browser_history.AddOneParams, db *gorm.DB, dClient distributed.DistributedClient) (*dto.BrowserHistory, *dto.Error) {
	logger := config.GetLogger()
	logger.Infow("Creating BrowserHistory record", "request", params.Body)

	record := model.ConvertFromBrowserHistoryDTO(params.Body)

	if err := db.Create(record).Error; err != nil {
		logger.Errorw("Failed to create new BrowserHistory record: %v", "error", err)
		restError := buildRestError(err, 500)
		return nil, restError
	}

	var count int64 = 0
	err := db.Table(model.BrowserHistoryTable).Count(&count).Error
	if err != nil {
		logger.Errorw("Failed to create new BrowserHistory record: %v", "error", err)
		restError := buildRestError(err, 500)
		return nil, restError
	}

	stringified := fmt.Sprintf("%d", count)
	dClient.Publish(distributed.NewBrowserHistoryCountMessage(stringified))

	restResp := model.ConvertToBrowserHistoryDTO(record)

	return restResp, nil
}

func findOneBrowserHistory(params browser_history.FindOneParams, db *gorm.DB) (*dto.BrowserHistory, *dto.Error) {
	logger := config.GetLogger()
	logger.Infow("Finding BrowserHistory record", "id", params.ID)

	var record model.BrowserHistory

	if err := db.Where("id = ?", params.ID).First(&record).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			logger.Infow("Failed to find BrowserHistory record", "id", params.ID)
			return nil, buildRestError(err, 400)
		}

		logger.Errorw("Failed to find BrowserHistory record due to unknown error",
			"id", params.ID, "error", err)
		return nil, buildRestError(err, 500)
	}

	response := model.ConvertToBrowserHistoryDTO(&record)
	return response, nil
}

func findAllBrowserHistory(params browser_history.FindAllParams, db *gorm.DB) (*dto.Pagination, []*dto.BrowserHistory, *dto.Error) {
	logger := config.GetLogger()
	logger.Info("Finding All BrowserHistory records")

	var records []model.BrowserHistory
	var count int64 = 0
	currentPageOffset := params.CurrentPageOffset
	itemCountPerPage := params.ItemCountPerPage

	dbOffset := int64(*currentPageOffset) * (*itemCountPerPage)

	err := db.
		Table(model.BrowserHistoryTable).
		Count(&count).
		Offset(int(dbOffset)).
		Limit(int(*itemCountPerPage)).
		Find(&records).
		Error

	if err != nil {
		logger.Errorw("Failed to find all BrowserHistory records", "error", err)
		restError := buildRestError(err, 500)
		return nil, nil, restError
	}

	rows := make([]*dto.BrowserHistory, 0)
	for i := range records {
		record := records[i]
		restmodel := model.ConvertToBrowserHistoryDTO(&record)
		rows = append(rows, restmodel)
	}

	pagination := dto.Pagination{
		ItemCountPerPage:  itemCountPerPage,
		CurrentPageOffset: currentPageOffset,
		TotalItemCount:    &count,
	}

	return &pagination, rows, nil
}

func removeOneBrowserHistory(params browser_history.RemoveOneParams, db *gorm.DB) *dto.Error {
	logger := config.GetLogger()
	logger.Infow("Deleting BrowserHistory record", "id", params.ID)

	// https://github.com/jinzhu/gorm/issues/1380
	// https://github.com/jinzhu/gorm/issues/371
	result := db.Where("id = ?", params.ID).Delete(&model.BrowserHistory{})

	if result.Error != nil {
		logger.Errorw("Failed to delete BrowserHistory record due to unknown error",
			"id", params.ID, "error", result.Error)
		restError := buildRestError(result.Error, 500)
		return restError
	}

	if result.RowsAffected < 1 {
		logger.Infow("Failed to find BrowserHistory record before removing", "id", params.ID)
		err := errors.New(gorm.ErrRecordNotFound.Error())
		restError := buildRestError(err, 400)
		return restError
	}

	return nil
}

// TODO(1ambda): PUT doesn't work
//func updateOneAccess(params browser_history.UpdateOneParams, db *gorm.DB) (*rest_model.Access, *rest_model.Error) {
//	logger := config.GetLogger()
//	logger.Infow("Updating BrowserHistory record", "id", params.ID)
//
//	record := model.ConvertFromBrowserHistoryDTO(params.Body)
//	var updated model.BrowserHistory
//
//	// https://github.com/jinzhu/gorm/issues/891
//	result := db.Model(&updated).Where("id = ?", params.ID).Update(record)
//
//	if result.Error != nil {
//		logger.Errorw("Failed to update BrowserHistory record due to unknown error",
//			"id", params.ID, "error", result.Error)
//		restError := buildRestError(result.Error, 500)
//		return nil, restError
//	}
//
//	if result.RowsAffected < 1 {
//		logger.Infow("Failed to find BrowserHistory record before updating", "id", params.ID)
//		err := errors.New(gorm.ErrRecordNotFound.Error())
//		restError := buildRestError(err, 400)
//		return nil, restError
//	}
//
//	response := model.ConvertToBrowserHistoryDTO(&updated)
//	return response, nil
//}