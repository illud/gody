package services

import (
	executionhistoryModel "github.com/gody-server/app/executionhistory/domain/models"
	executionhistoryInterface "github.com/gody-server/app/executionhistory/domain/repositories"
)

type Service struct {
	executionhistoryRepository executionhistoryInterface.IExecutionhistory
}

func NewService(executionhistoryRepository executionhistoryInterface.IExecutionhistory) *Service {
	return &Service{
		executionhistoryRepository: executionhistoryRepository,
	}
}

func (s *Service) CreateExecutionhistory(executionhistory executionhistoryModel.Executionhistory) error {
	err := s.executionhistoryRepository.CreateExecutionhistory(executionhistory)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetExecutionhistory(actionId int) ([]executionhistoryModel.Executionhistory, error) {
	result, err := s.executionhistoryRepository.GetExecutionhistory(actionId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetOneExecutionhistory(executionHistoryId int) (executionhistoryModel.Executionhistory, error) {
	result, err := s.executionhistoryRepository.GetOneExecutionhistory(executionHistoryId)
	if err != nil {
		return executionhistoryModel.Executionhistory{}, err
	}
	return result, nil
}

func (s *Service) UpdateExecutionhistory(executionHistoryId int) error {
	err := s.executionhistoryRepository.UpdateExecutionhistory(executionHistoryId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteExecutionhistory(executionHistoryId int) error {
	err := s.executionhistoryRepository.DeleteExecutionhistory(executionHistoryId)
	if err != nil {
		return err
	}
	return nil
}
