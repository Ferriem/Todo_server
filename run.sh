#!/bin/bash

# Run the Redis server on port 7890
redis-server --port 7890 &

sleep 2

#Run the Application
go run main.go