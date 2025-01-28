package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Addresses struct {
	db.ModelBase
	Id         uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Street     string     `json:"street,omitempty" column:"name:street;type:varchar;nullable:false"`
	District   string     `json:"district,omitempty" column:"name:district;type:varchar;nullable:false"`
	City       string     `json:"city,omitempty" column:"name:city;type:varchar;nullable:false"`
	PostalCode *string    `json:"postal_code,omitempty" column:"name:postal_code;type:varchar;nullable"`
	CreatedAt  time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable:false;default:now()"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"addresses" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	HealthcareFacilityAddresses []*HealthcareFacilities `json:"healthcare_facility_addresses,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:address_id"`
}
