//nolint
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Container struct {
	Msg     string  `json:"msg"`
	Payload Payload `json:"payload"`
}

//go:generate gobin -m -run github.com/dmarkham/enumer -type=PayloadType -json -trimprefix=Type
type PayloadType int

// It's good practice to use 0 as the undefined type. This prevents issues when unmarshalling invalid data where
// the integer value would be defaulted to 0 when the data is empty.
const (
	TypeUndefined PayloadType = iota
	TypeEncryptedPayload
	TypePublicPayload
)

type Payload interface {
	GetType() PayloadType
}

type PayloadHeader struct {
	Type PayloadType `json:"type"`
}

type PublicPayload struct {
	PayloadHeader
	Data string `json:"data"`
}

func (p *PublicPayload) GetType() PayloadType {
	return TypePublicPayload
}

type EncryptedPayload struct {
	PayloadHeader
	EncryptedPayloadIV  string `json:"iv"`
	EncryptedPayload    string `json:"cipherText"`
	EncryptedPayloadTag string `json:"tag"`
}

func (e *EncryptedPayload) GetType() PayloadType {
	return TypeEncryptedPayload
}

func (c *Container) UnmarshalJSON(data []byte) (err error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var payloadHeader PayloadHeader
	err = json.Unmarshal(raw["payload"], &payloadHeader)
	if err != nil {
		return
	}

	type alias Container
	switch payloadType := payloadHeader.Type; payloadType {
	case TypeUndefined:
		return errors.New("unable to unmarshall payload with undefined type")
	case TypeEncryptedPayload:
		c.Payload = &PublicPayload{}
	case TypePublicPayload:
		c.Payload = &EncryptedPayload{}
	default:
		return errors.Errorf("unknown payload type: %s", payloadHeader.Type)
	}
	return json.Unmarshal(data, (*alias)(c))
}

func main() {

	pubPayload := PublicPayload{
		PayloadHeader: PayloadHeader{
			Type: TypePublicPayload,
		},
		Data: "hello world!",
	}

	encPayload := EncryptedPayload{
		PayloadHeader: PayloadHeader{
			Type: TypeEncryptedPayload,
		},
		EncryptedPayloadIV:  "123",
		EncryptedPayload:    "hello encrypted world!",
		EncryptedPayloadTag: "v1",
	}

	pubContainer := Container{
		Msg:     "I'm public",
		Payload: &pubPayload,
	}
	encContainer := Container{
		Msg:     "I'm private",
		Payload: &encPayload,
	}

	pubJson, _ := json.Marshal(pubContainer)
	encJson, _ := json.Marshal(encContainer)
	untypedJson := []byte(`{"msg":"I'm untyped","payload":{"iv":"456","cipherText":"hello untyped encrypted payload!","tag":"v1"}}`)

	var result1 Container
	var result2 Container
	var result3 Container

	_ = json.Unmarshal(pubJson, &result1)
	_ = json.Unmarshal(encJson, &result2)
	_ = json.Unmarshal(untypedJson, &result3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
