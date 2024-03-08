# LPEdit

LPEdit is a reimplementation of Line6 HD Edit made through usb protocol reverse engineering.

## Supported features (legacy version)

- Reading the preset info when a preset stomp switch is activated on the POD:
  - Pedals parameters value;
  - Pedal Board Items position.
- Reading/Writing FX pedal type;
- Reading/Writing FX pedal parameters;
- Reading tempo info when pressing the TAP switch.

## Features in the TODO LIST

- Implementing the Amp UI
- Implementing position change for the elements
- Implementing setup functions (midi, tempo, etc.)
- Implementing more query messages for the POD.

## Supported hardware

- POD HD 500X

## Hardware that we want to support

- POD HD PRO (need kernel support first)
- POD HD 500
- POD HD 400
- POD HD 300

## Manual building

Requirements :

- NodeJS v18 minimum
- Golang 1.22
- Wails `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

Running in dev mode `wails dev`

Building the binary `wails build`

## Known bugs

- The software may crash when starting to listen to the POD events;
- The software may crash when stopping a POD and restarting it;

## Screenshots (legacy version)

![Alt text](/screenshots/Preset.png?raw=true "Preset management")
