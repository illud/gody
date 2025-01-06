package infraestructure

import (
	"time"

	actionsModel "github.com/gody-server/app/actions/domain/models"
	// uncomment this a change _ for db when you are making database queries
	db "github.com/gody-server/adapters/database"
)

type ActionsDb struct {
	// Add any dependencies or configurations related to the UserRepository here if needed.
}

func NewActionsDb() *ActionsDb {
	// Initialize any dependencies and configurations for the ActionsRepository here if needed.
	return &ActionsDb{}
}

func (a *ActionsDb) CreateActions(actions actionsModel.Actions) error {
	// Insert into database new user
	action := actionsModel.Actions{
		ActionName: actions.ActionName,
		Steps:      actions.Steps,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"), // Manually setting timestamp
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	result := db.Client().Create(&action)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *ActionsDb) GetActions() ([]actionsModel.Actions, error) {
	var actions []actionsModel.Actions
	result := db.Client().Order("id DESC").Find(&actions)
	if result.Error != nil {
		return nil, result.Error
	}
	return actions, nil
}

func (a *ActionsDb) GetOneActions(actionsId int) (actionsModel.Actions, error) {
	var action actionsModel.Actions
	result := db.Client().Where("id = ?", actionsId).First(&action)
	if result.Error != nil {
		return actionsModel.Actions{}, result.Error
	}
	return action, nil
}

func (a ActionsDb) UpdateActions(actionsId int, actions actionsModel.Actions) error {
	action := actionsModel.Actions{
		ActionName: actions.ActionName,
		Steps:      actions.Steps,
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	result := db.Client().Where("id = ?", actionsId).Updates(&action)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a ActionsDb) DeleteActions(actionsId int) error {
	result := db.Client().Where("id = ?", actionsId).Delete(&actionsModel.Actions{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a ActionsDb) Run(actions actionsModel.ActionRun) error {
	// Implement your creation logic here
	return nil
}
