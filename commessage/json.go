package commessage

import (
	"encoding/json"
	"time"

	"github.com/ecumeurs/upsilonserializer/data"
	"github.com/ecumeurs/upsilonserializer/serializable"
	"github.com/google/uuid"
)

// SendComMessage is the message that is sent between the client and the server. Will be serialized
type JSONSendComMessage struct {
	ID           data.MyID
	EmitedAt     time.Time
	Target       json.RawMessage `json:"Target,omitempty"`
	Content      json.RawMessage `json:"Content,omitempty"`
	Meta         json.RawMessage `json:"Meta,omitempty"`
	TargetMethod string          // The method to call on the target
	ContentType  string          // struct type of content
	MetaType     string          // struct type of meta
	HasError     bool
	ErrorMessage string
	ErrorKey     string
}

func NewJSONMessage(msg ComMessage) JSONSendComMessage {
	if msg.HasError {
		return NewJSONReplyError(msg.ID.ToUUID(), msg.ErrorMessage, msg.ErrorKey, msg.Meta)
	}
	return NewJSONReply(msg.ID.ToUUID(), msg.Target, msg.Content, msg.Meta)
}

// NewSendCommMessage creates a new SendComMessage
func NewJSONSendCommMessage(target serializable.Serializable, content serializable.Serializable, meta serializable.Serializable) JSONSendComMessage {
	return JSONSendComMessage{
		ID:           data.NewId(),
		EmitedAt:     time.Now(),
		Target:       target.Serialize(),
		Content:      content.Serialize(),
		Meta:         meta.Serialize(),
		TargetMethod: target.GetType(),
		ContentType:  content.GetType(),
		MetaType:     meta.GetType(),
		HasError:     false,
		ErrorMessage: "",
		ErrorKey:     "",
	}
}

// NewReply creates a new SendComMessage
func NewJSONReply(id uuid.UUID, target serializable.Serializable, content serializable.Serializable, meta serializable.Serializable) JSONSendComMessage {
	return JSONSendComMessage{
		ID:           data.FromUUID(id),
		EmitedAt:     time.Now(),
		Target:       target.Serialize(),
		Content:      content.Serialize(),
		Meta:         meta.Serialize(),
		TargetMethod: target.GetType(),
		ContentType:  content.GetType(),
		MetaType:     meta.GetType(),
		HasError:     false,
		ErrorMessage: "",
		ErrorKey:     "",
	}
}

// NewReply creates a new SendComMessage
func NewJSONReplyError(id uuid.UUID, errorMsg string, errorKey string, meta serializable.Serializable) JSONSendComMessage {
	return JSONSendComMessage{
		ID:           data.FromUUID(id),
		EmitedAt:     time.Now(),
		Target:       nil,
		Content:      nil,
		Meta:         meta.Serialize(),
		TargetMethod: "",
		ContentType:  "",
		MetaType:     meta.GetType(),
		HasError:     true,
		ErrorMessage: errorMsg,
		ErrorKey:     errorKey,
	}
}
