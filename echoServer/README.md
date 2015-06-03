#Echo-server needed for exercise 5 and 6

A complete server is provided. This server returns the HTTP-request back as a json document in the response.
- Method
- Url + parameters
- Headers
- Body (as string)

**Starting the server**
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

