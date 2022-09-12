package edit

import model "UNcademy_profile_ms/models"

type Service interface {
	EditService(input *InputEdit, username string) string
}

type service struct {
	repository Repository
}

func NewEditService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) EditService(input *InputEdit, username string) string {
	user := model.User{
		DepDocument:    input.DepDocument,
		CityDocument:   input.CityDocument,
		Genre:          input.Genre,
		Email:          input.Email,
		Cel:            input.Cel,
		Tel:            input.Tel,
		Age:            input.Age,
		BirthPlace:     input.BirthPlace,
		Country:        input.Country,
		BloodType:      input.BloodType,
		Address:        input.Address,
		ArmyCard:       input.ArmyCard,
		MotherFullName: input.MotherFullName,
		MotherDocument: input.MotherDocument,
		FatherFullName: input.FatherFullName,
		FatherDocument: input.FatherDocument,
	}

	//--
	_, errEdit := s.repository.EditRepository(&user, username)

	return errEdit
}
