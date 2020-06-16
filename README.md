# Globalsign Code Test

Please make sure go environment is setup

To build the binaries

  linux
  
  -  GOOS=linux GOARCH=amd64 go build -o top
  
  windows
  -  GOOS=windows GOARCH=amd64 go build -o top
  
  mac 
  -  GOOS=darwin GOARCH=amd64 go build -o top

# Running the program

 ```
 ./top filePath
 or
 go run main.go filePath
 ```
 
 # Note

```
Since container packages were not allowed , so the problem statement has been approached with slices of struct not with  maps,heap,priority queue etc
```
