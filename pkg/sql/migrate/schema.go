package migrate

import (
	"cweb/http/type/model"

	"github.com/jinzhu/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.Menu{})
}
