# fiber-comments
The repo adds useful information about creating a restful comment system in GO. This is built with ️Fiber, MongoDB, and of course, with lots of :heart:.

## Features
- Restful API
- Dockerized
- Using Fiber framework and MongoDB
- CI/CD
- JWT Authentication
- Documented
- Implemented the Repository Pattern

## Usage
Follow this steps:
```bash
    cp .env.example .env
    make up
```
#### If you don't want to use the api service on docker-compose.yml, follow this one:
- Comment the api service in docker-compose.yml
- Change DBURI value in the .env.example to mongodb://localhost:27017

```bash
    cp .env.example .env
    make up
    make run
```
