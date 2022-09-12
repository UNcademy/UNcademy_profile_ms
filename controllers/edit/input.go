package edit

type InputEdit struct {
	DepDocument    string `json:"dep_document" validate:"lowercase"`
	CityDocument   string `json:"city_document" validate:"lowercase"`
	Genre          string `json:"genre" validate:"lowercase"`
	Email          string `json:"email" validate:"email"`
	BirthPlace     string `json:"birthday_place" validate:"lowercase"`
	Cel            int    `json:"cel" validate:"numeric"`
	Tel            int    `json:"tel" validate:"numeric"`
	Age            int    `json:"age" validate:"numeric"`
	Country        string `json:"country" validate:"lowercase"`
	BloodType      string `json:"blood_type" validate:"lowercase"`
	Address        string `json:"address"`
	ArmyCard       bool   `json:"army_card" validate:"boolean"`
	MotherFullName string `json:"mother_full_name" validate:"lowercase"`
	MotherDocument int    `json:"mother_document" validate:"numeric"`
	FatherFullName string `json:"father_full_name" validate:"lowercase"`
	FatherDocument int    `json:"father_document" validate:"numeric"`
}
