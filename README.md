# Retrieve the Vendor of a given MAC address

`maciocall` is a Golang application that uses the `api.macaddress.io` API to retrieve the vendor of a given MAC accress.

## Usage
```sh
$ docker run --rm maciocall       
  -address string
        MAC Address to analise (Required)
  -apikey string
        API key to connect to macaddress.io (Required)
  -help
        Show this help menu.
  -output string
        Type of output. {text|json} (default "text")
```

### For building the Docker container:
```sh
$ docker build . -t maciocall
```

### For retrieving a vendor:
```sh
$ docker run --rm maciocall -address <MAC_ADDRESS> -apikey <YOUR_API_KEY>
```
Example:
```sh
$ docker run --rm maciocall -address 00:02:ba:ff:ef:57 -apikey <YOUR_API_KEY> -output json
{"MACAddress":"00:02:ba:ff:ef:57","CompanyName":"Cisco Systems, Inc"}
```