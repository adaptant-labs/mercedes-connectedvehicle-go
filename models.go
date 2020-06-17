package mercedes

import (
	"encoding/json"
	"errors"
)

type Vehicle struct {
	Id           string `json:"id"`
	LicensePlate string `json:"licenseplate"`
	VIN          string `json:"finorvin"`
}

type VehicleDetail struct {
	Id               string `json:"id"`
	LicensePlate     string `json:"licenseplate"`
	SalesDesignation string `json:"salesdesignation"`
	VIN              string `json:"finorvin"`
	ModelYear        string `json:"modelyear"`
	ColorName        string `json:"colorname"`
	FuelType         string `json:"fueltype"`
	PowerHP          string `json:"powerhp"`
	PowerKW          string `json:"powerkw"`
	NumberOfDoors    string `json:"numberofdoors"`
	NumberOfSeats    string `json:"numberofseats"`
}

type RetrievalStatus string

const (
	Valid        RetrievalStatus = "VALID"
	Initialized                  = "INITIALIZED"
	Invalid                      = "INVALID"
	NotSupported                 = "NOT_SUPPORTED"
)

func (rs *RetrievalStatus) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type RS RetrievalStatus
	var r = (*RS)(rs)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	switch *rs {
	case Valid, Initialized, Invalid, NotSupported:
		return nil
	}

	return errors.New("invalid retrieval status")
}

type PressureUnit string

const (
	KiloPascal PressureUnit = "KILOPASCAL"
)

func (pu *PressureUnit) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type PU PressureUnit
	var r = (*PU)(pu)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	switch *pu {
	case KiloPascal:
		return nil
	}

	return errors.New("invalid pressure unit")
}

type TirePressureStatus struct {
	Unit            PressureUnit    `json:"unit"`
	Value           float64         `json:"value"`
	RetrievalStatus RetrievalStatus `json:"retrievalstatus"`
	Timestamp       int64           `json:"timestamp"`
}

type Tires struct {
	TirePressureFrontLeft  TirePressureStatus `json:"tirepressurefrontleft"`
	TirePressureFrontRight TirePressureStatus `json:"tirepressurefrontright"`
	TirePressureRearLeft   TirePressureStatus `json:"tirepressurerearleft"`
	TirePressureRearRight  TirePressureStatus `json:"tirepressurerearright"`
}

type DoorState string

const (
	Open   DoorState = "OPEN"
	Closed           = "CLOSED"
)

func (ds *DoorState) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type DS DoorState
	var r = (*DS)(ds)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	switch *ds {
	case Open, Closed:
		return nil
	}

	return errors.New("invalid door state")
}

type DoorOpenStatus struct {
	Value           DoorState       `json:"value"`
	RetrievalStatus RetrievalStatus `json:"retrievalstatus"`
	Timestamp       int64           `json:"timestamp"`
}

type DoorLockState string

const (
	Locked   DoorLockState = "LOCKED"
	Unlocked               = "UNLOCKED"
)

func (ds *DoorLockState) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type DS DoorLockState
	var r = (*DS)(ds)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	switch *ds {
	case Locked, Unlocked:
		return nil
	}

	return errors.New("invalid door lock state")
}

type DoorLockStatus struct {
	Value           DoorLockState   `json:"value"`
	RetrievalStatus RetrievalStatus `json:"retrievalstatus"`
	Timestamp       int64           `json:"timestamp"`
}

type Doors struct {
	DoorStatusFrontLeft      DoorOpenStatus `json:"doorstatusfrontleft"`
	DoorStatusFrontRight     DoorOpenStatus `json:"doorstatusfrontright"`
	DoorStatusRearLeft       DoorOpenStatus `json:"doorstatusrearleft"`
	DoorStatusRearRight      DoorOpenStatus `json:"doorstatusrearright"`
	DoorLockStatusFrontLeft  DoorLockStatus `json:"doorlockstatusfrontleft"`
	DoorLockStatusFrontRight DoorLockStatus `json:"doorlockstatusfrontright"`
	DoorLockStatusRearLeft   DoorLockStatus `json:"doorlockstatusrearleft"`
	DoorLockStatusRearRight  DoorLockStatus `json:"doorlockstatusrearright"`
	DoorLockStatusDeckLid    DoorLockStatus `json:"doorlockstatusdecklid"`
	DoorLockStatusGas        DoorLockStatus `json:"doorlockstatusgas"`
	DoorLockStatusVehicle    DoorLockStatus `json:"doorlockstatusvehicle"`
}

type DoorLockCommand string

const (
	Lock   DoorLockCommand = "LOCK"
	Unlock                 = "UNLOCK"
)

func (dl *DoorLockCommand) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type DL DoorLockCommand
	var r = (*DL)(dl)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	switch *dl {
	case Lock, Unlock:
		return nil
	}

	return errors.New("invalid door lock/unlock command")
}

type DoorLockChangeRequestBody struct {
	Command DoorLockCommand `json:"command"`
}

type LocationCoordinate struct {
	Value           float64         `json:"value"`
	RetrievalStatus RetrievalStatus `json:"retrievalstatus"`
	Timestamp       int64           `json:"timestamp"`
}

type Location struct {
	Latitude  LocationCoordinate `json:"latitude"`
	Longitude LocationCoordinate `json:"longitude"`
	Heading   LocationCoordinate `json:"heading"`
}

type DistanceUnit string

const (
	Kilometers DistanceUnit = "KILOMETERS"
)

func (du *DistanceUnit) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type DU DistanceUnit
	var r = (*DU)(du)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	switch *du {
	case Kilometers:
		return nil
	}

	return errors.New("invalid distance unit")
}

type DistanceDriven struct {
	Unit            DistanceUnit    `json:"unit"`
	Value           int             `json:"value"`
	RetrievalStatus RetrievalStatus `json:"retrievalstatus"`
	Timestamp       int64           `json:"timestamp"`
}

type DistanceDrivenResponse struct {
	Odometer           DistanceDriven `json:"odometer"`
	DistanceSinceReset DistanceDriven `json:"distancesincereset"`
	DistanceSinceStart DistanceDriven `json:"distancesincestart"`
}

type PercentUnit string

const (
	Percent PercentUnit = "PERCENT"
)

func (pu *PercentUnit) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type PU PercentUnit
	var r = (*PU)(pu)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	switch *pu {
	case Percent:
		return nil
	}

	return errors.New("invalid percent unit")
}

type FuelLevel struct {
	Unit            PercentUnit     `json:"unit"`
	Value           int             `json:"value"`
	RetrievalStatus RetrievalStatus `json:"retrievalstatus"`
	Timestamp       int64           `json:"timestamp"`
}

type FuelLevelResponse struct {
	FuelLevelPercent FuelLevel `json:"fuellevelpercent"`
}

type StateOfCharge struct {
	Unit            PercentUnit     `json:"unit"`
	Value           int             `json:"value"`
	RetrievalStatus RetrievalStatus `json:"retrievalstatus"`
	Timestamp       int64           `json:"timestamp"`
}

type StateOfChargeResponse struct {
	StateOfCharge StateOfCharge `json:"stateofcharge"`
}
