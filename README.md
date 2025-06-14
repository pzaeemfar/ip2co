# oip2co

A simple command-line tool that converts IP addresses to country codes. It uses the MaxMind GeoLite2 database for accurate IP-to-country lookups.

> **Important**
> 
> Downloading the MaxMind GeoLite2 database normally requires signing up and going through a few steps. To make things easier, `oip2co` uses a pre-hosted version from [this link](https://github.com/PrxyHunter/GeoLite2/releases/latest/download/GeoLite2-Country.mmdb).
>
> However, if you prefer, you can download the official file yourself from the MaxMind website and place it at `/tmp/GeoLite2-Country.mmdb`. The tool will use that file if it's available.

## Features

- Fast IP-to-country lookups  
- Reads IPs from stdin or command-line arguments  
- Clean, simple output format  
- Automatic database download  
- Silent mode by default, use `-debug` for detailed logs  

## Installation

```bash
go install github.com/pzaeemfar/oip2co@latest
````

## Usage

```bash
# Look up Google's DNS (using stdin)
echo "8.8.8.8" | oip2co
# Output: 8.8.8.8 - US

# Look up a private IP
echo "192.168.1.1" | oip2co
# Output: 192.168.1.1 - Unknown

# Look up an IPv6 address
echo "2001:4860:4860::8888" | oip2co
# Output: 2001:4860:4860::8888 - US

# Look up IPs via command line arguments (if no stdin)
oip2co 8.8.8.8 1.1.1.1
```

## Options

* `-debug` : Enable debug output (default is silent mode)

## Notes

* The program automatically downloads the GeoLite2 database to `/tmp/GeoLite2-Country.mmdb` if missing
* Invalid IP addresses are skipped with optional debug info
* Private and unrecognized IP ranges return "Unknown" as country code
* If stdin is empty, IPs from CLI arguments are used
