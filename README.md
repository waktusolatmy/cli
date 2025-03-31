# Waktu Solat CLI

- [Waktu Solat CLI](#waktu-solat-cli)
  - [About](#about)
  - [Install](#install)
    - [Homebrew](#homebrew)
    - [Script](#script)
    - [Build from source](#build-from-source)
  - [Usage](#usage)
    - [List all zones](#list-all-zones)
    - [Set default zone](#set-default-zone)
    - [View default zone](#view-default-zone)
    - [Get waktu solat using default zone](#get-waktu-solat-using-default-zone)
    - [Get waktu solat for a specific zone (skip default zone)](#get-waktu-solat-for-a-specific-zone-skip-default-zone)
  - [Credits](#credits)

## About

Waktu Solat CLI provides accurate prayer times for various locations in Malaysia

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

## Usage

### List all zones

**Command**

```sh
waktusolat zones [-o json|yaml]
```

**Example**

```
$ waktusolat zones
JAKIMCODE NEGERI              DAERAH
JHR01     Johor               Pulau Aur dan Pulau Pemanggil
JHR02     Johor               Johor Bahru, Kota Tinggi, Mersing, Kulai
...
WLY01     Wilayah Persekutuan Kuala Lumpur, Putrajaya
WLY02     Wilayah Persekutuan Labuan
```

**Example (JSON)**

```
$ waktusolat zones -o json   
[
    {
        "jakimCode": "JHR01",
        "negeri": "Johor",
        "daerah": "Pulau Aur dan Pulau Pemanggil"
    },
    {
        "jakimCode": "JHR02",
        "negeri": "Johor",
        "daerah": "Johor Bahru, Kota Tinggi, Mersing, Kulai"
    },
    ...
    {
        "jakimCode": "WLY01",
        "negeri": "Wilayah Persekutuan",
        "daerah": "Kuala Lumpur, Putrajaya"
    },
    {
        "jakimCode": "WLY02",
        "negeri": "Wilayah Persekutuan",
        "daerah": "Labuan"
    }
]
```

### Set default zone

**Command**

```sh
waktusolat zone set <code>
```

**Example**

```
$ waktusolat zone set TRG01
New zone set to TRG01
```

### View default zone

**Command**

```sh
waktusolat zone current [-o json|yaml]
```

**Example**

```
$ waktusolat zone current
JAKIMCODE NEGERI     DAERAH
TRG01     Terengganu Kuala Terengganu, Marang, Kuala Nerus
```

**Example (JSON)**

```
$ waktusolat zone current -o json     
{
    "jakimCode": "TRG01",
    "negeri": "Terengganu",
    "daerah": "Kuala Terengganu, Marang, Kuala Nerus"
}
```

### Get waktu solat using default zone

**Command**

```sh
waktusolat [-o json|yaml]
```

**Example**

```
$ waktusolat   
ZONE  SUBUH  SYURUK ZOHOR  ASAR   MAGHRIB ISYAK
TRG01 5:59AM 7:07AM 1:14PM 4:15PM 7:18PM  8:27PM
```

**Example (JSON)**

```
$ waktusolat -o json     
{
    "zone": "TRG01",
    "subuh": "5:59AM",
    "syuruk": "7:07AM",
    "zohor": "1:14PM",
    "asar": "4:15PM",
    "maghrib": "7:18PM",
    "isyak": "8:27PM"
}
```

### Get waktu solat for a specific zone (skip default zone)

**Command**

```sh
waktusolat -z <code> [-o json|yaml]
```

**Example**

```
$ waktusolat -z WLY01
ZONE  SUBUH  SYURUK ZOHOR  ASAR   MAGHRIB ISYAK
WLY01 6:06AM 7:12AM 1:20PM 4:22PM 7:24PM  8:33PM
```

**Example (JSON)**

```
$ waktusolat -z WLY01 -o json
{
    "zone": "WLY01",
    "subuh": "6:06AM",
    "syuruk": "7:12AM",
    "zohor": "1:20PM",
    "asar": "4:22PM",
    "maghrib": "7:24PM",
    "isyak": "8:33PM"
}
```

## Credits

- [Waktu Solat API](https://api.waktusolat.app/docs)
