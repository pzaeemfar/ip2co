# oip2co

A simple command-line tool that converts IP addresses to country codes. It uses the **IP2Location LITE IP Geolocation Database** for accurate IP-to-country lookups.

> **Important**
>
> IPv6 addresses and domain names are not supported. Any such inputs will be skipped and ignored.

## Features

* Fast IP-to-country lookups
* Reads IPs from stdin or command-line arguments
* Clean, simple output format
* Automatic database download
* Optional JSON output with `-json` flag
* Silent mode by default, use `-debug` for detailed logs

## Installation

```bash
go install github.com/pzaeemfar/oip2co@latest
```

## Usage (Examples)

```bash
cat proxies.txt | oip2co

echo "192.168.1.1" | oip2co

oip2co 8.8.8.8 http://1.1.1.1:80

oip2co -json 8.8.8.8

oip2co -debug -json 8.8.8.8
```

## Options

* `-json`  : Output results in JSON format
* `-debug` : Enable debug output (default is silent mode)

## Notes

* Invalid IP addresses are skipped with optional debug info
* Private and unrecognized IP ranges return "Unknown" as country code
* If stdin is empty, IPs from CLI arguments are used