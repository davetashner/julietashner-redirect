# julietashner-redirect
This is a small Go application that uses the `net/http` package to perform a
301 redirect from `julietashner.com` to `julietashner.pixieset.com`.

# About
This application sets up a handler function, `redirectHandler`, that
performs the redirect using `http.Redirect` with the 
`http.StatusMovedPermanently` status code (which represents a 301 
redirect). The main function then sets up the HTTP server to listen 
on port 8080 and serve requests using the `http.HandleFunc` method. 
When a request is received, the `redirectHandler` is called to handle 
the request and perform the redirect.

# Running
To run the application, make sure you have Go installed on your system, 
and execute the following command in the terminal:

```go
go run main.go
```

The server will start listening on port 8080. When you access
`http://julietashner.com` in a web browser, it will redirect
you to `https://julietashner.pixieset.com`.