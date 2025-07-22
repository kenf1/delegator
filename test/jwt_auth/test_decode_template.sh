#!/bin/bash

export token=""

run_test(){
    hurl --variable token=$token $1.hurl
    hurl --test --variable token=$token $1.hurl
}

run_test decode_dev
run_test decode_prod