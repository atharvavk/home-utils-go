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
    if !isActionOn{
      errorCode = models.GEYSER_OFF_OR_INVALID_USER
    }
    return models.GeyserActionResponse{}, models.NewBadReqError(errorCode)
  }
  return models.GeyserActionResponse{Success: true}, nil
}

func (gs GeyserService) GetGeyserStatus() (models.GetGeyserStatusResponse, error) {
	repo := repository.New(gs.sqlConn)
	status, err := repo.GetGeyserStatus(context.TODO())
	if err != nil {
		gs.logger.Error("INTERNAL SERVER ERROR", zap.String("error", err.Error()))
		return models.GetGeyserStatusResponse{}, err
	}
	return models.NewGetGeyserStatusResponse(status), nil
}

func NewGeyserService(logger *zap.Logger, sqlConn *sql.DB) GeyserService {
	return GeyserService{
		logger:  logger,
		sqlConn: sqlConn,
	}
}
