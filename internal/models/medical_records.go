package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type MedicalRecords struct {
	db.ModelBase
	Id            uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId        uuid.UUID  `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	DoctorId      uuid.UUID  `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable:false"`
	AppointmentId uuid.UUID  `json:"appointment_id,omitempty" column:"name:appointment_id;type:uuid;nullable:false;unique"`
	Diagnose      string     `json:"diagnose,omitempty" column:"name:diagnose;type:varchar;nullable:false"`
	Treatment     string     `json:"treatment,omitempty" column:"name:treatment;type:varchar;nullable:false"`
	Prescription  string     `json:"prescription,omitempty" column:"name:prescription;type:text;nullable:false"`
	Notes         string     `json:"notes,omitempty" column:"name:notes;type:text;nullable:false"`
	CreatedAt     time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable:false;default:now()"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"medical_records" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Appointment *Appointments `json:"appointment,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:appointment_id"`
	Doctor      *Doctors      `json:"doctor,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	User        *Users        `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
