# go-training

This introduction training helps you to discover the go-language. To make it real-life applicable we use HTTP-client and server in the exercises.

The following presentation (http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide) can be used as training compagnon.

All examples from the presentation are also available at: https://github.com/MarcGrol/goopenkitchen


--

**Step 1:  [Create your environment](https://github.com/MarcGrol/go-training/tree/master/step_1)**

Instructions to get your invironment up and running

--

**Step 2: [Create a struct](https://github.com/MarcGrol/go-training/tree/master/step_2)**

Create and unit-test a 'struct'. The go equivalent of a 'class'.

--

**Step 3: [Create a http-client](https://github.com/MarcGrol/go-training/tree/master/step_3)**

Create and test a http-client

--

**Step 4: [Create a HTTP-server](https://github.com/MarcGrol/go-training/tree/master/step_4)**

Create and test a http-server

--

**Step 5: Make client fire concurrent request to server**

Fire 100 requests concurrently to the server. 

TIP: Use go-routines
Example: go doit()

--

**Step 6: Make client fire concurrent request to server with deadline**

Fire 100 requests concurrently to the server and quit after 500 msec.
TIP: Use go-routines, channels and select-loop
Examples can be found on: 
 - go-routines: TODO
 - channels: https://github.com/MarcGrol/goopenkitchen/blob/master/channels/channels.go
 - select: https://github.com/MarcGrol/goopenkitchen/blob/master/select/select.go

 









