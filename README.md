# NEJS GopherJS + Vecty Talk

March 6, 2017

https://www.meetup.com/nebraskajs/events/248262244/

# Running Everything

## Slides

There is a plain HTML export in the `presentation/` directory, or you can use the [go present](https://godoc.org/golang.org/x/tools/present) tool.

    $ present -notes

http://127.0.0.1:3999/nejs.slide

## Base Examples

    $ gopherjs serve

http://localhost:8080/github.com/jmhobbs/nejs-golang-talk/examples/

## Twitter Examples

The Twitter examples use their own server which calls out to the Twitter API because they don't support JSONP, etc.
You will need an OAuth token, stored in a environment variable `TWITTER_TOKEN`.

    $ cd examples/twitter/
    $ make all
    gopherjs build -o net-http.js net-http.go
    gopherjs build -o xhr.js xhr.go
    go build -o main main.go
    go build -o net-http net-http.go
    gopherjs build -o demo.js demo.go
    $ ./main
    2018/03/06 23:45:41 Listening on http://127.0.0.1:9090/

http://127.0.0.1:9090/

## jsgo.io Example

https://jsgo.io/github.com/jmhobbs/nejs-golang-talk/examples/minimal
