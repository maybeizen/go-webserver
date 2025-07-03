# Go Webserver/API

little thing im tinkering with to get better at Go

# Building

(yeah i know im using node scripts for building. shut up)

- `npm run build` - builds to ./bin/webserver.exe (Windows)
- `npm run build-linux` builds to ./bin/webserver

# Endpoints (see [routes.go](./internal/routes/routes.go))

- `/api/health` (GET) - Returns `200 OK` if the server is running.
- `/api/random` (GET) - Returns a random number 1-100

---

`/api/say` (POST) - Repeats back what you said.

Request with a JSON body like this:

```json
{
  "message": "Hello, world!"
}
```
