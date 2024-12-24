package model

import "time"

// IssueCategory represents the type of issue
type IssueCategory string

const (
	TaskCategory     IssueCategory = "task"
	StoryCategory    IssueCategory = "story"
	QuestionCategory IssueCategory = "question"
	BugCategory      IssueCategory = "bug"
	CRMCategory      IssueCategory = "crm"
)

// IssueStatus represents the current status of an issue
type IssueStatus string

const (
	NewStatus      IssueStatus = "new_status"
	TodoStatus     IssueStatus = "todo"
	InProgress     IssueStatus = "in_progress"
	Completed      IssueStatus = "completed"
	QuestionStatus IssueStatus = "question"
	BlockedStatus  IssueStatus = "blocked"
	VerifiedStatus IssueStatus = "verified"
)

// IssuePriority represents the priority level of an issue
type IssuePriority string

const (
	LowPriority      IssuePriority = "low"
	MediumPriority   IssuePriority = "medium"
	HighPriority     IssuePriority = "high"
	CriticalPriority IssuePriority = "critical"
)

// IssueSeverity represents the severity level of an issue
type IssueSeverity string

const (
	NormalSeverity   IssueSeverity = "normal"
	HighSeverity     IssueSeverity = "high"
	SeriousSeverity  IssueSeverity = "serious"
	DisasterSeverity IssueSeverity = "disaster"
)

// TimeEntry represents a time tracking entry for an issue
type TimeEntry struct {
	StartTime   time.Time  `json:"start_time"`
	EndTime     *time.Time `json:"end_time,omitempty"`
	Description string     `json:"description"`
	TimeSpent   float64    `json:"time_spent" description:"Time spent in hours"`
	UserID      string     `json:"user_id"`
}

// Issue represents a task, bug, or other trackable item
type Issue struct {
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	Category       IssueCategory `json:"category" binding:"required"`
	Status         IssueStatus   `json:"status" binding:"required"`
	Priority       IssuePriority `json:"priority" binding:"required"`
	Severity       IssueSeverity `json:"severity" binding:"required"`
	DueDate        time.Time     `json:"due_date"`
	Assignees      []string      `json:"assignees"`
	Labels         []string      `json:"labels"`
	EstimatedHours *float64      `json:"estimated_hours,omitempty"`
	ActualHours    *float64      `json:"actual_hours,omitempty"`
	Dependencies   []string      `json:"dependencies"`
	StartTime      time.Time     `json:"start_time"`
	EndTime        time.Time     `json:"end_time"`
	TimeEntries    []TimeEntry   `json:"time_entries"`
	TotalTimeSpent float64       `json:"total_time_spent" description:"Total time spent in hours"`
	ProjectID      string        `json:"project_id"`
}
