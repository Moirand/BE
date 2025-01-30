package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type DoctorSchedules struct {
	db.ModelBase
	Id          uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	DoctorId    uuid.UUID  `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable:false"`
	DayOfWeek   int32      `json:"day_of_week,omitempty" column:"name:day_of_week;type:integer;nullable:false"`
	StartTime   time.Time  `json:"start_time,omitempty" column:"name:start_time;type:time;nullable:false"`
	EndTime     time.Time  `json:"end_time,omitempty" column:"name:end_time;type:time;nullable:false"`
	QuotaPerDay int32      `json:"quota_per_day,omitempty" column:"name:quota_per_day;type:integer;nullable:false"`
	IsAvailable bool       `json:"is_available,omitempty" column:"name:is_available;type:boolean;nullable:false;default:true"`
	CreatedAt   time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:(now() AT TIME ZONE 'utc')"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctor_schedules" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor *Doctors `json:"doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
}
