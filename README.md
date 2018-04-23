# fov

A command-line utility for calculating horizontal or vertical FOV values for a
given aspect ratio. This is especially useful when tweaking the field of view
value in various 3D applications, such as games.

## Usage

### Examples

#### Determine the vertical FOV from an horizontal FOV

```
$ fov 90h 4:3
Horizontal FOV  90.00°
Vertical FOV    73.74°
Aspect ratio    4:3
```

#### Determine the horizontal FOV from a vertical FOV

```
$ fov 70v 16:9
Horizontal FOV  102.45°
Vertical FOV    70.00°
Aspect ratio    16:9
```

#### Convert an horizontal FOV to a wider aspect ratio

```
$ fov 90h 4:3 16:9
                Orig.   Converted
Horizontal FOV  90.00°  106.26°
Vertical FOV    73.74°  73.74°
Aspect ratio    4:3     16:9
```

### Reference

```
NAME:
   fov - Calculate horizontal or vertical FOV values for a given aspect ratio

USAGE:
   fov <FOV><h|v> <aspect ratio> [new aspect ratio]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## License

Copyright © 2018 Hugo Locurcio and contributors

Unless otherwise specified, files in this repository are licensed under the
MIT license, see [LICENSE.md](LICENSE.md) for more information.
