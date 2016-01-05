# gb-cobra
Plugin to http://getgb.io to allow easy running of the spf13/cobra generator wihin a gb project. 
The cobra generator is used to working with a $GOPATH.

## Installation

    go get github.com/kkeuning/gb-cobra/...

## Pre-reqs

	go get github.com/constabulary/gb/...
	go get -v github.com/spf13/cobra/cobra

$GOPATH/bin should be in your $PATH, gb-cobra assumes the cobra executable can be found in $PATH.  

This is based on Doug Clark's gb-run plug-in, the primary difference being that gb-run adds GB_PROJECT_DIR/vendor to the GOPATH, which is not allowable for the cobra tool.  

Or from the shell, if PROJ_DIR is set to the value of GB_PROJECT_DIR from `gb env` (assume /Users/kkeuning/gb-projects/test-app), then:

`gb-run cobra init $PROJ_DIR/src/newAppName`

Would be the equivalent of executing:

`env GOPATH=$PROJ_DIR:$PROJ_DIR/vendor cobra init $PROJ_DIR/src/newAppName`

which results in:

`Error: Cobra only supports project within $GOPATH`

With a few minor changes to create a gb-cobra plug-in:

`gb cobra init $PROJ_DIR/src/newAppName`

Is the equivalent of executing:

`env GOPATH=$PROJ_DIR cobra init $PROJ_DIR/src/newAppName`

results in:

```
Your Cobra application is ready at
/Users/kkeuning/gb-projects/test-app/src/newAppName
Give it a try by going there and running `go run main.go`
Add commands to it by running `cobra add [cmdname]`
```

Of course, within a gb project you ill be running `gb build`, not `go run main.go` and `gb cobra add [cmdname]` rather than just `cobra add [cmdname]` as instructed.

It could be overkill to create a plug-in for this, but its a convenience.
