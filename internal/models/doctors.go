package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Doctors struct {
	db.ModelBase
	Id                   uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId               uuid.UUID  `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	SpecialityId         uuid.UUID  `json:"speciality_id,omitempty" column:"name:speciality_id;type:uuid;nullable:false"`
	HealthcareFacilityId uuid.UUID  `json:"healthcare_facility_id,omitempty" column:"name:healthcare_facility_id;type:uuid;nullable:false"`
	LicenseNumber        string     `json:"license_number,omitempty" column:"name:license_number;type:varchar;nullable:false;unique"`
	Biography            *string    `json:"biography,omitempty" column:"name:biography;type:text;nullable"`
	ConcultationFee      float64    `json:"concultation_fee,omitempty" column:"name:concultation_fee;type:numeric;nullable:false"`
	ExperienceYears      int16      `json:"experience_years,omitempty" column:"name:experience_years;type:smallint;nullable:false"`
	CreatedAt            time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:(now() AT TIME ZONE 'utc')"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	AppointmentDoctors                      []*Appointments       `json:"appointment_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	AppointmentsThroughAppointmentsDoctor   []*Appointments       `json:"appointments_through_appointments_doctor,omitempty" join:"joinType:manyToMany;through:appointments;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	AppointmentsThroughMedicalRecordsDoctor []*Appointments       `json:"appointments_through_medical_records_doctor,omitempty" join:"joinType:manyToMany;through:medical_records;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	DoctorScheduleDoctors                   []*DoctorSchedules    `json:"doctor_schedule_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	HealthcareFacilityHealthcare            *HealthcareFacilities `json:"healthcare_facilities,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:healthcare_facility_id"`
	MedicalRecordDoctors                    []*MedicalRecords     `json:"medical_record_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	SpecialityCategorySpeciality            *SpecialityCategories `json:"speciality_categories,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:speciality_id"`
	User                                    *Users                `json:"users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
	UsersThroughAppointmentsDoctor          []*Users              `json:"users_through_appointments_doctor,omitempty" join:"joinType:manyToMany;through:appointments;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	UsersThroughMedicalRecordsDoctor        []*Users              `json:"users_through_medical_records_doctor,omitempty" join:"joinType:manyToMany;through:medical_records;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
}
