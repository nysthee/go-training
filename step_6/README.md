#Step 6: Combine multiple different kind of events in a single loop

- Create a program that fires 100 concurrent HTTP-GET-requests to the server. Pass a 'delay'-parameter within the url.
- The server is provided. Upon receipt of a call, the server will sleep for 'delay' before returning a response. See below on how to start the server.
- Report the duration from each goroutines back to the "main"-loop using a channel.
- Stop waiting for responses after 1000 ms and terminate the main-loop
- Report average waiting time and percentage of total request that completed within the deadline.

**Using the client**

A half-fabricate client is provided. Add concurrency and deadlining to it.
The client can be tested and started the following way:
``` sh
    go test             # run unit tests
    go build client.go  # build it
    ./client -h         # show usage
    ./client            # run it
```


**Starting the server**
A complete server is provided. It can be found in the echopServer-directory of the project.
The prefab client can be tested and started the following way:
``` sh
    cd ../echoServer
    go test                           # run unit tests
    
    go build delayedEchoServer.go     # build the server

    # start the server in background
    ./delayedEchoServer &

    # verify the server: delay with approx: 1000 ms (with random noise)
    time curl -X 'http://localhost:3000?delay=1000&arg=hoi'
    curl -X POST 'http://localhost:3000?delay=5000' --data hoi

    # run the above mentioned client against it
    ./client

```

**Examples can be found on:**
- goroutines and channels: http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide#31 etc
- select-loop: http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide#33
