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
	ConsultationFee      float64 `json:"consultation_fee"`
	ExperienceYears      int16   `json:"experience_years"`
}

type DoctorController1 struct {
	raiden.ControllerBase
	Http    string `path:"/doctor/{doctor_id}" type:"custom"`
	Payload *DoctorRequest
	Result  DoctorResponse
}

func (c *DoctorController1) Get(ctx raiden.Context) error {

	err := db.NewQuery(ctx).
		From(models.Doctors{}).
		Eq("id", c.Payload.Id).
		Single(&c.Result)

	if err != nil {
		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
	}

	return ctx.SendJson(c.Result)
}

// func (c *DoctorController1) Put(ctx raiden.Context) error {

// 	payload := models.Doctors{
// 		UserId:               c.Payload.Body.UserId,
// 		SpecialityId:         c.Payload.Body.SpecialityId,
// 		HealthcareFacilityId: c.Payload.Body.HealthcareFacilityId,
// 		LicenseNumber:        c.Payload.Body.LicenseNumber,
// 		Biography:            c.Payload.Body.Biography,
// 		ConcultationFee:      c.Payload.Body.ConcultationFee,
// 		ExperienceYears:      c.Payload.Body.ExperienceYears,
// 	}

// 	err := db.
// 		NewQuery(ctx).
// 		From(models.Doctors{}).
// 		Eq("id", c.Payload.Id).
// 		Update(payload, nil)

// 	if err != nil {
// 		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
// 	}

// 	return ctx.SendJson(c.Result)
// }

// func (c *DoctorController1) Delete(ctx raiden.Context) error {

// 	err := db.
// 		NewQuery(ctx).
// 		From(c.Model).
// 		Eq("id", c.Payload.Id).
// 		Delete()

// 	if err != nil {
// 		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
// 	}

// 	return ctx.SendJson(c.Result)
// }

type DoctorController2 struct {
	raiden.ControllerBase
	Http    string `path:"/doctor" type:"custom"`
	Payload *DoctorRequest
	Result  DoctorResponse
}

// func (c *DoctorController2) Post(ctx raiden.Context) error {
// 	if err := resource.Import(&resource.Flags{UpdateStateOnly: true}, ctx.Config()); err != nil {
// 		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
// 	}
//
// 	return ctx.SendJson(c.Result)
// }

func (c *DoctorController2) Get(ctx raiden.Context) error {

	doctorList := []DoctorResponse{c.Result}

	err := db.
		NewQuery(ctx).
		From(models.Doctors{}).
		Get(&doctorList)

	if err != nil {
		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
	}

	return ctx.SendJson(doctorList)
}
