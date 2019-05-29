# Lab 11 - Puppy Notifications

- Create directory `11_notify/USERNAME` containing a copy of upstream master `10_rest/USERNAME`
- Create `cmd/lostpuppy-service/main.go` running single endpoint:
```
POST   /api/lostpuppy/          Payload: { id: PUPPY_ID }
This stubbed endpoint returns with 2 second delay:
HTTP status 201 for even IDs
HTTP status 500 for odd IDs
```
Update Puppy Delete method to notify lostpuppy-service in a goroutine and log response code asynchronously.
