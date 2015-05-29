# go-training
Introduction in the go programming language.

--

*Step 1:  Create your environment*
Download and instruction on https://golang.org/doc/install

Prepare your environment:

    mkdir ~/go
    export GOPATH=~/go                                # define GOPATH
    go get github.com/MarcGrol/go-training            # to fetch and install
    cd ${GOPATH}/src/github.com/MarcGrol/go-training  # start editing
    ${GOPATH}/bin/go-training                         # to run program

*Step 2:Create a struct and unit-test it*

The struct should carry your name, age and interests.
Examples can be found on: 
 - defining a struct: https://github.com/MarcGrol/goopenkitchen/blob/master/struct/struct.go
 - unit-testing: https://github.com/MarcGrol/goopenkitchen/blob/master/testit/reverse_test.go

*Step 3: Create a http-client and write a test to proove it*

A running server can be reached at: TODO
Examples can be found on: 
 - http-client: TODO
 - testing a http-client: TODO

*Step 4: Create a HTTP-server and write a test to proove it*

Examples can be found on: 
 - http-server: https://github.com/MarcGrol/goopenkitchen/blob/master/webserver/webserver.go
 - testing a http-server: TODO

*Step 5:  Make client fire concurrent request to server*

Create 100 goroutines that talk to the server concurrently

*Step 6: Make client fire concurrent request to server with deadline*

Create 100 goroutines that talk to the server concurrently and quit after 500 msec.









