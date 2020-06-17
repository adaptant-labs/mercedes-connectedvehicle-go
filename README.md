# Go SDK for Mercedes-Benz Connected Vehicle API

## Overview

The [Mercedes-Benz Connected Vehicle API][connected_vehicle_api] is an experimental API for accessing connected vehicle
data and for prototyping connected vehicle services. It provides extensive information about the vehicle itself (vehicle
information, tire pressure, door status, location & heading, odometer, fuel level, as well as the state of battery
charge for electric vehicles), while also providing limited actuation controls (door lock / unlock).

[connected_vehicle_api]: https://developer.mercedes-benz.com/products/connected_vehicle/

## Installation

If not using Go modules, the SDK can be installed the same way as for other Go projects:

```
$ go get github.com/adaptant-labs/mercedes-connectedvehicle-go
```

## Getting Started

1\. Initiate a new client connection with your API key
```
// Initiate a new client connection. Set 'false' for tryout, 'true' for production API.
client := mercedes.NewClient(<API Key>, false)
```

2\. Obtain a list of vehicles
```
vehicles, err := client.GetVehicles(context.TODO())
```

3\. Choose a vehicle to perform operations on
```
vehicle := client.NewVehicle(vehicles[0].Id)
```

4\. Carry out vehicle-specific operations
```
// Vehicle APIs
detail, err := vehicle.GetVehicleDetail(context.TODO())
tires, err := vehicle.GetTirePressure(context.TODO())
err := vehicle.LockDoors(context.TODO())
err := vehicle.UnockDoors(context.TODO())
doors, err := vehicle.GetDoorStatus(context.TODO())
location, err := vehicle.GetLocation(context.TODO())
distance, err := vehicle.GetDistanceDriven(context.TODO())
level, err := vehicle.GetFuelLevel(context.TODO())
charge, err := vehicle.GetStateOfCharge(context.TODO())
```
 
## Testing with the Mercedes-Benz Car Simulator

The SDK itself can be used together with the [Mercedes-Benz Car Simulator][simulator], but must go through the
appropriate Oauth2 authentication flows in order to become accessible from the Connected Vehicle API. Note that in this
case, the client must be configured for using the production API, and initiate the connection with the exchanged access
token. A simple example of this is included in the `examples` directory. 

[simulator]: https://car-simulator.developer.mercedes-benz.com/

## Features and bugs

Please file feature requests and bugs concerning the SDK itself in the [issue tracker][tracker]. Note that as this is a
third-party SDK and we have no direct affiliation with Mercedes-Benz, we are unable to handle feature requests for the
REST API itself.

[tracker]: https://github.com/adaptant-labs/mercedes-connectedvehicle-go/issues

## License

`mercedes-connectedvehicle-go` is released under the terms of the MIT license, the full
version of which can be found in the LICENSE file included in the distribution.