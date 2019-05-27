# Lab 10 - Puppy REST

- Create directory `10_rest/USERNAME` containing a copy of upstream master `09_json/USERNAME`
- Add file `pkg/puppy/rest.go` implementing:

```
GET    /api/puppy/{id}
POST   /api/puppy/          Payload: Puppy JSON without ID
PUT    /api/puppy/{id}      Payload: Puppy JSON without ID
DELETE /api/puppy/{id}
```

- Use [net/http/httptest](https://golang.org/pkg/net/http/httptest/) for testing
- Add flag `-p PORT` with long flag `--port PORT` to command line flags
- Add flag `-s STORE` with long flag `--store STORE` with accepted values:

```
map, sync, db
```
Document the API in README.md
