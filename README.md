# UserExtraService

A velocity which will be handling user extra data.

## Functions

### Add

Takes in user extra data and adds it into database.

```go
response, err := client.Add(ctx, &userData)
```

### Get

Takes in user id and returns the extra data.

```go 
response, err := client.Get(ctx, &proto.GetRequest{UserID: "UserID"})
```

### Update

Takes in user id and update. then causes that update.

```go
response, err := client.Update(ctx, &proto.UpdateRequest{
    UserID: "UserID",
    Update: Update,
})
```
### Validate

Takes in user data and returns if it is valid.

```go
response, _ := client.Validate(ctx, UserData)
```
