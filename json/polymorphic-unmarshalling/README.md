# JSON

Unmarshalling polymorphic types in Go isn't quite as straightforward as other languages. Since you can't implement
methods directly on the Interface as you would be able to in a language like Java, you have to either create
a custom unmarshaller for the containing object, or create an intermediary envelope. Both are effectively using
the same concept, but the envelope attempts to contain custom unmarshalling code to the dynamic type.

The thing I dislike about custom unmarshalling for dynamic types in Go is that you have to create a custom
unmarshaller for every object that contains an abstract type, which violates the DRY principle. The packages
within this directory explore mechanisms to minimize the impact of custom unmarshalling.

[payloadmarshalling](payloadmarshalling/main.go) package uses a `PayloadEnvelope` to (un)marshal the `Payload`
interface. This works relatively well and allows the `Payload` type to keep all of its unmarshalling within the
`PayloadEnvelope`. However, this approach has a definite code smell. We add a bunch of complexity to our types just
for JSON (un)marshalling.

[commitmarshalling](commitmarshalling/main.go) package delegates the custom (un)marshalling to the container type. In this
case the `Commit` type contains a `Payload` member and must provide a custom unmarshaller to read the JSON into the
proper subclass of `Payload`. Some of the pain can be mitigated by extracting the `Payload` specific code in the
`Commit.UnmarshalJSON` function, but for this example I left the code inline.

[commitmarshalling_enumertype](commitmarshalling_enumertype/main.go) package tests what happens when an enumerated type
is missing from the JSON. This is common when a new type is added and a key is required to know what type to read. Since
the type was missing from the payload initially, it is important to know that the unmarshaller can default the type.
