## Filesystem Demo

# Run instruction:
bash
$ cd output
$ chmod +x *
$ ./fsdemo_mac_arm   //mac m1 or m2 
$ ./fsdemo_mac_intel // mac intel
$ ./fsdemo_linux     // linux 64bit
$ fsdemo_win.exe     // windows

--------------------------------------------------------------------------
# project structure

entrypoint : app/main.go#

interface : 
internal/filesystem/filesystem.go

concrete implementation of above interface : 
internal/inmemoryfs/inmemoryfs.go

unit tests:
inmemoryfs_test.go

--------------------------------------------------------------------------
# Golang app Build instructions
1. Install go 1.24
2. Check out code and open the root folder - fsdemo
3. mac intel
   a. env GOOS=darwin GOARCH=amd64 go build -o output/fsdemo_mac_intel app/main.go 
4. mac m1/m2
   a. env GOOS=darwin GOARCH=arm64 go build -o output/fsdemo_mac_arm app/main.go 
5. linux
   a. env GOOS=linux GOARCH=amd64 go build -o output/fsdemo_linux app/main.go 
6. windows
   a. env GOOS=windows GOARCH=amd64 go build -o output/fsdemo_win.exe app/main.go 

--------------------------------------------------------------------------
# Run test files

go test -v ./...
