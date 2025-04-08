# skey

A Go Native cli for configuring secure keys.

## Installation

```
go install github.com/qubesome/skey/cmd/skey@latest
```

## Usage

List fido keys:
```
skey fido list
```

List PIV cards:
```
skey piv list
```

Reset PIV cards:
```
skey piv reset
```

## Security

The project aims to have limited number of dependencies, to decrease changes
of supply chain compromises.

Landlock is being used in BestEffort mode, so when running in Linux `5.13`
or above, file and network restrictions will be imposed at Kernel level.

## Hardware Support

This has been tested with NitroKey 3, Yubikey 4 and 5.

## License
This project is licensed under the Apache 2.0 License. Refer to the [LICENSE](LICENSE) file for more information.
