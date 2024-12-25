package models

import (
	"fmt"
	"home-utils/internal/repository"
	"time"
)

type GetGeyserStatusResponse struct {
	IsOn        bool   `json:"isOn"`
	UpdatedAt   string `json:"changeTime"`
	DisplayName string `json:"actionBy"`
}

type GeyserActionResponse struct {
	Success bool `json:"success"`
}

type GeyserActionRequest struct {
	TurnGeyserOn bool `json:"turnGeyserOn" `
}

func formatTimeString(t time.Time) string {
	year, month, day := t.Date()
	hour := t.Hour()
	minute := t.Minute()
	return fmt.Sprintf("%d %s %d %d:%d", day, month.String(), year, hour, minute)
}

func NewGetGeyserStatusResponse(row repository.GetGeyserStatusRow) GetGeyserStatusResponse {
	return GetGeyserStatusResponse{
		IsOn:        row.IsOn,
		DisplayName: row.DisplayName,
		UpdatedAt:   formatTimeString(row.UpdatedAt),
	}
}
