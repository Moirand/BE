package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type SpecialityCategories struct {
	db.ModelBase
	Id          uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name        string     `json:"name,omitempty" column:"name:name;type:varchar;nullable:false;unique"`
	Description string     `json:"description,omitempty" column:"name:description;type:varchar;nullable:false"`
	CreatedAt   time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable:false;default:now()"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"speciality_categories" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorSpecialities                           []*Doctors              `json:"doctor_specialities,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:speciality_id"`
	HealthcareFacilitiesThroughDoctorsSpeciality []*HealthcareFacilities `json:"healthcare_facilities_through_doctors_speciality,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:speciality_id;targetPrimaryKey:id;targetForeign:speciality_id"`
	UsersThroughDoctorsSpeciality                []*Users                `json:"users_through_doctors_speciality,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:speciality_id;targetPrimaryKey:id;targetForeign:speciality_id"`
}
