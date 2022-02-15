#!/bin/bash
echo "enter id"
read id
echo "enter name"
read name
echo "enter role"
read role
curl http://localhost:8080/consultant \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data "{\"id\": \"$id\",\"name\": \"$name\",\"role\": \"$role\"}"