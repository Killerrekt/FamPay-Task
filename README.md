# FamPay Development

## Installation

1. Fork this repository to your github account
2. Clone the forked repository
3. Copy the .example.env into .env and update the values according to your needs
4. Now can use the following commands for the functionality associated with it
5. For the api documentation, click [here](https://documenter.getpostman.com/view/26244894/2sAXxP9Cjf)

## Step for the Task

1. Run the server by the commands provided below
2. Hit the route /add-query, /set-mode, /remove-query to adjust the application to your needs
3. Hit the /settings route to see the current setting of the application
4. Hit the /get-data route with the query parameter of "query" and a optional query parameter of "publish"

## Extra Feature

1. Can run multiple query goroutine at once server instance
2. Nothing is hardcoded and everything is flexible

### Server Build

Use the following commands to build the server in case of any changes.
Server will be expose at port mentioned in the .env

```
make server-build
```

### Server Start

Server will be expose at port mentioned in the .env

```
make server-up
```

### Server Down

Use the following command to stop the server

```
make server-down
```

### Database Up

Use this commands to just run the postgres container standalone without the api

```
make database-up
```

### Database Down

Use this commands to stop the postgres container

```
make database-down
```
