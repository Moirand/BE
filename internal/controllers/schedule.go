package controllers

import (
	"fmt"
	"medpoint/internal/models"
	"net/http"
	"strings"
	"time"

	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type ScheduleRequest struct {
	DoctorId string `path:"doctor_id"`
	Id       string `path:"schedule_id"`
	Body     models.DoctorSchedules
}

type ScheduleResponse struct {
	Id          string     `json:"id"`
	DoctorId    string     `json:"doctor_id"`
	DayOfWeek   int32      `json:"day_of_week"`
	StartTime   CustomTime `json:"start_time"`
	EndTime     CustomTime `json:"end_time"`
	QuotaPerDay int32      `json:"quota_per_day"`
	IsAvailable bool       `json:"is_available"`
}

// CustomTime adalah wrapper untuk time.Time agar bisa menggunakan format khusus
type CustomTime struct {
	time.Time
}

const timeFormat = "15:04:05" // Format sesuai database

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// Hilangkan tanda kutip dari string JSON
	str := strings.Trim(string(b), "\"")

	// Parse format waktu dari database
	t, err := time.Parse(timeFormat, str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	// Konversi time.Time ke format yang sesuai untuk JSON
	return []byte(fmt.Sprintf("\"%s\"", ct.Format(timeFormat))), nil
}

// path: /schedule
// type ScheduleController1 struct {
// 	raiden.ControllerBase
// 	Http    string `path:"/schedule" type:"custom"`
// 	Payload *ScheduleRequest
// 	Result  DoctorResponse
// }

// func (c *ScheduleController1) Post(ctx raiden.Context) error {
// 	// startTime, err := time.Parse(time.RFC3339, "2025-01-27T04:31:26.038927Z")
// 	// if err != nil {
// 	// 	return ctx.SendErrorWithCode(http.StatusBadRequest, fmt.Errorf("invalid start time format: %v", err))
// 	// }
//
// 	// endTime, err := time.Parse(time.RFC3339, "2026-01-27T04:31:26.038927Z")
// 	// if err != nil {
// 	// 	return ctx.SendErrorWithCode(http.StatusBadRequest, fmt.Errorf("invalid start time format: %v", err))
// 	// }
//
// 	payload := models.DoctorSchedules{
// 		DoctorId:    uuid.MustParse("776781d7-9a35-412a-98c3-d1ab17f94934"),
// 		DayOfWeek:   1,
// 		StartTime:   time.Now(),
// 		EndTime:     time.Now(),
// 		QuotaPerDay: 25,
// 		IsAvailable: true,
// 	}
//
// 	err := db.NewQuery(ctx).
// 		From(models.DoctorSchedules{}).
// 		Insert(&payload, nil)
//
// 	if err != nil {
// 		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
// 	}
//
// 	fmt.Print(payload)
//
// 	return nil
// }

// path: /schedule/{doctor_id}
type ScheduleController2 struct {
	raiden.ControllerBase
	Http    string `path:"/schedule/{doctor_id}" type:"custom"`
	Payload *ScheduleRequest
	Result  ScheduleResponse
}

func (c *ScheduleController2) Get(ctx raiden.Context) error {

	doctorScheduleList := []ScheduleResponse{c.Result}

	err := db.NewQuery(ctx).
		From(models.DoctorSchedules{}).
		Eq("doctor_id", c.Payload.DoctorId).
		Get(&doctorScheduleList)

	if err != nil {
		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
	}

	return ctx.SendJson(doctorScheduleList)
}

// path: /schedule/{doctor_id}/{schedule_id}
type ScheduleController3 struct {
	raiden.ControllerBase
	Http    string `path:"/schedule/{doctor_id}/{schedule_id}" type:"custom"`
	Payload *ScheduleRequest
	Result  ScheduleResponse
}

func (c *ScheduleController3) Get(ctx raiden.Context) error {

	err := db.NewQuery(ctx).
		From(models.DoctorSchedules{}).
		Eq("id", c.Payload.Id).
		Eq("doctor_id", c.Payload.DoctorId).
		Single(&c.Result)

	if err != nil {
		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
	}

	return ctx.SendJson(c.Result)
}

func (c *ScheduleController3) Delete(ctx raiden.Context) error {

	err := db.NewQuery(ctx).
		From(models.DoctorSchedules{}).
		Eq("id", c.Payload.Id).
		Eq("doctor_id", c.Payload.DoctorId).
		Delete()

	if err != nil {
		return ctx.SendErrorWithCode(http.StatusBadRequest, err)
	}

	return ctx.SendJson(c.Result)
}
