package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Appointments struct {
	db.ModelBase
	Id                 uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId             uuid.UUID   `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	DoctorId           uuid.UUID   `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable:false"`
	AppointmentDate    time.Time   `json:"appointment_date,omitempty" column:"name:appointment_date;type:date;nullable:false"`
	RescheduleFrom     *uuid.UUID  `json:"reschedule_from,omitempty" column:"name:reschedule_from;type:uuid;nullable;unique"`
	CancellationReason *string     `json:"cancellation_reason,omitempty" column:"name:cancellation_reason;type:text;nullable"`
	Status             interface{} `json:"status,omitempty" column:"name:status;nullable:false"`
	CreatedAt          time.Time   `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable:false;default:now()"`
	UpdatedAt          *time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"appointments" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	AppointmentReschedule                    *Appointments     `json:"appointment_reschedule,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:reschedule_from"`
	Doctor                                   *Doctors          `json:"doctor,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	DoctorsThroughAppointmentsRescheduleFrom []*Doctors        `json:"doctors_through_appointments_reschedule_from,omitempty" join:"joinType:manyToMany;through:appointments;sourcePrimaryKey:id;sourceForeignKey:reschedule_from;targetPrimaryKey:id;targetForeign:reschedule_from"`
	DoctorsThroughMedicalRecordsAppointment  []*Doctors        `json:"doctors_through_medical_records_appointment,omitempty" join:"joinType:manyToMany;through:medical_records;sourcePrimaryKey:id;sourceForeignKey:appointment_id;targetPrimaryKey:id;targetForeign:appointment_id"`
	MedicalRecordAppointments                []*MedicalRecords `json:"medical_record_appointments,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:appointment_id"`
	PaymentAppointments                      []*Payments       `json:"payment_appointments,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:appointment_id"`
	User                                     *Users            `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
	UsersThroughAppointmentsRescheduleFrom   []*Users          `json:"users_through_appointments_reschedule_from,omitempty" join:"joinType:manyToMany;through:appointments;sourcePrimaryKey:id;sourceForeignKey:reschedule_from;targetPrimaryKey:id;targetForeign:reschedule_from"`
	UsersThroughMedicalRecordsAppointment    []*Users          `json:"users_through_medical_records_appointment,omitempty" join:"joinType:manyToMany;through:medical_records;sourcePrimaryKey:id;sourceForeignKey:appointment_id;targetPrimaryKey:id;targetForeign:appointment_id"`
}
