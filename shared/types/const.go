package types

type EventType string

const (
	EventTypeSendToPlayer EventType = "send_to_ply"
	EvenetTypeBroadcast   EventType = "broadcast"
)
