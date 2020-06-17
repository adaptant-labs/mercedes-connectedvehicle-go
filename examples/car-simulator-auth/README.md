# car-simulator-auth

This is a simple example of negotiating the Oauth2 login and consent flows described by the
[Mercedes-Benz OAuth Authentication Documentation][oauth2-docs] in order to access vehicle data, to be used in
conjunction with the [Mercedes-Benz Simulator][simulator]. Once consent has been given by the user, the door locks of the
simulated vehicle will be toggled, which can be observed directly in the simulator:

```
11:16:01 - Mercedes-Benz Simulator - Starting initialization
11:16:01 - Mercedes-Benz Simulator - Initialized capabilities
11:16:02 - Mercedes-Benz Simulator - Car with serial number 7410ED2522A7AA3256 initialized
11:19:55 - Sandbox device - CapabilityManager - Incoming telematics message for Door Locks - ACASAQAEAQABAQ==
11:19:55 - Sandbox device - Door Locks - Lock doors
11:20:01 - Sandbox device - CapabilityManager - Incoming telematics message for Door Locks - ACAA
11:20:01 - Sandbox device - Door Locks - Get Lock State
11:20:07 - Sandbox device - CapabilityManager - Incoming telematics message for Door Locks - ACASAQAEAQABAA==
11:20:07 - Sandbox device - Door Locks - Unlock doors
```

# Environment Variables

In order to carry out the OAuth flow, the client ID and secret must be provided. These are to be set in the following
environment variables:

```
MERCEDES_CLIENT_ID
MERCEDES_CLIENT_SECRET
```

These are available from the [Mercedes-Benz Developer Console][console].

[oauth2-docs]: https://developer.mercedes-benz.com/content-page/oauth-documentation
[simulator]: https://car-simulator.developer.mercedes-benz.com/
[console]: https://developer.mercedes-benz.com/console
