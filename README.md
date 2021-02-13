# Braccio CLI

## Setup Embedded

The Arudino script found in `arm-embedded` should be flashed to whichever Arudino the Braccio hat is plugged into.

## Setup CLI

In order to build the binary you should have go installed. Once that's done, run the following command:

```bash
make build
```

The CLI will look for the environment variable `ARM_SERIAL_CHANNEL`. Set that before running the cli like so:

```bash
export ARM_SERIAL_CHANNEL=YOUR_SERIAL_CHANNEL
```

You can usually find your channel in Linux by running the command:

```bash
ls -l /dev/tty*
```

Then unplug the USB connector of the arm, and run the command again. The missing channel will be that of the arm.

With the channel found and set, run the binary called `arm-cli`. Run the binary with the command:

```shell
./arm-cli
```
