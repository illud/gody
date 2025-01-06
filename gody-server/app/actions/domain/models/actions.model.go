package models

type Actions struct {
	ID         int    `json:"id" db:"id" gorm:"primaryKey"`
	ActionName string `json:"action_name" db:"action_name"`
	Steps      string `json:"steps" db:"steps"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}

type ActionRun struct {
	ActionId int `json:"action_id" db:"action_id"`
}

type Steps struct {
	StepType int    `json:"step_type" db:"step_type"` // 1: windows .bat, 2: linux .sh
	Step     string `json:"step" db:"step"`
}

type RunAction struct {
	Github struct {
		GithubExecute     bool   `json:"github_execute" db:"github_execute"`
		GithubToken       string `json:"github_token" db:"github_token"`
		RepositoryOwner   string `json:"repository_owner" db:"repository_owner"`
		RepositoryName    string `json:"repository_name" db:"repository_name"`
		BranchName        string `json:"branch_name" db:"branch_name"`
		GithubProjectPath string `json:"github_project_path" db:"github_project_path"`
	}

	Ftp struct {
		FtpExecute   bool   `json:"ftp_execute" db:"ftp_execute"`
		FtpServer    string `json:"ftp_server" db:"ftp_server"`
		Username     string `json:"username" db:"username"`
		Password     string `json:"password" db:"password"`
		ProjectPath  string `json:"project_path" db:"project_path"`
		FtpDirectory string `json:"ftp_directory" db:"ftp_directory"`
	}
	StepsPath string  `json:"steps_path" db:"steps_path"`
	Steps     []Steps `json:"steps" db:"steps"`
}
