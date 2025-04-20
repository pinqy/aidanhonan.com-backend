# AidanHonan.com Backend Server

## Local Testing

Quick way to run the backend locally:

```
go run .
```

To mimic the actual backend server:

```
go build -tags netgo -ldflags '-s -w' -o app
```
```
./app
```

To hit the endpoints:
```
curl -d {<arguments>} localhost:8080/<api_name>
```