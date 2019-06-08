# Deprecated

Infoblox has created an official Go library.  This library was designed for Infoblox API version 1.4 in mind, and I've been avoiding changes that make the library incompatible with that version.  I recommend folks migrate to the official library.

https://github.com/infobloxopen/infoblox-go-client

# go-infoblox

This project implements a Go client library for the Infoblox WAPI. This library supports version 1.4.1 and user/pass
auth.

It was originally written for an early version of Go, and I'm concerned about breaking backwards compatability.

## Installing

Run

    go get github.com/fanatic/go-infoblox

Include in your source:

    import "github.com/fanatic/go-infoblox"

## Godoc

See http://godoc.org/github.com/fanatic/go-infoblox

## Using

    go run ./example/example.go

## Debugging

To see what requests are being issued by the library, set up an HTTP proxy such as Charles Proxy and then set the
following environment variable:

    export HTTP_PROXY=http://localhost:8888

## To Do

* Need to add other WAPI objects, but they should be trivial to add.
* Unit tests
* Responses as objects rather than interfaces
