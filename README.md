# trs - Time Recording System

trs is a command-line based time recording system written in Go, with an easy to use interface that works 100% offline.

## Installation

With Go installed (and your `$GOPATH/bin` added to `$PATH`), you can install trs via

```shell
go get github.com/shimst3r/trs
```

To test the installation, run `trs help`.

## How to use trs

Before you can use trs, you have to set it up via

```shell
trs init
```

This will create a [SQLite](https://sqlite.org/index.html) database at `$HOME/.trs.db` where all app-related data is stored.

After the setup, trs can do three things:

- `trs start` to start a new time entry.
- `trs stop` to stop the last time entry.
- `trs today` to print the time recorded on the same day.

## Disclaimer

Go has a caring and friendly community where everyone is welcome. To reflect this I encourage you to read and follow the [Conde of Conduct](https://www.gophercon.com/page/1475132/code-of-conduct) before interacting with this project. Thank you! 🤗 

## License

[MIT](./LICENSE)