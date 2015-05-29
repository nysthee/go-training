# go-training

This introduction training helps you to discover the go-language. To make it real-life applicable we use HTTP-client and server in the exercises.

The following presentation (http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide) can be used as training compagnon.  The full examples are also available at: https://github.com/MarcGrol/goopenkitchen


--
**Step 0: Read through the [presentation](http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide)**

This will give you a good overview of the language. 
Note that many examples are executable within your web-browser.

**Step 1:  Create your environment**

Download and instructions on https://golang.org/doc/install

Prepare your environment:

    mkdir ~/go
    export GOPATH=~/go                                # define GOPATH
    go get github.com/MarcGrol/go-training            # to fetch and install
    cd ${GOPATH}/src/github.com/MarcGrol/go-training  # start editing
    ${GOPATH}/bin/go-training                         # to run program

**Step 2: Create a struct**

Create a struct with the following attributes: your name, age and interests.

- Print the values to stdout
- Create a unit-test to proove that the values a correct. Use the "testing"-package.

Examples can be found on: 
 - defining a struct: https://github.com/MarcGrol/goopenkitchen/blob/master/struct/struct.go
 - unit-testing: https://github.com/MarcGrol/goopenkitchen/blob/master/testit/reverse_test.go

**Step 3: Create a http-client**

- Perform HTTP GET to http://httpbin.org/get (returns your request as a json-response)
- Test it manually: A running server can be reached at: http://httpbin.org/get
- Write a test to proove it works. Use the "net/http/httptest"-package.

Examples can be found on: 
 - http-client: https://github.com/MarcGrol/goopenkitchen/blob/master/httpClient/httpClient.go
 - json-encoding:  https://github.com/MarcGrol/goopenkitchen/blob/master/encoding/encoding.go
 - testing a http-client: https://github.com/MarcGrol/goopenkitchen/blob/master/httpClient/httpClient_test.go

**Step 4: Create a HTTP-server**

- Test it manually from your browser or using curl
- Write a test to proove it works

Examples can be found on: 
 - http-server: https://github.com/MarcGrol/goopenkitchen/blob/master/webserver/webserver.go
 - testing a http-server: TODO

**Step 5: Make client fire concurrent request to server**

Fire 100 requests concurrently to the server. 

TIP: Use go-routines
Example: go doit()


**Step 6: Make client fire concurrent request to server with deadline**

Fire 100 requests concurrently to the server and quit after 500 msec.
TIP: Use go-routines, channels and select-loop
Examples can be found on: 
 - go-routines: TODO
 - channels: https://github.com/MarcGrol/goopenkitchen/blob/master/channels/channels.go
 - select: https://github.com/MarcGrol/goopenkitchen/blob/master/select/select.go

 









