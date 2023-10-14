## Welcome to Go-Base-Service
A fast codebase for building a service to serve HTTP requests. Built with love by me and the previous team. In the base-service, we have developed the following features, with ongoing features in the pipeline:
- [x]  Support for working with configurations (Dev, Production, Default Values)
- [x]  Support for managing logs (Format, Output, Level)
- [x]  Support for database operations (RDBMS, NoSQL)
- [x]  Support for Authentication with JWT (Sample Payload, Create, GetClaims)
- [x]  Support for working with Kafka (Schema, Broker, Consumer)
- [x]  Support for Redis operations (Connect, Get, Set, Invalidate)
- [x]  Support for AWS S3 operations (Connect, Get, Upload)
- [x]  Support for Elasticsearch (Connect, Insert, Search)
- [x]  Support for middleware (Rate Limit, Blocking)
- [x]  Easy integration with HTTP request handling (Status, Response, graceful shutdown)
- [ ]  Support for GRPC
- [ ]  Support for Websocket
- [ ]  Support for Ethereum

## Demo of a Basic System

- Support standard authentication functions (Sign In, Sign Up, Change Password, Get Profile)
- Support admin functions: Get all users, Block IP, Upgrade role, Rate Limit
- Support sending messages to Kafka brokers and consuming these messages,then inserting them into Elasticsearch
- Support full-text search using Elasticsearch's API
You can import this API collection into Postman as a JSON file using this postman.json file

![Image](https://res.cloudinary.com/dtmebo99b/image/upload/v1697304940/github/base_yrbzf0.png)
## How to Run
To install database, redis, kafka, elasticsearch. I defined in docker-compose.yml. You need run to start app 
```text
docker-compose up --build -d
```
Then we run go app with Makefile
```text
make run
```
## Prerequisites
- Go version 1.18

## Contributing or Maintaining
This project may have several issues, although I'm not aware of them at the moment. You can contribute to the go-base-service project by submitting documentation issues and pull requests to the repository.

### Thanks for visting me