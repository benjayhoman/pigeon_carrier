# Pigeon Carrier

This is a sample app that I have created in order to learn the bubble tea library. It's function is to act as a terminal based, striped down version of postman. I prefer terminal based applications so here we are.

## Development

Requires Go 1.26 or later. The UI uses Bubble Tea v2, Bubbles v2, and Lip Gloss v2 from `charm.land`.

- Build: `make build`
- Run: `make run`
- Test: `go test ./...`

## Interface
```
>> [GET] https://sample.app.com/path/to/something?query=nothing
>> Headers: [+]
    Header1 : value1 [-]
    Header2 : value2 [-]

Response: 200
1    {
2       "field": 42,
3       "other_field": "something"
4    }

([Ctrl+Enter] to send, [Ctrl+C] to quit)
```

## Todo
- Cleanup Program's update function
- Make AddRemoveBox more generic
    - rename it to Button
    - pass in its functionality instead of it being hard coded to add remove headers