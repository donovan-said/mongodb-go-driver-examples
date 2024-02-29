# MongoDB Go Driver Examples

- [MongoDB Go Driver Examples](#mongodb-go-driver-examples)
  - [Requirements](#requirements)
  - [MongoDB Container](#mongodb-container)
    - [Run Docker MongoDB Container](#run-docker-mongodb-container)
    - [Shutdown the Docker Container](#shutdown-the-docker-container)
  - [VSCode MongoDB Extension](#vscode-mongodb-extension)
    - [Running the Sample Go Scripts](#running-the-sample-go-scripts)
  - [Appendix](#appendix)
    - [Connect with Mongosh](#connect-with-mongosh)

## Requirements

| Tool                                                                                              | Description                                     |
| :------------------------------------------------------------------------------------------------ | :---------------------------------------------- |
| [Docker](https://www.docker.com/products/docker-desktop/)                                         | This is used to launch the MongoDB container    |
| [mongosh](https://www.mongodb.com/docs/mongodb-shell/)                                            | The MongoDB Shell                               |
| [MongoDB for VS Code](https://marketplace.visualstudio.com/items?itemName=mongodb.mongodb-vscode) | This is used to create a playground for MongoDB |

## MongoDB Container

For the purpose of a development environment, I'm using the `mongo:latest` image
to launch a local development environment.

### Run Docker MongoDB Container

```shell
docker-compose up -d
docker ps
docker volume ls
```

### Shutdown the Docker Container

```shell
# Shutdown without deleting all containers
docker-compose stop
# Shutdown with deleting all containers
docker-compose down
```

## VSCode MongoDB Extension

Please see the [VSCode MongoDB extension](https://code.visualstudio.com/docs/azure/mongodb)
documentation to see how to connect via the extension.

The connection string is `mongodb://root:rootpassword@127.0.0.1/`

Below is an example of what the connected database will look like via the
extension:

![MongodbExtention](./img/mongodb_extension.png)

### Running the Sample Go Scripts

To tun the scripts, run the following command:

```shell
go run cmd/cli/main.go
```

You'll be prompted asking if you want to populate the DB with sample data. This
only need to be done on the first run. You can see an example of this below:

![UserPrompt](./img/user_prompt.png)

## Appendix

### Connect with Mongosh

```shell
mongosh admin -u root -p rootpassword
```
