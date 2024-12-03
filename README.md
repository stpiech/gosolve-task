# Gosolve recruitment task

This is simple API service which allows you to check in which index is the number you are looking for (or on which index there is a close number if exact was not found).

## Run

To run this service locally just run `make run`

## Test 

To run tests run `make test`

## Config

To configure your app (`port` and `logLevel`) use `config.json` file

## Packages

### Api

Registers and handles endpoints

## Loader 

Loads data from files (both numbers file and config)

## Logger 

Set the loggers for different logging levels. Has a setter method for global log level variable

## Search

Performs binary search to find a number. If number is not found, checks if the closest one is in 10% range. If so, it returns it.
