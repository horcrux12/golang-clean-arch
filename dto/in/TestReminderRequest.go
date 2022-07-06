package in

import "time"

type ToDoListRequest struct {
	ID                 int64     `json:"id"`
	UUIDKey            string    `json:"uuid_key"`
	ListID             int64     `json:"list_id"`
	CustomerID         int64     `json:"customer_id"`
	RepeatUntilDateStr string    `json:"repeat_until_date"`
	RepeatUntilDate    time.Time `json:"-"`
	RepeatFromDateStr  string    `json:"repeat_from_date"`
	RepeatFromDate     time.Time `json:"-"`
	TaskName           string    `json:"task_name"`
	Email              string    `json:"email"`
	Description        string    `json:"description"`
	DueDateStr         string    `json:"due_date"`
	DueDate            time.Time `json:"-"`
	ReminderDateStr    string    `json:"reminder_date"`
	ReminderDate       time.Time `json:"-"`
	IsImportant        bool      `json:"is_important"`
	RepeatType         string    `json:"repeat_type"`
	RepeatEvery        int64     `json:"repeat_every"`
	CreatedClient      string    `json:"created_client"`
	UpdatedClient      string    `json:"updated_client"`
	CreatedBy          int64     `json:"created_by"`
	UpdatedBy          int64     `json:"updated_by"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAtStr       string    `json:"updated_at"`
	UpdatedAt          time.Time `json:"-"`
}
