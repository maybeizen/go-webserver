# Go Webserver/API

little thing im tinkering with to get better at Go

## Building

(yeah im using Node scripts to build a Go project. Shut up)

- `npm run build` - Builds a Windows executable to `./bin/webserver.exe`
- `npm run build-linux` - Builds a Linux executable to `./bin/webserver`

## Endpoints

(check out [routes.go](./internal/routes/routes.go))

- `/api/health` (GET) - Responses with `200 OK` if the server is online and responding.
- `/api/random` (GET) - Gives you a random number between 1-100.

---

- `/api/say` (POST) - Repeats whatever you say to it.

Send this as a JSON body:

```json
{
  "message": "Hello, world!"
}
```
