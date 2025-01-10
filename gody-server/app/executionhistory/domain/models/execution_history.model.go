package models

type Step struct {
	ExecutionName   string `json:"execution_name" db:"execution_name"`
	ExecutionTime   string `json:"execution_time" db:"execution_time"`
	ExecutionDate   string `json:"execution_date" db:"execution_date"`
	ExecutionStatus string `json:"execution_status" db:"execution_status"`
	ExecutionError  string `json:"execution_error" db:"execution_error"`
}

type Executionhistory struct {
	ID        uint   `json:"id" db:"id" gorm:"primaryKey"`
	ActionID  int    `json:"action_id" db:"action_id"`
	Step      string `json:"step" db:"step" gorm:"type:jsonb"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
