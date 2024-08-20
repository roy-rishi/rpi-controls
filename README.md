# rpi-controls

## configuration
In `/server/.env`: Set `TOKEN` equal to a random string in `server/.env`. If `TOKEN` is set to `"none"`, the server will not check requests for authentication. Set `PORT` equal to the network port the server should listen to.

## usage
If authentication is disabled, enter the Raspberry Pi's address in a browser search bar, including the port and the /shutdown path. Eg: http://rpi.local:3009/shutdown . To use authentication, set a cookie or use another client to send the bearer token along with the GET request.
