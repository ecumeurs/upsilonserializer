package jsonserializer

import (
	"encoding/json"

	"github.com/ecumeurs/upsilonserializer/commessage"
	"github.com/ecumeurs/upsilonserializer/serializable"
	"github.com/sirupsen/logrus"
)

// jsonserializer is a serializer using JSON as a format.
type JsonSerializer struct{}

func (js JsonSerializer) Serialize(registry serializable.SerializableRegistry, cm commessage.ComMessage) ([]byte, error) {
	toSend := commessage.NewJSONMessage(cm)
	return json.Marshal(toSend)
}

func (js JsonSerializer) Deserialize(registry serializable.SerializableRegistry, b []byte) (commessage.ComMessage, error) {
	received := commessage.JSONSendComMessage{}
	cm := commessage.ComMessage{}

	err := json.Unmarshal(b, &received)

	if err != nil {
		return cm, err
	}

	cm.Target = registry.Create(received.TargetMethod)
	if cm.Target != nil {
		err = json.Unmarshal(received.Target, cm.Target)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"component": "serializer",
				"type":      "mpack",
			}).WithError(err).Error("Error deserializing target")
			return cm, err
		}
	} else {
		if received.TargetMethod != "" {

			logrus.WithFields(logrus.Fields{
				"component": "serializer",
				"type":      "mpack",
			}).Error("Target not found")
		} else {
			logrus.WithFields(logrus.Fields{
				"component": "serializer",
				"type":      "mpack",
			}).Error("Target method is empty")
		}
	}

	cm.Content = registry.Create(received.ContentType)
	if cm.Content != nil {
		err = json.Unmarshal(received.Content, cm.Content)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"component": "serializer",
				"type":      "mpack",
			}).WithError(err).Error("Error deserializing content")
			return cm, err
		}
	} else {
		logrus.WithFields(logrus.Fields{
			"component": "serializer",
			"type":      "mpack",
		}).Error("Content not found")
	}

	cm.Meta = registry.Create(received.MetaType)
	if cm.Meta != nil {
		err = json.Unmarshal(received.Meta, cm.Meta)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"component": "serializer",
				"type":      "mpack",
			}).WithError(err).Error("Error deserializing meta")
			return cm, err
		}
	} else {
		logrus.WithFields(logrus.Fields{
			"component": "serializer",
			"type":      "mpack",
		}).Error("Meta not found")
	}

	cm.ID = received.ID
	cm.EmitedAt = received.EmitedAt
	cm.HasError = received.HasError
	cm.ErrorMessage = received.ErrorMessage
	cm.ErrorKey = received.ErrorKey

	return cm, nil
}
