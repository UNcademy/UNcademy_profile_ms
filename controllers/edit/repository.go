package edit

import (
	model "UNcademy_profile_ms/models"
	"gorm.io/gorm"
)

// El input ya trae todos los campos que voy a modificar
type Repository interface {
	EditRepository(input *model.User, username string) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewEditRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) EditRepository(input *model.User, username string) (*model.User, string) {
	var users model.User
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("user_name = ?", username).Find(&users)

	if checkUserAccount.RowsAffected == 0 {
		errorCode <- "USER_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "USER_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	if input.DepDocument != "" {
		users.DepDocument = input.DepDocument
	}
	if input.CityDocument != "" {
		users.CityDocument = input.CityDocument
	}
	if input.Genre != "" {
		users.Genre = input.Genre
	}
	if input.Email != "" {
		users.Email = input.Email
	}
	if input.Cel != 0 {
		users.Cel = input.Cel
	}
	if input.Tel != 0 {
		users.Tel = input.Tel
	}
	if input.Age != 0 {
		users.Age = input.Age
	}
	if input.BirthPlace != "" {
		users.BirthPlace = input.BirthPlace
	}
	if input.Country != "" {
		users.Country = input.Country
	}
	if input.BloodType != "" {
		users.BloodType = input.BloodType
	}
	if input.Address != "" {
		users.Address = input.Address
	}
	if input.ArmyCard != false {
		users.ArmyCard = input.ArmyCard
	}
	if input.MotherFullName != "" {
		users.MotherFullName = input.MotherFullName
	}
	if input.MotherDocument != 0 {
		users.MotherDocument = input.MotherDocument
	}
	if input.FatherFullName != "" {
		users.FatherFullName = input.FatherFullName
	}
	if input.FatherDocument != 0 {
		users.FatherDocument = input.FatherDocument
	}

	changeFields := db.Save(&users)

	/*changeFields := db.Model(&users).Select("dep_document", "city_document", "genre",
	"email", "cel", "tel", "age", "birth_place", "country", "blood_type", "address", "army_card", "mother_full_name",
	"mother_document", "father_full_name", "father_document").Where("user_name = ?", username).Updates(model.User{DepDocument: users.DepDocument,
	CityDocument: input.CityDocument, Genre: input.Genre, Email: input.Email, Cel: input.Cel, Tel: input.Tel,
	Age: input.Age, BirthPlace: input.BirthPlace, Country: input.Country, BloodType: input.BloodType, Address: input.Address,
	ArmyCard: input.ArmyCard, MotherFullName: input.MotherFullName, MotherDocument: input.MotherDocument,
	FatherFullName: input.FatherFullName, FatherDocument: input.FatherDocument})*/

	if changeFields.Error != nil {
		errorCode <- "CHANGE_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
