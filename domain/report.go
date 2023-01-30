package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Report struct {
	ID           int    `json:"id"`
	Location     string `json:"location"`
	Country      string `json:"country"`
	Quality      string `json:"quality"`
	ReporterID   string `json:"reporter_name"`
	ReporterName string `json:"reporter_id"`
	Date         int64  `json:"date"`
	Type         string `json:"type"`
	CreatedOn    int64  `json:"created_on"`
	ModifiedOn   int64  `json:"modified_on"`
}

type ReportRepository interface {
	CreateReport(user *User) error
	GetReport(data map[string]interface{}) (res Report, err error)
}

func (u *Report) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (u *Report) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}
