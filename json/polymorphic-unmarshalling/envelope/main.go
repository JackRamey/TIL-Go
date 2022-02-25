//nolint
package main

import (
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	EncryptedPayloadType = "encrypted"
	PublicPayloadType    = "public"
)

type Container struct {
	Payload PayloadEnvelope `json:"payload"`
}

type PayloadEnvelope struct {
	Payload
}

type Payload interface {
	GetType() string
}

type PayloadHeader struct {
	Type string `json:"type"`
}

type PublicPayload struct {
	PayloadHeader
	Data string `json:"data"`
}

func (p *PublicPayload) GetType() string {
	return PublicPayloadType
}

type EncryptedPayload struct {
	PayloadHeader
	EncryptedPayloadIV  string `json:"iv"`
	EncryptedPayload    string `json:"cipherText"`
	EncryptedPayloadTag string `json:"tag"`
}

func (e *EncryptedPayload) GetType() string {
	return EncryptedPayloadType
}

func (e PayloadEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Payload)
}

func (e *PayloadEnvelope) UnmarshalJSON(data []byte) (err error) {
	var payloadHeader PayloadHeader
	err = json.Unmarshal(data, &payloadHeader)
	if err != nil {
		return
	}

	switch payloadHeader.Type {
	case PublicPayloadType:
		var payload PublicPayload
		err = json.Unmarshal(data, &payload)
		if err != nil {
			return
		}
		e.Payload = &payload
	case EncryptedPayloadType:
		var payload EncryptedPayload
		err = json.Unmarshal(data, &payload)
		if err != nil {
			return
		}
		e.Payload = &payload
	default:
		return errors.Errorf("unknown payload type: %s", payloadHeader.Type)
	}
	return
}

func main() {

	pubPayload := PublicPayload{
		PayloadHeader: PayloadHeader{
			Type: "public",
		},
		Data: "hello world!",
	}

	encPayload := EncryptedPayload{
		PayloadHeader: PayloadHeader{
			Type: "encrypted",
		},
		EncryptedPayloadIV:  "123",
		EncryptedPayload:    "hello encrypted world!",
		EncryptedPayloadTag: "v1",
	}

	pubContainer := Container{Payload: PayloadEnvelope{&pubPayload}}
	encContainer := Container{Payload: PayloadEnvelope{&encPayload}}

	pubJson, _ := json.Marshal(pubContainer)
	encJson, _ := json.Marshal(encContainer)

	var result1 Container
	var result2 Container

	_ = json.Unmarshal(pubJson, &result1)
	_ = json.Unmarshal(encJson, &result2)

	print("got to end")
}
