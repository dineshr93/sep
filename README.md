# sep

File separator in go
```
// in Windows
// set GOOS=linux or windows
// set GOARCH=amd64 
go build
```
```
go run .
```

```
./sep FULL_FOLDER_PATH_OF_FILES_TO_BE_SEPARATED
```
```console
/sep$ make
all                            performs clean build and test
clean                          Move back the files to original state in test folder
build                          Generate the windows and linux builds for sep
test                           Separates the test folder
git                            commits and push the changes if commit msg m is given without spaces ex m=added_files
```
