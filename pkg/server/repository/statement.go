package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type StatementRepository interface {
	GetStatements() []model.Statements
	CreateStatement(statement model.Statements) *gorm.DB
	UpdateStatement(statement model.Statements) *gorm.DB
	DeleteStatement(uuid string) *gorm.DB
}

type statementRepository struct {
	BaseConfig config.BaseConfig
}


func (statementRepository statementRepository) GetStatements() []model.Statements {
	var statements []model.Statements
	statementRepository.BaseConfig.DBConnection.Find(&statements)
	return statements
}

func (statementRepository statementRepository) CreateStatement(statement model.Statements) *gorm.DB {
	results := statementRepository.BaseConfig.DBConnection.Create(&statement)
	return results
}

func (statementRepository statementRepository) UpdateStatement(statement model.Statements) *gorm.DB {
	results := statementRepository.BaseConfig.DBConnection.Model(model.Statements{}).Where("uuid = ?", statement.UUID).Updates(statement)
	return results
}

func (statementRepository statementRepository) DeleteStatement(uuid string) *gorm.DB {
	results := statementRepository.BaseConfig.DBConnection.Model(model.Statements{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewStatementRepository(conf config.BaseConfig) StatementRepository {
	return &statementRepository{BaseConfig: conf}
}