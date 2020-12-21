[![License](https://img.shields.io/github/license/lil5/go-mail-cli.svg)](https://github.com/lil5/go-mail-cli/blob/master/LICENSE.md)

# GO CLI

A CLI tool to generate html emails. Using [`go-premailer`](https://github.com/vanng822/go-premailer) and the [`go template`](https://golang.org/pkg/text/template/) engine.

## Usage

1. Download the latest binary from the [releases](https://github.com/lil5/go-mail-cli/releases). If you don't find your required architecture or OS please follow the build instructions.

2. Create a directory structure like so:
   - `build`
	- `templates`
	- `pages`

3. Run the binary, this builds html files in a build directory from each `*.gohtml` file inside the `pages` directory, make sure to run the binary from the root of the directory structure.

```
$ ./gomail-<os>-<arch>
```

## Build

1. Git clone this repository.
2. `cd` into `v2` subdirectory.
3. Run `go build -o bin/gomail`
