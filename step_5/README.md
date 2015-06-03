#Step 5: Concurrently fire requests to a server

- The server is provided: So below on how to start the server
- Create a program that fire 100 concurrent HTTP-GET-requests to the server
- Report the result from each goroutines back to the "main"-loop using a channel.


``` sh
    # build the server
    go build server.go

    # start the server 
    ./server

    # verify the server
    curl 'http://localhost:3000/'

```

**Examples can be found on:**
- goroutines: http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide#31 etc
