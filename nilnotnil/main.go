package main

import (
	"fmt"
	"reflect"
)

type Payload interface {
	GetType() string
}

type EncryptedPayload struct{}

func (payload *EncryptedPayload) GetType() string {
	return fmt.Sprintf("%T", payload)
}

type Object interface {
	GetPayload() Payload
}

type ObjectA struct {
	payload Payload
}

func (obj *ObjectA) GetPayload() Payload {
	return obj.payload
}

type ObjectB struct {
	payload *EncryptedPayload
}

func (obj *ObjectB) GetPayload() Payload {
	return obj.payload
}

type ObjectC struct {
	payload *EncryptedPayload
}

func (obj *ObjectC) GetPayload() Payload {
	if obj.payload == nil {
		return nil
	}
	return obj.payload
}

func main() {
	var object Object
	var value reflect.Value
	object = &ObjectA{}
	value = reflect.ValueOf(object)
	fmt.Println("ObjectA")
	fmt.Printf("[%s], [%s]\n", value.Elem(), value.Type())
	payload := object.GetPayload()
	payloadValue := reflect.ValueOf(payload)
	fmt.Println(payloadValue)
	//fmt.Printf("[%s], [%s]\n", payloadValue.Elem(), payloadValue.Type())

	if object.GetPayload() == nil {
		fmt.Println("nil payload")
	} else {
		fmt.Println("non nil payload")
	}

	object = &ObjectB{}
	value = reflect.ValueOf(object)
	fmt.Println("ObjectB")
	fmt.Printf("[%s], [%s]\n", value, value.Type())
	payload = object.GetPayload()
	payloadValue = reflect.ValueOf(payload)
	fmt.Println(payloadValue)
	fmt.Printf("[%s], [%s]\n", payloadValue.Elem(), payloadValue.Type())
	if payload == nil {
		fmt.Println("nil payload")
	} else {
		fmt.Println("non nil payload")
	}

	object = &ObjectC{}
	value = reflect.ValueOf(object)
	fmt.Println("ObjectC")
	fmt.Printf("[%s], [%s]\n", value.Elem(), value.Type())
	if object.GetPayload() == nil {
		fmt.Println("nil payload")
	} else {
		fmt.Println("non nil payload")
	}
}
