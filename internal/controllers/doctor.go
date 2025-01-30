package controllers

import (
	"medpoint/internal/models"
	"net/http"

	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type DoctorRequest struct {
	Id   string `path:"doctor_id"`
	Body models.Doctors
}

type DoctorResponse struct {
	Id                   string  `json:"id"`
	UserId               string  `json:"user_id"`
	SpecialityId         string  `json:"speciality_id"`
	HealthcareFacilityId string  `json:"healthcare_facility_id"`
	LicenseNumber        string  `json:"license_number"`
	Biography            *string `json:"biography"`
	ConsultationFee      float64 `json:"concultation_fee"`
	ExperienceYears      int16   `json:"experience_years"`
}

// path: /doctor
type DoctorController1 struct {
	raiden.ControllerBase
	Http    string `path:"/doctor" type:"custom"`
	Model   models.Doctors
	Payload *DoctorRequest
	Result  DoctorResponse
}

//
// func (c *DoctorController1) Post(ctx raiden.Context) error {
// 	// err := resource.Import(&resource.Flags{UpdateStateOnly: true}, ctx.Config())
//
// 	bio := "asd"
// 	payload := models.Doctors{
// 		UserId:               uuid.MustParse("2b1f7246-cccb-4f01-a1c6-0e46da3a9060"),
// 		SpecialityId:         uuid.MustParse("4873e827-a5be-4336-a0f6-398e4a09113a"),
// 		HealthcareFacilityId: uuid.MustParse("19207e70-725c-4071-9ec9-7bbc33700e7e"),
// 		LicenseNumber:        "2b1f7246-cccb-4f01-a1c6-0e46da3a9061",
// 		Biography:            &bio,
// 		ConcultationFee:      50000,
// 		ExperienceYears:      10,
// 	}
//
// 	// payload := models.Doctors{
// 	// 	UserId:               c.Payload.Body.UserId,
// 	// 	SpecialityId:         c.Payload.Body.SpecialityId,
// 	// 	HealthcareFacilityId: c.Payload.Body.HealthcareFacilityId,
// 	// 	LicenseNumber:        c.Payload.Body.LicenseNumber,
// 	// 	Biography:            c.Payload.Body.Biography,
// 	// 	ConcultationFee:      c.Payload.Body.ConcultationFee,
// 	// 	ExperienceYears:      c.Payload.Body.ExperienceYears,
// 	// }
//
// 	err := db.NewQuery(ctx).
// 		From(c.Model).
// 		Insert(&payload, nil)
//
// 	if err != nil {
// 		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
// 	}
//
// 	fmt.Print(payload)
//
// 	return ctx.SendErrorWithCode(200, nil)
// }

func (c *DoctorController1) Get(ctx raiden.Context) error {

	doctorList := []DoctorResponse{c.Result}

	err := db.NewQuery(ctx).
		From(c.Model).
		Get(&doctorList)

	if err != nil {
		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
	}

	return ctx.SendJson(doctorList)
}

// path: /doctor/{doctor_id}
type DoctorController2 struct {
	raiden.ControllerBase
	Http    string `path:"/doctor/{doctor_id}" type:"custom"`
	Payload *DoctorRequest
	Result  DoctorResponse
}

func (c *DoctorController2) Get(ctx raiden.Context) error {

	err := db.NewQuery(ctx).
		From(models.Doctors{}).
		Eq("id", c.Payload.Id).
		Single(&c.Result)

	if err != nil {
		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
	}

	return ctx.SendJson(c.Result)
}

// func (c *DoctorController2) Put(ctx raiden.Context) error {
//
// 	payload := models.Doctors{
// 		UserId:               c.Payload.Body.UserId,
// 		SpecialityId:         c.Payload.Body.SpecialityId,
// 		HealthcareFacilityId: c.Payload.Body.HealthcareFacilityId,
// 		LicenseNumber:        c.Payload.Body.LicenseNumber,
// 		Biography:            c.Payload.Body.Biography,
// 		ConcultationFee:      c.Payload.Body.ConcultationFee,
// 		ExperienceYears:      c.Payload.Body.ExperienceYears,
// 	}
//
// 	err := db.
// 		NewQuery(ctx).
// 		From(models.Doctors{}).
// 		Eq("id", c.Payload.Id).
// 		Update(payload, nil)
//
// 	if err != nil {
// 		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
// 	}
//
// 	return ctx.SendJson(c.Result)
// }

// func (c *DoctorController2) Delete(ctx raiden.Context) error {
//
// 	err := db.NewQuery(ctx).
// 		From(models.Doctors{}).
// 		Eq("id", c.Payload.Id).
// 		Delete()
//
// 	if err != nil {
// 		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
// 	}
//
// 	return ctx.SendJson(c.Result)
// }
