package model

import "time"

type ZendeskTicket struct {
	URL         string    `json:"url"`
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Type        string    `json:"type"`
	Subject     string    `json:"subject"`
	RawSubject  string    `json:"raw_subject"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	GroupID     int       `json:"group_id"`
}

func (z ZendeskTicket) ShortenedSubject() string {
	if len(z.Subject) > 100 {
		return z.Subject[:97] + "..."
	}
	return z.Subject
}

func (z ZendeskTicket) GetGroupName() string {
	switch z.GroupID {
	case 24037409:
		return "Monitoring"
	case 20436884:
		return "Platform"
	default:
		return "Unknown"
	}
}
