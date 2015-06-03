# Step 2: Create a struct

Create a struct with the following attributes: your name, age and interests.
- Print the values as json to stdout
- Create a unit-test to proove that the values a correct. Use the "testing"-package.

**Examples can be found on:**
 - presentation: http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide#20
 - defining a struct: https://github.com/MarcGrol/goopenkitchen/blob/master/struct/struct.go
 - json-encoding:  https://github.com/MarcGrol/goopenkitchen/blob/master/encoding/encoding.go
 - unit-testing: https://github.com/MarcGrol/goopenkitchen/blob/master/testit/reverse_test.go

**Tips:**

- To run unit-tests, use:
``` sh
    go test
````

- To compile and build your application, use:
``` sh
    go run person.go     # no intermediate binary created
    go build             # binary 'step_2' ends up in current directory
    go install           # binary 'step_2' ends up in '${GOPATH}/bin'-directory
    ${GOPATH}/bin/step_2 # to start your application
```

     

    
