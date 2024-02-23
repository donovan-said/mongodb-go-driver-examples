# Template Go Github Repository

A repository to use as a template for Go repositories.

## Requirements

| Tool                                                                                              | Description                                     |
| :------------------------------------------------------------------------------------------------ | :---------------------------------------------- |
| [Docker](https://www.docker.com/products/docker-desktop/)                                         | This is used to launch the MongoDB container    |
| [mongosh](https://www.mongodb.com/docs/mongodb-shell/)                                            | The MongoDB Shell                               |
| [MongoDB for VS Code](https://marketplace.visualstudio.com/items?itemName=mongodb.mongodb-vscode) | This is used to create a playground for MongoDB |

## Run Docker MongoDB Container

```shell
docker-compose up -d
docker ps
docker volume ls
```

## Connect with Mongosh

```shell
mongosh admin -u root -p rootpassword
```

## Shutdown the Docker Container

```shell
# Shutdown without deleting all containers
docker-compose stop
# Shutdown with deleting all containers
docker-compose down
```

## VSCode MongoDb Extension

Please see the [VSCode MongoDB extension](https://code.visualstudio.com/docs/azure/mongodb) 
documentation

### Connection String

```shell
mongodb://root:rootpassword@127.0.0.1/
```
