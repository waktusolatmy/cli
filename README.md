# Waktu Solat CLI

- [Waktu Solat CLI](#waktu-solat-cli)
  - [About](#about)
  - [Install](#install)
    - [Homebrew](#homebrew)
    - [Script](#script)
    - [Build from source](#build-from-source)
  - [Credits](#credits)

## About

TODO

## Install

There are multiple ways to get this CLI installed:

### Homebrew

> You must have [Homebrew](https://brew.sh) installed first

```sh
brew untap waktusolatmy/tools &>/dev/null || true
brew tap waktusolatmy/tools

brew install waktusolat
waktusolat version
```

### Script

```sh
curl -sfL https://raw.githubusercontent.com/waktusolatmy/cli/refs/heads/master/install.sh | sh
```

### Build from source

> You must have [Go](https://go.dev) installed first

```sh
git clone https://github.com/waktusolatmy/cli.git
cd cli
go build -o waktusolat
sudo mv waktusolat /usr/local/bin
waktusolat version
```

## Credits

- [Waktu Solat API](https://api.waktusolat.app/docs)
