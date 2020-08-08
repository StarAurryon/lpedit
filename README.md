# LPEdit

LPEdit is a reimplementation of Line6 HD Edit made through usb protocol reverse engineering.

## Supported features

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

You will need my patched verion of QT bindings for golang available at:
https://github.com/StarAurryon/qt

Then extract this repository into your $GOPATH/src folder, go into the lpedit folder and run:
```
$(go env GOPATH)/bin/qtdeploy
```

## Known bugs

- The software may crash when starting to listen to the POD events;
- I need to kill the software when exiting or stopping the communication with the pod:
  - You need the kernel 5.8 at least to fix this issue.

## Screenshots

![Alt text](/screenshots/Preset.png?raw=true "Preset management")
