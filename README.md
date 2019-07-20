# Twitter Proxy (GoLang)

### Brief Description

This project was created to help me better understand `Go` language.
It is a simple server proxying the twitter API with 2 endpoints;

1. Get tweets by username

   ```sh
   GET http://localhost:3000/tweets/{username}
   ```

2. Get top 10 tweets by username

   ```sh
   GET http://localhost:3000/tweets/{username}/top-monthly
   ```

## Running the App

```sh
go install
```

```sh
go run main.go
```

### Time to Complete

In total ~1.5 Days

### Pacakges

| Module                                 | Why?                                                                                                    |
| :------------------------------------- | :------------------------------------------------------------------------------------------------------ |
| github.com/dghubble/go-twitter/twitter | The twitter package exposes a `Client` that helps simplify working with the twitter API                 |
| github.com/dghubble/oauth1             | The oauth1 is a package that exposes a `http.Client` that helps simplify working with oauth             |
| github.com/gorilla/mux                 | The mux package is a request router that helps simplify how we handle routing in our server application |

### Hardest Area

I had difficulties working with the data types, working out how to transform the data into a json format and wrapping my head around responding with raw bytes

### Proudest Area

Working out how to handle network requests within `Go` working with query params and routing.
