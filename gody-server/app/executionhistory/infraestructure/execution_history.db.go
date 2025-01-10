package infraestructure

import (
	executionhistoryModel "github.com/gody-server/app/executionhistory/domain/models"
	// uncomment this a change _ for db when you are making database queries
	db "github.com/gody-server/adapters/database"
)

type ExecutionhistoryDb struct {
	// Add any dependencies or configurations related to the UserRepository here if needed.
}

func NewExecutionhistoryDb() *ExecutionhistoryDb {
	// Initialize any dependencies and configurations for the ExecutionhistoryRepository here if needed.
	return &ExecutionhistoryDb{}
}

func (e *ExecutionhistoryDb) CreateExecutionhistory(executionhistory executionhistoryModel.Executionhistory) error {
	history := executionhistoryModel.Executionhistory{
		ActionID:  executionhistory.ActionID,
		Step:      executionhistory.Step,
		CreatedAt: executionhistory.CreatedAt,
		UpdatedAt: executionhistory.UpdatedAt,
	}

	result := db.Client().Create(&history)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (e *ExecutionhistoryDb) GetExecutionhistory(actionId int) ([]executionhistoryModel.Executionhistory, error) {
	var executionHistory []executionhistoryModel.Executionhistory
	result := db.Client().Order("id DESC").Where("action_id = ?", actionId).Find(&executionHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return executionHistory, nil
}

func (e *ExecutionhistoryDb) GetOneExecutionhistory(executionHistoryId int) (executionhistoryModel.Executionhistory, error) {
	// Implement your single retrieval logic here
	return executionhistoryModel.Executionhistory{}, nil
}

func (e ExecutionhistoryDb) UpdateExecutionhistory(executionHistoryId int) error {
	// Implement your update logic here
	return nil
}

func (e ExecutionhistoryDb) DeleteExecutionhistory(executionHistoryId int) error {
	// Implement your deletion logic here
	return nil
}
