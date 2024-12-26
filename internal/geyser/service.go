package geyser

import (
	"context"
	"database/sql"
	"home-utils/internal/models"
	"home-utils/internal/repository"

	"go.uber.org/zap"
)

type GeyserService struct {
	logger  *zap.Logger
	sqlConn *sql.DB
}

func (gs GeyserService) GetGeyserHistory(req models.GetGeyserHistoryRequest) (models.GetGeyserHistoryResponse, error) {
	repo := repository.New(gs.sqlConn)
	sqlOffset := (req.PageNumber - 1) * req.RowsPerPage
	rows, err := repo.GetGeyserHistoryPaginated(context.TODO(), repository.GetGeyserHistoryPaginatedParams{
		Limit:  int32(req.RowsPerPage),
		Offset: int32(sqlOffset),
	})
	if err != nil {
		gs.logger.Error("INTERNAL SERVER ERROR", zap.String("error", err.Error()))
		return models.GetGeyserHistoryResponse{}, err
	}
  count, err := repo.GetGeyserHistoryCount(context.TODO())
  if err != nil {
		gs.logger.Error("INTERNAL SERVER ERROR", zap.String("error", err.Error()))
    return models.GetGeyserHistoryResponse{}, err
  }
	return models.NewGetGeyserHistoryResponse(rows, int(count)), nil
}

func (gs GeyserService) insertGeyserHistory(isActionOn bool, userKey string) {
	actionString := "TURN_ON"
	if !isActionOn {
		actionString = "TURN_OFF"
	}
	repo := repository.New(gs.sqlConn)
	if err := repo.InsertGeyserHistory(context.TODO(), repository.InsertGeyserHistoryParams{Userkey: userKey, Actionvalue: actionString}); err != nil {
		gs.logger.Error("INTERNAL SERVER ERROR", zap.String("error", err.Error()))
	}
}

func (gs GeyserService) DoGeyserAction(userKey string, isActionOn bool) (models.GeyserActionResponse, error) {
	repo := repository.New(gs.sqlConn)
	var rowsUpdated int64
	var err error
	if isActionOn {
		rowsUpdated, err = repo.TurnOnGeyser(context.TODO(), userKey)
	} else {
		rowsUpdated, err = repo.TurnOffGeyser(context.TODO(), userKey)
	}
	if err != nil {
		gs.logger.Error("INTERNAL SERVER ERROR", zap.String("error", err.Error()))
		return models.GeyserActionResponse{}, err
	}
	if rowsUpdated == 0 {
		var errorCode models.ErrorCode = models.GEYSER_ALREADY_ON
		if !isActionOn {
			errorCode = models.GEYSER_OFF_OR_INVALID_USER
		}
		return models.GeyserActionResponse{}, models.NewBadReqError(errorCode)
	}
	gs.insertGeyserHistory(isActionOn, userKey)
	return models.GeyserActionResponse{Success: true}, nil
}

func (gs GeyserService) GetGeyserStatus(userKey string) (models.GetGeyserStatusResponse, error) {
	repo := repository.New(gs.sqlConn)
	status, err := repo.GetGeyserStatus(context.TODO())
	if err != nil {
		gs.logger.Error("INTERNAL SERVER ERROR", zap.String("error", err.Error()))
		return models.GetGeyserStatusResponse{}, err
	}
	return models.NewGetGeyserStatusResponse(userKey, status), nil
}

func NewGeyserService(logger *zap.Logger, sqlConn *sql.DB) GeyserService {
	return GeyserService{
		logger:  logger,
		sqlConn: sqlConn,
	}
}
