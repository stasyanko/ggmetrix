package models

import (
	"github.com/jinzhu/gorm"
)

// Data data model
type Data struct {
	gorm.Model
	Title  string `sql:"type:varchar;"`
	Type   string `sql:"type:varchar;"`
	Value  uint   `sql:"type:integer;"`
	UnixTs uint   `sql:"type:bigint;"`
}

func (Data) TableName() string {
	return "ggmetrix_data"
}
