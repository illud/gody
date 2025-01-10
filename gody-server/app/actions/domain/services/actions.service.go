package services

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	ftp "github.com/gody-server/adapters/ftp"
	githubAdatpter "github.com/gody-server/adapters/gogithub"
	actionsModel "github.com/gody-server/app/actions/domain/models"
	actionsInterface "github.com/gody-server/app/actions/domain/repositories"
	executionhistoryModel "github.com/gody-server/app/executionhistory/domain/models"
	executionhistoryServices "github.com/gody-server/app/executionhistory/domain/services"
	executionhistoryDatabase "github.com/gody-server/app/executionhistory/infraestructure"
)

// Create a executionhistoryDb instance
var executionhistoryDb = executionhistoryDatabase.NewExecutionhistoryDb()

// Create a Service instance using the executionhistoryDb
var executionService = executionhistoryServices.NewService(executionhistoryDb)

type Service struct {
	actionsRepository actionsInterface.IActions
}

func NewService(actionsRepository actionsInterface.IActions) *Service {
	return &Service{
		actionsRepository: actionsRepository,
	}
}

func (s *Service) CreateActions(actions actionsModel.Actions) error {
	err := s.actionsRepository.CreateActions(actions)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetActions() ([]actionsModel.Actions, error) {
	result, err := s.actionsRepository.GetActions()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetOneActions(actionsId int) (actionsModel.Actions, error) {
	result, err := s.actionsRepository.GetOneActions(actionsId)
	if err != nil {
		return actionsModel.Actions{}, err
	}
	return result, nil
}

func (s *Service) UpdateActions(actionsId int, actions actionsModel.Actions) error {
	err := s.actionsRepository.UpdateActions(actionsId, actions)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteActions(actionsId int) error {
	err := s.actionsRepository.DeleteActions(actionsId)
	if err != nil {
		return err
	}
	return nil
}

const (
	GITHUB = iota + 1
)

const (
	BAT = iota + 1
	SH
)

func (s *Service) Run(actions actionsModel.ActionRun) error {
	var history []executionhistoryModel.Step

	actionResult, err := s.GetOneActions(actions.ActionId)
	if err != nil {
		return err
	}

	var action actionsModel.RunAction
	err = json.Unmarshal([]byte(actionResult.Steps), &action)
	if err != nil {
		return err
	}

	originalDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if action.Github.GithubExecute {
		start := time.Now()

		// Change the working directory to the project path
		absPath, err := filepath.Abs(action.Github.GithubProjectPath)
		if err != nil {
			history = append(history, executionhistoryModel.Step{
				ExecutionName:   "Github",
				ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
				ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
				ExecutionStatus: "Failed",
				ExecutionError:  err.Error(),
			})
			return err
		}

		err = os.Chdir(absPath)
		if err != nil {
			history = append(history, executionhistoryModel.Step{
				ExecutionName:   "Github",
				ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
				ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
				ExecutionStatus: "Failed",
				ExecutionError:  err.Error(),
			})
			return err
		}

		err = githubAdatpter.CheckAndUpdateRepository(action.Github.GithubToken, action.Github.RepositoryOwner, action.Github.RepositoryName, action.Github.BranchName, action.Github.GithubProjectPath)
		if err != nil {
			history = append(history, executionhistoryModel.Step{
				ExecutionName:   "Github",
				ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
				ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
				ExecutionStatus: "Failed",
				ExecutionError:  err.Error(),
			})
			return err
		}

		history = append(history, executionhistoryModel.Step{
			ExecutionName:   "Github",
			ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
			ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
			ExecutionStatus: "Success",
			ExecutionError:  "",
		})
	}

	for _, step := range action.Steps {
		if step.StepType == BAT {
			start := time.Now()

			// Step 1: Create a .bat file
			batFileName := "run.bat"
			batFileContent := `@echo off ` + "\n" + step.Step

			// Write the content to a .bat file
			err = os.WriteFile(action.StepsPath+"/"+batFileName, []byte(batFileContent), 0644)
			if err != nil {
				history = append(history, executionhistoryModel.Step{
					ExecutionName:   step.Step,
					ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
					ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
					ExecutionStatus: "Failed",
					ExecutionError:  err.Error(),
				})
				return err
			}

			// Step 2: Run the .bat file
			cmd := exec.Command(action.StepsPath + "/" + batFileName)
			// stdout, err := cmd.StdoutPipe() // Capture stdout separately
			// if err != nil {
			// 	fmt.Println("Error setting up pipes:", err)
			// 	return err
			// }
			stderr, err := cmd.StderrPipe() // Capture stderr separately
			if err != nil {
				history = append(history, executionhistoryModel.Step{
					ExecutionName:   step.Step,
					ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
					ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
					ExecutionStatus: "Failed",
					ExecutionError:  err.Error(),
				})
				return err
			}

			err = cmd.Start()
			if err != nil {
				history = append(history, executionhistoryModel.Step{
					ExecutionName:   step.Step,
					ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
					ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
					ExecutionStatus: "Failed",
					ExecutionError:  err.Error(),
				})
				return err
			}

			// Read the output from stdout and stderr
			// outBytes, _ := io.ReadAll(stdout)
			errBytes, _ := io.ReadAll(stderr)

			err = cmd.Wait() // Wait for the command to finish
			if err != nil {
				history = append(history, executionhistoryModel.Step{
					ExecutionName:   step.Step,
					ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
					ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
					ExecutionStatus: "Failed",
					ExecutionError:  string(errBytes),
				})
				return errors.New("Command failed: " + string(errBytes))
			}

			// Step 3: Delete the .bat file after successful execution
			err = os.Remove(action.StepsPath + "/" + batFileName)
			if err != nil {
				history = append(history, executionhistoryModel.Step{
					ExecutionName:   step.Step,
					ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
					ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
					ExecutionStatus: "Failed",
					ExecutionError:  err.Error(),
				})
				return err
			}

			history = append(history, executionhistoryModel.Step{
				ExecutionName:   step.Step,
				ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
				ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
				ExecutionStatus: "Success",
				ExecutionError:  "",
			})
		}

		if step.StepType == SH {
			start := time.Now()
			// Step 1: Create a .sh file
			shFileName := "run.sh"
			shFileContent := `#!/bin/bash ` + "\n" + step.Step

			// Write the content to a .sh file
			err = os.WriteFile(action.StepsPath+"/"+shFileName, []byte(shFileContent), 0644)
			if err != nil {
				history = append(history, executionhistoryModel.Step{
					ExecutionName:   step.Step,
					ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
					ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
					ExecutionStatus: "Failed",
					ExecutionError:  err.Error(),
				})
				return err
			}

			// Step 2: Run the .sh file
			cmd := exec.Command("sh", action.StepsPath+"/"+shFileName)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				history = append(history, executionhistoryModel.Step{
					ExecutionName:   step.Step,
					ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
					ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
					ExecutionStatus: "Failed",
					ExecutionError:  err.Error(),
				})
				return err
			}

			// Step 3: Delete the .sh file after successful execution
			err = os.Remove(action.StepsPath + "/" + shFileName)
			if err != nil {
				history = append(history, executionhistoryModel.Step{
					ExecutionName:   step.Step,
					ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
					ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
					ExecutionStatus: "Failed",
					ExecutionError:  err.Error(),
				})
				return err
			}

			history = append(history, executionhistoryModel.Step{
				ExecutionName:   step.Step,
				ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
				ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
				ExecutionStatus: "Success",
				ExecutionError:  "",
			})
		}
	}

	if action.Ftp.FtpExecute {
		start := time.Now()
		err = ftp.Ftp(action.Ftp.FtpServer, action.Ftp.Username, action.Ftp.Password, action.Ftp.ProjectPath, action.Ftp.FtpDirectory)
		if err != nil {
			history = append(history, executionhistoryModel.Step{
				ExecutionName:   "Ftp",
				ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
				ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
				ExecutionStatus: "Failed",
				ExecutionError:  err.Error(),
			})
			return err
		}

		history = append(history, executionhistoryModel.Step{
			ExecutionName:   "Ftp",
			ExecutionTime:   strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64),
			ExecutionDate:   time.Now().Format("2006-01-02 15:04:05"),
			ExecutionStatus: "Success",
			ExecutionError:  "",
		})
	}

	// Run `git pull origin main` to pull the latest changes
	// cmd := exec.Command("go", "build", "main.go")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err = cmd.Run()
	// if err != nil {
	// 	return err
	// }

	// cmd2 := exec.Command("./main")
	// cmd2.Stdout = os.Stdout
	// cmd2.Stderr = os.Stderr
	// err = cmd2.Start()
	// if err != nil {
	// 	return err
	// }

	err = os.Chdir(originalDir)
	if err != nil {
		return err
	}

	// convert history to string including json format
	historyString, err := json.Marshal(history)
	if err != nil {
		return err
	}

	err = executionService.CreateExecutionhistory(executionhistoryModel.Executionhistory{
		ActionID:  actions.ActionId,
		Step:      string(historyString),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return err
	}
	// err = s.actionsRepository.Run(actions)
	// if err != nil {
	// 	return err
	// }

	return nil
}
