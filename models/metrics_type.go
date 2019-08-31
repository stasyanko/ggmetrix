package models

import (
	"github.com/jinzhu/gorm"
)

// Data data model
type MetricsType struct {
	gorm.Model
	Title string `sql:"type:varchar;"`
	Type  string `sql:"type:varchar;"`
}

func (MetricsType) TableName() string {
	return "ggmetrix_metrics_types"
}
