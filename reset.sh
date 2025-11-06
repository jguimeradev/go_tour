#!/usr/bin/env bash

set -xe

create_main_file(){
    echo -e "package main \\n\\nfunc main(){\\n}" > main.go
    return 0
}


if [ -f main.go ]; then # [ -f ] is the same than test -f command
    rm main.go && create_main_file
    else 
    touch main.go && create_main_file
fi

echo $?