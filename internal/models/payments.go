package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Payments struct {
	db.ModelBase
	Id            uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	AppointmentId uuid.UUID  `json:"appointment_id,omitempty" column:"name:appointment_id;type:uuid;nullable:false;unique"`
	Amount        float64    `json:"amount,omitempty" column:"name:amount;type:numeric;nullable:false"`
	Method        string     `json:"method,omitempty" column:"name:method;type:text;nullable:false"`
	Status        string     `json:"status,omitempty" column:"name:status;type:text;nullable:false"`
	RefundStatus  *string    `json:"refund_status,omitempty" column:"name:refund_status;type:text;nullable"`
	CreatedAt     time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable:false;default:now()"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"payments" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Appointment *Appointments `json:"appointments,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:appointment_id"`
}
