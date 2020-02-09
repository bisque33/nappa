# nappa

nappa is a command tool that converts various formats to [vegeta](https://github.com/tsenart/vegeta) input format.

## Usage

### curl

```
Convert curl command format to vegeta input format.

- Arguments and flags conform to curl command.
- Supported flags are listed in Flags.

Usage:
  nappa curl [flags]

Flags:
  -d, --data string             HTTP POST data
      --data-ascii string       HTTP POST ASCII data
      --data-binary string      HTTP POST binary data
      --data-raw string         HTTP POST data, '@' allowed
      --data-urlencode string   HTTP POST data url encoded
  -H, --header stringArray      headers
  -h, --help                    help for curl
  -X, --request string          method (default "GET")
```

example

```
$ nappa <paste a curl command> | vegeta attack -format=json -duration=1s -rate=1/s | vegeta encode
```
