package services_domain_utils

import "time"

type TaskState int

const (
	// Indicates that the task is currently being processed by Handler.
	TaskStateActive TaskState = iota + 1

	// Indicates that the task is ready to be processed by Handler.
	TaskStatePending

	// Indicates that the task is scheduled to be processed some time in the future.
	TaskStateScheduled

	// Indicates that the task has previously failed and scheduled to be processed some time in the future.
	TaskStateRetry

	// Indicates that the task is archived and stored for inspection purposes.
	TaskStateArchived

	// Indicates that the task is processed successfully and retained until the retention TTL expires.
	TaskStateCompleted

	// Indicates that the task is waiting in a group to be aggregated into one task.
	TaskStateAggregating
)

func (s TaskState) String() string {
	switch s {
	case TaskStateActive:
		return "active"
	case TaskStatePending:
		return "pending"
	case TaskStateScheduled:
		return "scheduled"
	case TaskStateRetry:
		return "retry"
	case TaskStateArchived:
		return "archived"
	case TaskStateCompleted:
		return "completed"
	case TaskStateAggregating:
		return "aggregating"
	}
	panic("asynq: unknown task state")
}

type TaskInfo struct {
	// ID is the identifier of the task.
	ID string

	// Queue is the name of the queue in which the task belongs.
	Queue string

	// Type is the type name of the task.
	Type string

	// Payload is the payload data of the task.
	Payload []byte

	// State indicates the task state.
	State TaskState

	// MaxRetry is the maximum number of times the task can be retried.
	MaxRetry int

	// Retried is the number of times the task has retried so far.
	Retried int

	// LastErr is the error message from the last failure.
	LastErr string

	// LastFailedAt is the time time of the last failure if any.
	// If the task has no failures, LastFailedAt is zero time (i.e. time.Time{}).
	LastFailedAt time.Time

	// Timeout is the duration the task can be processed by Handler before being retried,
	// zero if not specified
	Timeout time.Duration

	// Deadline is the deadline for the task, zero value if not specified.
	Deadline time.Time

	// Group is the name of the group in which the task belongs.
	//
	// Tasks in the same queue can be grouped together by Group name and will be aggregated into one task
	// by a Server processing the queue.
	//
	// Empty string (default) indicates task does not belong to any groups, and no aggregation will be applied to the task.
	Group string

	// NextProcessAt is the time the task is scheduled to be processed,
	// zero if not applicable.
	NextProcessAt time.Time

	// IsOrphaned describes whether the task is left in active state with no worker processing it.
	// An orphaned task indicates that the worker has crashed or experienced network failures and was not able to
	// extend its lease on the task.
	//
	// This task will be recovered by running a server against the queue the task is in.
	// This field is only applicable to tasks with TaskStateActive.
	IsOrphaned bool

	// Retention is duration of the retention period after the task is successfully processed.
	Retention time.Duration

	// CompletedAt is the time when the task is processed successfully.
	// Zero value (i.e. time.Time{}) indicates no value.
	CompletedAt time.Time

	// Result holds the result data associated with the task.
	// Use ResultWriter to write result data from the Handler.
	Result []byte
}
