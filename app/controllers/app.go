package controllers

import (
	"github.com/fasilz/co-memo/app/models"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
)

type App struct {
	gormc.TxnController
	CurrentUser *models.User
}
