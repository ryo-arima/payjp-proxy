package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type StatementRepository interface {
	GetStatements() []model.Statements
	CreateStatement(statement model.Statements) model.Statements
	UpdateStatement(statement model.Statements) model.Statements
	DeleteStatement(statement model.Statements) model.Statements
}

type statementRepository struct {
	BaseConfig config.BaseConfig
}

func (statementRepository statementRepository) GetStatements() []model.Statements {
	var statements []model.Statements
	return statements
}

func (statementRepository statementRepository) CreateStatement(statement model.Statements) model.Statements {
	return statement
}

func (statementRepository statementRepository) UpdateStatement(statement model.Statements) model.Statements {
	return statement
}

func (statementRepository statementRepository) DeleteStatement(statement model.Statements) model.Statements {
	return statement
}

func NewStatementRepository(conf config.BaseConfig) StatementRepository {
	return &statementRepository{BaseConfig: conf}
}
