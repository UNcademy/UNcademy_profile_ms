package view

import (
	model "UNcademy_profile_ms/models"
	"gorm.io/gorm"
)

type Repository interface {
	ViewRepository(username string) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewViewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// Todos los repositorios tienen relacion directa con la base de datos
func (r *repository) ViewRepository(username string) (*model.User, string) {
	//Es el modelo de usuario que tenemos definidos pero nulos (estructura temporal)
	var users model.User
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.UserName = username

	//Miro si el nombre de usuario existe en la base de datos y guarda en users (estructura temporal) la informacion que esta guardada de ese usuario
	checkUserAccount := db.Select("*").Find(&users)

	if checkUserAccount.Error != nil {
		errorCode <- "USER_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "USER_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	return &users, <-errorCode
}
