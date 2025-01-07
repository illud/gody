package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	ftp "github.com/gody-server/adapters/ftp"
	githubAdatpter "github.com/gody-server/adapters/gogithub"
	actionsModel "github.com/gody-server/app/actions/domain/models"
	actionsInterface "github.com/gody-server/app/actions/domain/repositories"
)

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
	fmt.Println("originalDir github: ", originalDir)
	if err != nil {
		return err
	}

	if action.Github.GithubExecute {
		// Change the working directory to the project path
		absPath, err := filepath.Abs(action.Github.GithubProjectPath)
		if err != nil {
			return err
		}

		err = os.Chdir(absPath)
		if err != nil {
			return err
		}

		err = githubAdatpter.CheckAndUpdateRepository(action.Github.GithubToken, action.Github.RepositoryOwner, action.Github.RepositoryName, action.Github.BranchName, action.Github.GithubProjectPath)
		if err != nil {
			return err
		}

	}

	for _, step := range action.Steps {
		if step.StepType == BAT {
			// Step 1: Create a .bat file
			batFileName := "run.bat"
			batFileContent := `@echo off ` + "\n" + step.Step

			// Write the content to a .bat file
			err = os.WriteFile(action.StepsPath+"/"+batFileName, []byte(batFileContent), 0644)
			if err != nil {
				fmt.Println("Error creating .bat file:", err)
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
				fmt.Println("Error setting up pipes:", err)
				return err
			}

			err = cmd.Start()
			if err != nil {
				fmt.Println("Error starting command:", err)
				return err
			}

			// Read the output from stdout and stderr
			// outBytes, _ := io.ReadAll(stdout)
			errBytes, _ := io.ReadAll(stderr)

			err = cmd.Wait() // Wait for the command to finish
			if err != nil {
				fmt.Println("Command execution failed:", err)
				return errors.New("Command failed: " + string(errBytes))
			}

			// Step 3: Delete the .bat file after successful execution
			err = os.Remove(action.StepsPath + "/" + batFileName)
			if err != nil {
				fmt.Println("Error deleting .bat file:", err)
				return err
			}
		}

		if step.StepType == SH {
			// Step 1: Create a .sh file
			shFileName := "run.sh"
			shFileContent := `#!/bin/bash ` + "\n" + step.Step

			// Write the content to a .sh file
			err = os.WriteFile(action.StepsPath+"/"+shFileName, []byte(shFileContent), 0644)
			if err != nil {
				fmt.Println("Error creating .sh file:", err)
				return err
			}

			// Step 2: Run the .sh file
			cmd := exec.Command("sh", action.StepsPath+"/"+shFileName)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println("Error running .sh file:", err)
				return err
			}

			// Step 3: Delete the .sh file after successful execution
			err = os.Remove(action.StepsPath + "/" + shFileName)
			if err != nil {
				fmt.Println("Error deleting .sh file:", err)
				return err
			}
		}
	}

	if action.Ftp.FtpExecute {
		err = ftp.Ftp(action.Ftp.FtpServer, action.Ftp.Username, action.Ftp.Password, action.Ftp.ProjectPath, action.Ftp.FtpDirectory)
		if err != nil {
			fmt.Println(err)
		}
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

	err = s.actionsRepository.Run(actions)
	if err != nil {
		return err
	}

	err = os.Chdir(originalDir)
	if err != nil {
		return err
	}

	return nil
}
