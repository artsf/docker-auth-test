AuthZ Plugin Test for Docker
============================

This example is to test plugin for Docker. It disallows all actions.

This project is Mac OS X specific and require [brew](https://brew.sh/) installed.
However it can be easily modified to be run on Linux (just remove brew).

Install:

0. Setup your `GOPATH` environment variable (usually pointing to `~/work`)
1. `mkdir -p $GOPATH/src/github.com/artsf/`
2. `cd $GOPATH/src/github.com/artsf/`
3. `git clone git@github.com:artsf/docker-auth-test.git`
4. `cd docker-auth-test`
5. `make setupmac`
6. `make`

and test that it disallows anything by:
```
docker run docker/whalesay cowsay boo
```
