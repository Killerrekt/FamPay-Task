# FamPay Development

## Installation

1. Fork this repository to your github account
2. Clone the forked repository
3. Copy the .example.env into .env and update the values according to your needs
4. Now can use the following commands for the functionality associated with it
5. For the api documentation, click [here](https://documenter.getpostman.com/view/26244894/2sAXxP9Cjf)

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
