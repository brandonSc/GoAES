# GoAES

Straight up [FIPS-197](http://csrc.nist.gov/publications/fips/fips197/fips-197.pdf), written in Go.

### About

The motivation for this project is to understand the original publication of AES, and implement an educational example of the original algorithm. The code is designed to read as closely to the original paper as possible. Go is an excellent choice in language to accomplish this goal. Not only does it optimize very well for programs of this nature, but the syntax it self resembles the algorithm in the original FIPS-197 paper quite nicely. 

Currently, the 128-bit encrpytion algorithm is working for single 16-byte blocks of data. Decryption is not quite finished, however. Take a look at the code in the `cipher` package, and compare it to the original sudo code in [FIPS-197](http://csrc.nist.gov/publications/fips/fips197/fips-197.pdf). Some of the code in `cipher/invcipher.go` is incomplete, however `cipher/cipher.go` can be mapped nicely to the original paper, and references to sections of the original paper are listed in comments in this file.

Next steps for this project are to finish the decryption (in `invcipher.go`) and then implement CBC and CTR modes.

### Running Instructions

First download and install Go from [here](https://golang.org/dl/) (test that it works by running `go version`)

Now, there are a couple of options for building and running this project. The easiest way (but not recommended in general for running bigger Go projects) is to use `go run`. You do not need to compile the project. Just execute:
```
go run main.go
```

The other, more standard way of running go projects is to issue a `go get` from my Github repoistory. 
```
go get github.com/brandonSc/GoAES
```
Build the executible:
```
go build
```
then run with
```
./GoAES
```

Note that Go usually expects packages to reside in your `$GOPATH/src` directory, with a directory structure matching the import path structure (ie it should be in `$GOPATH/src/github.com/brandonSc/GoAES`). If you have problems running the code through either options listed above, you may have to ensure your $GOPATH is configured, then option 2 (via `go get`) should be the easiest method.

If you're familiar with Docker and don't want to deal with setting up Go, I also included a Dockerfile. You can run this way with `docker build -t goaes . && docker run goaes`.

Instructions are assuming a unix-like operating system.
