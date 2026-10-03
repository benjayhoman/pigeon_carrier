# Pigeon Carrier

This is a sample app that I have created in order to learn the bubble tea library. It's function is to act as a terminal based, striped down version of postman. I prefer terminal based applications so here we are.

## Development

Requires Go 1.26 or later. The UI uses Bubble Tea v2, Bubbles v2, and Lip Gloss v2 from `charm.land`.

- Build: `make build`
- Run: `make run`
- Test: `go test ./...`

## Interface
```
[Send] [Load] [Env]

Request: [Save] [Default Env]
    ──────────────────────────────────────────
    [GET] https://sample.app.com/path/to/something?query=nothing
    Headers: [Add Header]
    Body: [vv]
    1 
    Config: [vv]
        script: ./script1.lua
        client certificate: ./client-cert.crt
        client key: ./client-key.crt

Response:
    ──────────────────────────────────────────
    Status: 200 OK
    Headers: [>>]
    Body:
    1 {
    2   "field": 42,
    3   "other_field": "something"
    4 }

([Ctrl+Enter] to send, [Ctrl+C] to quit)
```

## Todo
- [do] revamp the UI to look better
- [do] fix focus to make it more obvious and consistant
- [spike] figure out where config will live (in a config file or in the TUI)