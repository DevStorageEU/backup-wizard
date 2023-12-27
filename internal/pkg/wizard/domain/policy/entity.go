package policy

import (
	"github.com/google/uuid"
	"time"
)

type RetentionKind string
type Weekday string

const (
	RetentionKindHourly  RetentionKind = "hourly"
	RetentionKindDaily   RetentionKind = "daily"
	RetentionKindWeekly  RetentionKind = "weekly"
	RetentionKindMonthly RetentionKind = "monthly"

	WeekdayMonday    Weekday = "monday"
	WeekdayTuesday   Weekday = "tuesday"
	WeekdayWednesday Weekday = "wednesday"
	WeekdayThursday  Weekday = "thursday"
	WeekdayFriday    Weekday = "friday"
	WeekdaySaturday  Weekday = "saturday"
	WeekdaySunday    Weekday = "sunday"
)

type Policy struct {
	ID         uuid.UUID
	Retentions []*Retention
	Schedules  []*Schedule
}

type Retention struct {
	PolicyID uuid.UUID
	Kind     RetentionKind
	Amount   uint
}

type Schedule struct {
	PolicyID uuid.UUID
	Kind     RetentionKind
	Weekday  Weekday
	Time     time.Duration
}
