# Step 2: Create a struct

Create a struct with the following attributes: your name, age and interests.
- Print the values to stdout
- Create a unit-test to proove that the values a correct. Use the "testing"-package.

**Examples can be found on:**
 - defining a struct: https://github.com/MarcGrol/goopenkitchen/blob/master/struct/struct.go
 - unit-testing: https://github.com/MarcGrol/goopenkitchen/blob/master/testit/reverse_test.go

**Tips:**
- To compile and build your application, use:
``` sh
    go run person.go     # no intermediate binary created
    go build             # binary 'step_2' ends up in current directory
    go install           # binary 'step_2' ends up in '${GOPATH}/bin'-directory
    ${GOPATH}/bin/step_2 # to start your application
```
- To run unit-tests, use:
``` sh
    go test
````
     

    
