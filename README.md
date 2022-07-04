# The Coolest Shuffler
_API to handle the deck and cards to be used in any game like Poker or Blackjack_

![The standard 52-card deck of French playing cards illustration](assets/the-coolest-shuffler.png)

## Build

This project has only one dependency to run: **docker-compose**. For development purposes you may need to: 

- **Go 1.18** 
- **make**

To start you must run:
```sh
docker-compose up -d
```
To stop the app you must run:
```sh
docker-compose stop
```

## Testing

If you want to run the tests you can call:
```sh
make test
``` 
or
```sh
make test-all
``` 
for testing everything (integration tests included).

## API
Create a new **Deck**:

```sh
curl -i -X GET "http://localhost:8916/the-coolest-shuffler/v1/deck/new?shuffle=true&amount=2&suits=CLUBS"
```

Open a **Deck**:

```sh
curl -i -X GET "http://localhost:8916/the-coolest-shuffler/v1/deck/:id"
```

Draw a **Card**:

```sh
curl -i -X GET "http://localhost:8916/the-coolest-shuffler/v1/deck/:id/draw?count=1"
```

## Troubleshooting

Some [Makefile](Makefile) goals need to run global dependencies, if you have problems running _mockery_, please, try:

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

> HINT: your `go` installation path can be different.

Problems like:
```log
ERRO[0000] failed to connect to `host=the-coolest-shuffler-postgres user=postgres database=postgres`: hostname resolving error (lookup the-coolest-shuffler-postgres: Temporary failure in name resolution)
```

Try this `configs/application.yaml`:

```yaml
version: "1.0.0"
postgres:
  host: "localhost"
  port: "5432"
  database: "postgres"
  user: "postgres"
  password: "postgres"
  timezone: "America/Fortaleza"
  sslmode: "disable"
redis:
  host: "localhost"
  port: "6379"
  database: 0
  password: ""
app:
  host: "localhost"
  port: "8916"
```
