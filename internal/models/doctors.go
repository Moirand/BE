package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Doctors struct {
	db.ModelBase
	Id                   uuid.UUID  `json:"id" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId               uuid.UUID  `json:"user_id" column:"name:user_id;type:uuid;nullable:false"`
	SpecialityId         uuid.UUID  `json:"speciality_id" column:"name:speciality_id;type:uuid;nullable:false"`
	HealthcareFacilityId uuid.UUID  `json:"healthcare_facility_id" column:"name:healthcare_facility_id;type:uuid;nullable:false"`
	LicenseNumber        string     `json:"license_number" column:"name:license_number;type:varchar;nullable:false;unique"`
	Biography            *string    `json:"biography" column:"name:biography;type:text;nullable"`
	ConcultationFee      float64    `json:"concultation_fee" column:"name:concultation_fee;type:numeric;nullable:false"`
	ExperienceYears      int16      `json:"experience_years" column:"name:experience_years;type:smallint;nullable:false"`
	CreatedAt            time.Time  `json:"created_at" column:"name:created_at;type:timestamp;nullable:false;default:now()"`
	UpdatedAt            *time.Time `json:"updated_at" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	AppointmentDoctors                      []*Appointments       `json:"appointment_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	AppointmentsThroughAppointmentsDoctor   []*Appointments       `json:"appointments_through_appointments_doctor,omitempty" join:"joinType:manyToMany;through:appointments;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	AppointmentsThroughMedicalRecordsDoctor []*Appointments       `json:"appointments_through_medical_records_doctor,omitempty" join:"joinType:manyToMany;through:medical_records;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	DoctorScheduleDoctors                   []*DoctorSchedules    `json:"doctor_schedule_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	HealthcareFacilityHealthcare            *HealthcareFacilities `json:"healthcare_facility_healthcare,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:healthcare_facility_id"`
	MedicalRecordDoctors                    []*MedicalRecords     `json:"medical_record_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	SpecialityCategorySpeciality            *SpecialityCategories `json:"speciality_category_speciality,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:speciality_id"`
	User                                    *Users                `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
	UsersThroughAppointmentsDoctor          []*Users              `json:"users_through_appointments_doctor,omitempty" join:"joinType:manyToMany;through:appointments;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	UsersThroughMedicalRecordsDoctor        []*Users              `json:"users_through_medical_records_doctor,omitempty" join:"joinType:manyToMany;through:medical_records;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
}
