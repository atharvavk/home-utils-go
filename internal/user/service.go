package user

import (
	"context"
	"database/sql"
	"fmt"
	"home-utils/internal/models"
	"home-utils/internal/repository"

	"go.uber.org/zap"
)

type UserService struct {
	logger  *zap.Logger
	sqlConn *sql.DB
}

func NewUserService(logger *zap.Logger, sqlConn *sql.DB) UserService {
	return UserService{
		logger:  logger,
		sqlConn: sqlConn,
	}
}

func (us UserService) GetUser(key string) (models.GetResidentResponse, error) {
	repo := repository.New(us.sqlConn)
	ctx := context.TODO()
	resident, err := repo.GetResidentByKey(ctx, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.GetResidentResponse{}, models.NewBadReqError(models.NO_SUCH_USER)
		}
		us.logger.Error("INTERNAL SERVER ERROR", zap.String("error", err.Error()))
		return models.GetResidentResponse{}, fmt.Errorf("Error while getting resident: %s", err.Error())
	}
	return models.GetResidentResponse{DisplayName: resident.DisplayName}, nil
}
