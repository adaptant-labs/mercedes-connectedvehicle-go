package mercedes

import "context"

type ConnectedVehicle struct {
	VehicleID string
	client    *ConnectedVehicleClient
}

func (v *ConnectedVehicle) GetVehicleDetail(ctx context.Context) (VehicleDetail, error) {
	var detail VehicleDetail
	ret := v.client.getJson(ctx, "/vehicles/"+v.VehicleID, &detail)
	return detail, ret
}

func (v *ConnectedVehicle) GetTirePressure(ctx context.Context) (Tires, error) {
	var tires Tires
	ret := v.client.getJson(ctx, "/vehicles/"+v.VehicleID+"/tires", &tires)
	return tires, ret
}

func (v *ConnectedVehicle) GetDoorStatus(ctx context.Context) (Doors, error) {
	var doors Doors
	ret := v.client.getJson(ctx, "/vehicles/"+v.VehicleID+"/doors", &doors)
	return doors, ret
}

func (v *ConnectedVehicle) LockDoors(ctx context.Context) error {
	cmd := DoorLockChangeRequestBody{Command: Lock}
	return v.client.postJson(ctx, "/vehicles/"+v.VehicleID+"/doors", &cmd)
}

func (v *ConnectedVehicle) UnlockDoors(ctx context.Context) error {
	cmd := DoorLockChangeRequestBody{Command: Unlock}
	return v.client.postJson(ctx, "/vehicles/"+v.VehicleID+"/doors", &cmd)
}

func (v *ConnectedVehicle) GetLocation(ctx context.Context) (Location, error) {
	var location Location
	ret := v.client.getJson(ctx, "/vehicles/"+v.VehicleID+"/location", &location)
	return location, ret
}

func (v *ConnectedVehicle) GetDistanceDriven(ctx context.Context) (DistanceDrivenResponse, error) {
	var response DistanceDrivenResponse
	ret := v.client.getJson(ctx, "/vehicles/"+v.VehicleID+"/odometer", &response)
	return response, ret
}

func (v *ConnectedVehicle) GetFuelLevel(ctx context.Context) (FuelLevelResponse, error) {
	var level FuelLevelResponse
	ret := v.client.getJson(ctx, "/vehicles/"+v.VehicleID+"/fuel", &level)
	return level, ret
}

func (v *ConnectedVehicle) GetStateOfCharge(ctx context.Context) (StateOfChargeResponse, error) {
	var state StateOfChargeResponse
	ret := v.client.getJson(ctx, "/vehicles/"+v.VehicleID+"/stateofcharge", &state)
	return state, ret
}
