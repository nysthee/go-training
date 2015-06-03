#Step 6: Combine multiple different kind of events in a single loop

- Create a program that fires 100 concurrent HTTP-GET-requests to the server. Pass a 'delay'-parameter within the url.
- The server is provided. Upon receipt of a call, the server will sleep for 'delay' before returning a response. See below on how to start the server.
- Report the duration from each goroutines back to the "main"-loop using a channel.
- Stop waiting for responses after 1000 ms and terminate the main-loop
- Report average waiting time and percentage of total request that completed within the deadline.


``` sh
    # build the server
    go build delayed_server.go

    # start the server 
    ./delayed_server

    # verify the server: delay with approx: 1000 ms (with random noise)
    time curl 'http://localhost:3000?delay=1000'

```

**Examples can be found on:**
- goroutines and channels: http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide#31 etc
- select-loop: http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide#33
