package models

import (
	"fmt"
	"home-utils/internal/repository"
	"time"
)

type GetGeyserStatusResponse struct {
	IsOn         bool   `json:"isOn"`
	UpdatedAt    string `json:"changeTime"`
	DisplayName  string `json:"actionBy"`
	IsUserAction bool   `json:"isUserAction"`
}

type GeyserActionResponse struct {
	Success bool `json:"success"`
}

type GeyserActionRequest struct {
	TurnGeyserOn bool `json:"turnGeyserOn" `
}

type GetGeyserHistoryRequest struct {
	RowsPerPage int `form:"rowsPerPage"`
	PageNumber  int `form:"pageNumber"`
}

type GetGeyserHistoryResponse struct {
	NumRecord int                   `json:"numRecords"`
	Records   []GeyserHistoryRecord `json:"records"`
}

type GeyserHistoryRecord struct {
	Action   string `json:"action"`
	Resident string `json:"resident"`
	Time     string `json:"time"`
}

func formatTimeString(t time.Time) string {
	year, month, day := t.Date()
	hour := t.Hour()
	minute := t.Minute()
	return fmt.Sprintf("%d %s %d %d:%d", day, month.String(), year, hour, minute)
}

func NewGetGeyserStatusResponse(userKey string, row repository.GetGeyserStatusRow) GetGeyserStatusResponse {
	return GetGeyserStatusResponse{
		IsOn:         row.IsOn,
		UpdatedAt:    formatTimeString(row.UpdatedAt),
		DisplayName:  row.DisplayName,
		IsUserAction: row.Key == userKey,

	}
}

func NewGetGeyserHistoryResponse(rows []repository.GetGeyserHistoryPaginatedRow, numRows int) GetGeyserHistoryResponse {
	records := make([]GeyserHistoryRecord, len(rows))
	for i, row := range rows {
		records[i] = GeyserHistoryRecord{
			Action:   row.Action,
			Resident: row.DisplayName,
			Time:     formatTimeString(row.CreatedAt),
		}
	}
	return GetGeyserHistoryResponse{
		NumRecord: numRows,
		Records:   records,
	}
}
