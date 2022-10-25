package commessage

import (
	"time"

	"github.com/ecumeurs/upsilonserializer/data"
	"github.com/ecumeurs/upsilonserializer/serializable"
	"github.com/google/uuid"
)

// Message received from the client, not serialized, THIS one will be used by this APP
type ComMessage struct {
	ID           data.MyID
	EmitedAt     time.Time
	Target       serializable.Serializable
	Content      serializable.Serializable
	Meta         serializable.Serializable
	HasError     bool
	ErrorMessage string
	ErrorKey     string
}

// NewComMessage creates a new ComMessage
func NewComMessage(target serializable.Serializable, content serializable.Serializable, meta serializable.Serializable) ComMessage {
	return ComMessage{
		ID:           data.NewId(),
		EmitedAt:     time.Now(),
		Target:       target,
		Content:      content,
		Meta:         meta,
		HasError:     false,
		ErrorMessage: "",
		ErrorKey:     "",
	}
}

func NewReply(id uuid.UUID, target serializable.Serializable, content serializable.Serializable, meta serializable.Serializable) ComMessage {
	return ComMessage{
		ID:           data.FromUUID(id),
		EmitedAt:     time.Now(),
		Target:       target,
		Content:      content,
		Meta:         meta,
		HasError:     false,
		ErrorMessage: "",
		ErrorKey:     "",
	}
}

// NewComMessageError creates a new ComMessage with an error
func NewComMessageError(id uuid.UUID, errorMessage string, errorKey string) ComMessage {
	return ComMessage{
		ID:           data.FromUUID(id),
		EmitedAt:     time.Now(),
		Target:       nil,
		Content:      nil,
		Meta:         nil,
		HasError:     true,
		ErrorMessage: errorMessage,
		ErrorKey:     errorKey,
	}
}
