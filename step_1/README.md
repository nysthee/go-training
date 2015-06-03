#Step 1:  Setup your environment


**Prepare your environment:**

Download go and follow instructions on https://golang.org/doc/install.
Find more info in [presentation](http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide#6).


**Additional steps:**
``` sh
    mkdir ~/go
    export GOPATH=~/go                                # define GOPATH
    export PATH=${GOPATH}/bin:${PATH}                 # add "go install"-path to PATH
```

**Fetch the training project**
``` sh
    go get github.com/MarcGrol/go-training/...        # to fetch and install
    cd ${GOPATH}/src/github.com/MarcGrol/go-training  # start editing
    ${GOPATH}/bin/step_2                              # to run program (there are more)
```

**Verify installation**
``` sh
    cd ${GOPATH}/src/github.com/MarcGrol/go-training/step_1
    go build hello.go                                        # build locally: ./hello
    or
    go install                                               # build to $GOPATH/bin/hello
```

**Optional: Fetch the presentation project**
``` sh
    go get github.com/MarcGrol/goopenkitchen/...        # to fetch and install
    cd ${GOPATH}/src/github.com/MarcGrol/goopenkitchen  # start editing
    ${GOPATH}/bin/select                                # to run program (there are more)
```

