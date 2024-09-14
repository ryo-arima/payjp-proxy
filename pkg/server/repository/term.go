package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type TermRepository interface {
	GetTerms() []model.Terms
	CreateTerm(term model.Terms) *gorm.DB
	UpdateTerm(term model.Terms) *gorm.DB
	DeleteTerm(uuid string) *gorm.DB
}

type termRepository struct {
	BaseConfig config.BaseConfig
}


func (termRepository termRepository) GetTerms() []model.Terms {
	var terms []model.Terms
	termRepository.BaseConfig.DBConnection.Find(&terms)
	return terms
}

func (termRepository termRepository) CreateTerm(term model.Terms) *gorm.DB {
	results := termRepository.BaseConfig.DBConnection.Create(&term)
	return results
}

func (termRepository termRepository) UpdateTerm(term model.Terms) *gorm.DB {
	results := termRepository.BaseConfig.DBConnection.Model(model.Terms{}).Where("uuid = ?", term.UUID).Updates(term)
	return results
}

func (termRepository termRepository) DeleteTerm(uuid string) *gorm.DB {
	results := termRepository.BaseConfig.DBConnection.Model(model.Terms{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewTermRepository(conf config.BaseConfig) TermRepository {
	return &termRepository{BaseConfig: conf}
}