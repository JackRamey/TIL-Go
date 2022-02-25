package container

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFleet_UnmarshalJSON(t *testing.T) {
	car := Car{
		VehicleMetadata: VehicleMetadata{
			Type:  VehicleTypeCar,
			Make:  "Volkswagen",
			Model: "Golf",
		},
	}
	truck := Truck{
		VehicleMetadata: VehicleMetadata{
			Type:  VehicleTypeTruck,
			Make:  "Ford",
			Model: "F-250",
		},
	}
	fleetID := uuid.New().String()
	fleet := Fleet{
		ID:       fleetID,
		Vehicles: []Vehicle{&car, &truck},
	}
	data, err := json.Marshal(fleet)
	require.NoError(t, err)

	var fromJson Fleet
	err = json.Unmarshal(data, &fromJson)
	require.NoError(t, err)

	assert.Equal(t, fleetID, fromJson.ID)
	assert.Len(t, fleet.Vehicles, 2)
}
