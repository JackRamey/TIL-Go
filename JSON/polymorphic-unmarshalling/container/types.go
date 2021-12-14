package container

const (
	VehicleTypeCar   = "Car"
	VehicleTypeTruck = "Truck"
)

type Vehicle interface {
	GetType() string
}

type VehicleMetadata struct {
	Type  string `json:"type"`
	Make  string `json:"make"`
	Model string `json:"model"`
}

func (v VehicleMetadata) GetType() string {
	return v.Type
}

type Car struct {
	VehicleMetadata
}

// GetType
// Note: Must be a pointer receiver in order to properly cast to Vehicle during json.Unmarshal
func (c *Car) GetType() string {
	return VehicleTypeCar
}

type Truck struct {
	VehicleMetadata
}

// GetType
// Note: Must be a pointer receiver in order to properly cast to Vehicle during json.Unmarshal
func (t *Truck) GetType() string {
	return VehicleTypeTruck
}

type Fleet struct {
	ID       string    `json:"id"`
	Vehicles []Vehicle `json:"vehicles"`
}
