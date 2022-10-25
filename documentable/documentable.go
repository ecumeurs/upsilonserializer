package documentable

import (
	"fmt"
	"time"
)

type Documentable interface {
	// Return the type of the object as a string
	GetType() string
	// For ArrayDocumentable, return the type of the array
	GetSubType() string
	// Return the documentation of the field
	GetDocumentation() string
	// Is the field required
	IsRequired() bool
	// Validation
	GetValidationRequirements() string

	GetDefaultValue() string

	// Get children documentation
	GetChildrenDocumentation() []Documentable
}

type IntDocumentable struct {
	Documentation string
	Required      bool
	Validation    string
	DefaultValue  int
}

// MakeIntDocumentable
func MakeIntDocumentable(doc string, required bool, validation string, defaultValue int) IntDocumentable {
	return IntDocumentable{
		Documentation: doc,
		Required:      required,
		Validation:    validation,
		DefaultValue:  defaultValue,
	}
}

func (i IntDocumentable) GetType() string {
	return "int"
}

func (i IntDocumentable) GetSubType() string {
	return ""
}

func (i IntDocumentable) GetDocumentation() string {
	return i.Documentation
}

func (i IntDocumentable) IsRequired() bool {
	return i.Required
}

func (i IntDocumentable) GetValidationRequirements() string {
	return i.Validation
}

func (i IntDocumentable) GetChildrenDocumentation() []Documentable {
	return []Documentable{}
}

func (i IntDocumentable) GetDefaultValue() string {
	return fmt.Sprintf("%d", i.DefaultValue)
}

type StringDocumentable struct {
	Documentation string
	Required      bool
	Validation    string
	DefaultValue  string
}

// MakeStringDocumentable
func MakeStringDocumentable(doc string, required bool, validation string, defaultValue string) StringDocumentable {
	return StringDocumentable{
		Documentation: doc,
		Required:      required,
		Validation:    validation,
		DefaultValue:  defaultValue,
	}
}

func (s StringDocumentable) GetType() string {
	return "string"
}

func (s StringDocumentable) GetSubType() string {
	return ""
}

func (s StringDocumentable) GetDocumentation() string {
	return s.Documentation
}

func (s StringDocumentable) IsRequired() bool {
	return s.Required
}

func (s StringDocumentable) GetValidationRequirements() string {
	return s.Validation
}

func (s StringDocumentable) GetChildrenDocumentation() []Documentable {
	return []Documentable{}
}

func (s StringDocumentable) GetDefaultValue() string {
	return s.DefaultValue
}

type BoolDocumentable struct {
	Documentation string
	Required      bool
	Validation    string
	DefaultValue  bool
}

// MakeBoolDocumentable
func MakeBoolDocumentable(doc string, required bool, validation string, defaultValue bool) BoolDocumentable {
	return BoolDocumentable{
		Documentation: doc,
		Required:      required,
		Validation:    validation,
		DefaultValue:  defaultValue,
	}
}

func (b BoolDocumentable) GetType() string {
	return "bool"
}

func (b BoolDocumentable) GetSubType() string {
	return ""
}

func (b BoolDocumentable) GetDocumentation() string {
	return b.Documentation
}

func (b BoolDocumentable) IsRequired() bool {
	return b.Required
}

func (b BoolDocumentable) GetValidationRequirements() string {
	return b.Validation
}

func (b BoolDocumentable) GetChildrenDocumentation() []Documentable {
	return []Documentable{}
}

func (b BoolDocumentable) GetDefaultValue() string {
	return fmt.Sprintf("%t", b.DefaultValue)
}

type FloatDocumentable struct {
	Documentation string
	Required      bool
	Validation    string
	DefaultValue  float64
}

// MakeFloatDocumentable
func MakeFloatDocumentable(doc string, required bool, validation string, defaultValue float64) FloatDocumentable {
	return FloatDocumentable{
		Documentation: doc,
		Required:      required,
		Validation:    validation,
		DefaultValue:  defaultValue,
	}
}

func (f FloatDocumentable) GetType() string {
	return "float"
}

func (f FloatDocumentable) GetSubType() string {
	return ""
}

func (f FloatDocumentable) GetDocumentation() string {
	return f.Documentation
}

func (f FloatDocumentable) IsRequired() bool {
	return f.Required
}

func (f FloatDocumentable) GetValidationRequirements() string {
	return f.Validation
}

func (f FloatDocumentable) GetChildrenDocumentation() []Documentable {
	return []Documentable{}
}

func (f FloatDocumentable) GetDefaultValue() string {
	return fmt.Sprintf("%f", f.DefaultValue)
}

type ArrayDocumentable struct {
	Documentation string
	Required      bool
	Validation    string
	SubDocument   Documentable
}

// MakeArrayDocumentable
func MakeArrayDocumentable(doc string, required bool, validation string, subDoc Documentable) ArrayDocumentable {
	return ArrayDocumentable{
		Documentation: doc,
		Required:      required,
		Validation:    validation,
		SubDocument:   subDoc,
	}
}

func (a ArrayDocumentable) GetType() string {
	return "array"
}

func (a ArrayDocumentable) GetSubType() string {
	return a.SubDocument.GetType()
}

func (a ArrayDocumentable) GetDocumentation() string {
	return a.Documentation
}

func (a ArrayDocumentable) IsRequired() bool {
	return a.Required
}

func (a ArrayDocumentable) GetValidationRequirements() string {
	return a.Validation
}

func (a ArrayDocumentable) GetChildrenDocumentation() []Documentable {
	return []Documentable{a.SubDocument}
}

func (a ArrayDocumentable) GetDefaultValue() string {
	return ""
}

type ObjectDocumentable struct {
	Type          string
	Documentation string
	Required      bool
	Validation    string
	Children      []Documentable
}

// MakeObjectDocumentable
func MakeObjectDocumentable(tpe string, doc string, required bool, validation string, children []Documentable) ObjectDocumentable {
	return ObjectDocumentable{
		Type:          tpe,
		Documentation: doc,
		Required:      required,
		Validation:    validation,
		Children:      children,
	}
}

func (o ObjectDocumentable) GetType() string {
	return o.Type
}

func (o ObjectDocumentable) GetSubType() string {
	return ""
}

func (o ObjectDocumentable) GetDocumentation() string {
	return o.Documentation
}

func (o ObjectDocumentable) IsRequired() bool {
	return o.Required
}

func (o ObjectDocumentable) GetValidationRequirements() string {
	return o.Validation
}

func (o ObjectDocumentable) GetChildrenDocumentation() []Documentable {
	return o.Children
}

func (o ObjectDocumentable) GetDefaultValue() string {
	return ""
}

type TimeDocumentable struct {
	Documentation string
	Required      bool
	Validation    string
	DefaultValue  time.Time
}

// MakeTimeDocumentable
func MakeTimeDocumentable(doc string, required bool, validation string, defaultValue time.Time) TimeDocumentable {
	return TimeDocumentable{
		Documentation: doc,
		Required:      required,
		Validation:    validation,
		DefaultValue:  defaultValue,
	}
}

func (t TimeDocumentable) GetType() string {
	return "time"
}

func (t TimeDocumentable) GetSubType() string {
	return ""
}

func (t TimeDocumentable) GetDocumentation() string {
	return t.Documentation
}

func (t TimeDocumentable) IsRequired() bool {
	return t.Required
}

func (t TimeDocumentable) GetValidationRequirements() string {
	return t.Validation
}

func (t TimeDocumentable) GetChildrenDocumentation() []Documentable {
	return []Documentable{}
}

func (t TimeDocumentable) GetDefaultValue() string {
	return t.DefaultValue.Format(time.RFC3339)
}
