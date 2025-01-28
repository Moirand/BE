package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type HealthcareFacilities struct {
	db.ModelBase
	Id          uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name        string     `json:"name,omitempty" column:"name:name;type:varchar;nullable:false"`
	AddressId   uuid.UUID  `json:"address_id,omitempty" column:"name:address_id;type:uuid;nullable:false"`
	PhoneNumber string     `json:"phone_number,omitempty" column:"name:phone_number;type:varchar;nullable:false;unique"`
	CreatedAt   time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable:false;default:now()"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"healthcare_facilities" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Address                                              *Addresses              `json:"address,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:address_id"`
	DoctorHealthcares                                    []*Doctors              `json:"doctor_healthcares,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:healthcare_facility_id"`
	SpecialityCategoriesThroughDoctorsHealthcareFacility []*SpecialityCategories `json:"speciality_categories_through_doctors_healthcare_facility,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:healthcare_facility_id;targetPrimaryKey:id;targetForeign:healthcare_facility_id"`
	UsersThroughDoctorsHealthcareFacility                []*Users                `json:"users_through_doctors_healthcare_facility,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:healthcare_facility_id;targetPrimaryKey:id;targetForeign:healthcare_facility_id"`
}
