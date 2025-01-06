package repositories

import (
	actionsModel "github.com/gody-server/app/actions/domain/models"
)

type IActions interface {
	CreateActions(actions actionsModel.Actions) error
	GetActions() ([]actionsModel.Actions, error)
	GetOneActions(actionsId int) (actionsModel.Actions, error)
	UpdateActions(actionsId int, actions actionsModel.Actions) error
	DeleteActions(actionsId int) error
	Run(actions actionsModel.ActionRun) error
}
