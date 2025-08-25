package models

import "time"

type CommandResult struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	BatchID          int    `json:"batchId"`
	CategoryID       int    `json:"categoryId"`
	EntityID         int    `json:"entityId"`
	CurrentActionIdx int    `json:"currentActionIdx"`
	ExecuteAt        int64  `json:"executeAt"`
	LastActivityAt   int64  `json:"lastActivityAt"`
	FinishedAt       int64  `json:"finishedAt"`
	ResultCode       int    `json:"resultCode"`
	ResultText       string `json:"resultText"`
}

func NewCommand() *CommandResult {
	timestamp := time.Now().Unix()
	return &CommandResult{
		ID:               8900,
		Name:             "command",
		BatchID:          8900,
		CategoryID:       8900,
		EntityID:         8900,
		CurrentActionIdx: 8900,
		ExecuteAt:        timestamp,
		LastActivityAt:   timestamp,
		FinishedAt:       timestamp,
		ResultCode:       127,
		ResultText:       "command",
	}
}
