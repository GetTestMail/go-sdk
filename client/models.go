package client

import "time"

type Problem struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Status int    `json:"status,omitempty"`
}

type Attachment struct {
	Filename string `json:"filename,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
	Content  string `json:"content,omitempty"`
}

type Message struct {
	ID          string       `json:"id,omitempty"`
	From        string       `json:"from,omitempty"`
	To          string       `json:"to,omitempty"`
	Subject     string       `json:"subject,omitempty"`
	Text        string       `json:"text,omitempty"`
	HTML        string       `json:"html,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type GetTestMail struct {
	ID           string    `json:"id,omitempty"`
	EmailAddress string    `json:"emailAddress,omitempty"`
	ExpiresAt    time.Time `json:"expiresAt,omitempty"`
	Message      *Message  `json:"message,omitempty"`
}
