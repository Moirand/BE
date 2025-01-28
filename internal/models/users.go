package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Users struct {
	db.ModelBase
	Id          uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	Username    string      `json:"username,omitempty" column:"name:username;type:varchar;nullable:false"`
	Email       string      `json:"email,omitempty" column:"name:email;type:varchar;nullable:false"`
	Password    string      `json:"password,omitempty" column:"name:password;type:varchar;nullable:false"`
	Sex         interface{} `json:"sex,omitempty" column:"name:sex;nullable:false"`
	Role        interface{} `json:"role,omitempty" column:"name:role;nullable:false"`
	DateOfBirth time.Time   `json:"date_of_birth,omitempty" column:"name:date_of_birth;type:date;nullable:false"`
	PhoneNumber *string     `json:"phone_number,omitempty" column:"name:phone_number;type:varchar;nullable;unique"`
	CreatedAt   time.Time   `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable:false;default:now()"`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"users" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	AppointmentUsers                       []*Appointments         `json:"appointment_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	AppointmentsThroughAppointmentsUser    []*Appointments         `json:"appointments_through_appointments_user,omitempty" join:"joinType:manyToMany;through:appointments;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	AppointmentsThroughMedicalRecordsUser  []*Appointments         `json:"appointments_through_medical_records_user,omitempty" join:"joinType:manyToMany;through:medical_records;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	DoctorUsers                            []*Doctors              `json:"doctor_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	DoctorsThroughAppointmentsUser         []*Doctors              `json:"doctors_through_appointments_user,omitempty" join:"joinType:manyToMany;through:appointments;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	DoctorsThroughMedicalRecordsUser       []*Doctors              `json:"doctors_through_medical_records_user,omitempty" join:"joinType:manyToMany;through:medical_records;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	HealthcareFacilitiesThroughDoctorsUser []*HealthcareFacilities `json:"healthcare_facilities_through_doctors_user,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	MedicalRecordUsers                     []*MedicalRecords       `json:"medical_record_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	SpecialityCategoriesThroughDoctorsUser []*SpecialityCategories `json:"speciality_categories_through_doctors_user,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	UserId                                 *Users                  `json:"user_id,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:id"`
}
