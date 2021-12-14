package container

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// UnmarshalJSON is a custom unmarshaler that we place on the struct that contains our polymorphic type. This mechanism
// is required because we cannot define an implementation of UnmarshalJSON on the Vehicle interface.
func (f *Fleet) UnmarshalJSON(bytes []byte) (err error) {
	// Unmarshal the top level object abstractly. This allows us to look up the set of Vehicle objects by the
	// key and use our polymorphic unmarshalling directly on each object
	var raw map[string]json.RawMessage
	err = json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}

	var rawVehicles []json.RawMessage
	err = json.Unmarshal(raw["vehicles"], &rawVehicles)
	if err != nil {
		return err
	}

	// Optimization: if you know the length of the slice required it's more performant to create the new slice
	// the same size, rather than using append()
	vehicles := make([]Vehicle, len(rawVehicles))
	for idx, rawVehicle := range rawVehicles {
		vehicle, err := unmarshalVehicle(rawVehicle)
		if err != nil {
			return err
		}

		err = json.Unmarshal(rawVehicle, &vehicle)
		if err != nil {
			return err
		}

		vehicles[idx] = vehicle
	}

	// Since the Fleet struct has more attributes than just the set of Vehicles, unmarshalling those attributes needs to
	// be accounted for in a way that would not require us to update the UnmarshalJSON function every time an attribute
	// is added. The `alias` trick allows us to use the default unmarshalling mechanism for the Fleet struct by re-typing
	// the Fleet as an `alias` and then setting the values that required the custom unmarshalling from above.
	type alias Fleet
	err = json.Unmarshal(bytes, (*alias)(f))

	f.Vehicles = vehicles

	return nil
}

// unmarshalVehicle looks for the type identifier inside the raw JSON and returns an empty struct of the correct type.
func unmarshalVehicle(bytes []byte) (Vehicle, error) {
	var err error
	var metadata VehicleMetadata
	err = json.Unmarshal(bytes, &metadata)
	if err != nil {
		return nil, err
	}

	return parseVehicleType(metadata.Type)

}

// parseVehicleType maps the typeStr to the Vehicle type and returns an empty struct of that type.
func parseVehicleType(typeStr string) (Vehicle, error) {
	switch typeStr {
	case VehicleTypeCar:
		return &Car{}, nil
	case VehicleTypeTruck:
		return &Truck{}, nil
	default:
		return nil, errors.Errorf("unknown vehicle type: %s", typeStr)
	}
}

