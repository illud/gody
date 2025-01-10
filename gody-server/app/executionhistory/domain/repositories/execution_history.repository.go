package repositories

import (
	executionhistoryModel "github.com/gody-server/app/executionhistory/domain/models"
)

type IExecutionhistory interface {
	CreateExecutionhistory(executionhistory executionhistoryModel.Executionhistory) error
	GetExecutionhistory(actionId int) ([]executionhistoryModel.Executionhistory, error)
	GetOneExecutionhistory(executionHistoryId int) (executionhistoryModel.Executionhistory, error)
	UpdateExecutionhistory(executionHistoryId int) error
	DeleteExecutionhistory(executionHistoryId int) error
}
