package serializer

import (
	"github.com/ecumeurs/upsilonserializer/commessage"
	"github.com/ecumeurs/upsilonserializer/serializable"
)

type Serializer interface {
	Serialize(registry serializable.SerializableRegistry, cm commessage.ComMessage) ([]byte, error)
	Deserialize(registry serializable.SerializableRegistry, b []byte) (commessage.ComMessage, error)
}
