#!/bin/bash

shellExit()
{
if [ $1 -eq 1 ]; then
    printf "\nfailed!!!\n\n"
    exit 1
fi
}

printf "\nRegenerating file\n\n"
time go run -v ./cmd/mysqlmd/main.go  -addr $1 -user $2 -pass $3 -name $4 -tables $5 -prefix $6

printf "\n2222222222222222\n"
 
shellExit $?

printf "\ncreate curd code : \n"
time go build -o gormgen ./cmd/gormgen/main.go
shellExit $?

mv gormgen $GOPATH/bin
shellExit $?

go generate ./...
shellExit $?

printf "\nFormatting code\n\n"
time go run -v ./cmd/mfmt/main.go
shellExit $?

printf "\nDone.\n\n"
