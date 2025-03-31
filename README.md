# Waktu Solat CLI

- [Waktu Solat CLI](#waktu-solat-cli)
  - [Install](#install)
    - [Script](#script)
    - [Build from source](#build-from-source)
  - [Credits](#credits)

## Install

There are multiple ways to get this CLI installed:

### Script

```sh
curl -sfL https://raw.githubusercontent.com/waktusolatmy/cli/refs/heads/master/install.sh | sh
```

### Build from source

```sh
git clone https://github.com/waktusolatmy/cli.git
cd cli
go build -o waktusolat
sudo mv waktusolat /usr/local/bin
waktusolat version
```

## Credits

- [Waktu Solat API](https://api.waktusolat.app/docs)
