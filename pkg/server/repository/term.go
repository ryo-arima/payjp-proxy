package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type TermRepository interface {
	GetTerms() []model.Terms
	CreateTerm(term model.Terms) model.Terms
	UpdateTerm(term model.Terms) model.Terms
	DeleteTerm(term model.Terms) model.Terms
}

type termRepository struct {
	BaseConfig config.BaseConfig
}

func (termRepository termRepository) GetTerms() []model.Terms {
	var terms []model.Terms
	return terms
}

func (termRepository termRepository) CreateTerm(term model.Terms) model.Terms {
	return term
}

func (termRepository termRepository) UpdateTerm(term model.Terms) model.Terms {
	return term
}

func (termRepository termRepository) DeleteTerm(term model.Terms) model.Terms {
	return term
}

func NewTermRepository(conf config.BaseConfig) TermRepository {
	return &termRepository{BaseConfig: conf}
}
