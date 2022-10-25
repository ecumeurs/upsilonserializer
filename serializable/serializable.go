package serializable

import "github.com/ecumeurs/upsilonserializer/documentable"

type Serializable interface {
	// Return the type of the object as a string
	GetType() string
	Serialize() []byte
	Deserialize([]byte) error

	// validate provided object (and subs)
	IsValid() (bool, string)

	GetDocumentation() documentable.Documentable
}

type SerializableGenerator interface {
	// Create creates a new instance of a serializable object
	Create() Serializable
	GetType() string
}

// SerializableRegistry is a registry of serializable objects. It allow creation of a struct based on name.
type SerializableRegistry struct {
	registry map[string]SerializableGenerator
}

// NewSerializableRegistry creates a new SerializableRegistry
func NewSerializableRegistry() SerializableRegistry {
	return SerializableRegistry{
		registry: make(map[string]SerializableGenerator),
	}
}

// Register registers a serializable object in the registry
func (sr *SerializableRegistry) Register(s SerializableGenerator) {
	sr.registry[s.GetType()] = s
}

// Create creates a new instance of a serializable object based on its type
func (sr *SerializableRegistry) Create(t string) Serializable {
	return sr.registry[t].Create()
}
