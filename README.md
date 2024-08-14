# go-qrcode-generator

### Things todo list

1. Clone this repository: `git clone https://github.com/hendisantika/go-qrcode-generator.git`
2. Navigate to the folder: `cd go-qrcode-generator`
3. Run the application: `go run main.go`

### Generate the QRCode

```shell
curl -X POST \
    --form "size=256" \
    --form "content=https://s.id/hendisantika" \
    --output data/qrcode.png \
    http://localhost:8080/generate
```

### See the error messages

If you'd like to see the error messages, run one or both of the following commands.

```shell
curl -X POST --form "content=https://s.id/hendisantika" http://localhost:8080/generate

"Could not determine the desired QR code size."


curl -X POST --form "size=256" http://localhost:8080/generate

"Could not determine the desired QR code content."

```