package main

type EventStatus string

const (
	Scheduled EventStatus = "Scheduled"
	Started   EventStatus = "Started"
)

type ResourceType string

const (
	VirtualMachine ResourceType = "VirtualMachine"
)

type EventType string

const (
	Freeze    EventType = "Freeze"
	Reboot    EventType = "Reboot"
	Redeploy  EventType = "Redeploy"
	Preempt   EventType = "Preempt"
	Terminate EventType = "Terminate"
)

type EventSource string

const (
	Platform EventSource = "Platform"
	User     EventSource = "User"
)

type ScheduledEventsBatch struct {
	DocumentIncarnation int `json:"DocumentIncarnation"`
	Events              []struct {
		EventID           string       `json:"EventId"`
		EventStatus       EventStatus  `json:"EventStatus"`
		EventType         EventType    `json:"EventType"`
		ResourceType      ResourceType `json:"ResourceType"`
		Resources         []string     `json:"Resources"`
		NotBefore         string       `json:"NotBefore"`
		Description       string       `json:"Description"`
		EventSource       EventSource  `json:"EventSource"`
		DurationInSeconds int          `json:"DurationInSeconds"`
	} `json:"Events"`
}

type ConfirmScheduledEvent struct {
	StartRequests []StartRequest `json:"StartRequests"`
}

type StartRequest struct {
	EventID string `json:"EventId"`
}

type patchStringValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value bool   `json:"value"`
}
