# API to work with products
This project is a pet project to try fiber. Here's a CRUD for 1 entity.

In project were used the following libraries:
* Fiber - Routing
* Viper - Managing the environment variables
* GORM - ORM
* Postgres - DB
* Docker - Containerization

## Quick start
You can use Make (see Makefile for more information):

```make server``` to run app

```make build``` to build app into a binary file located in bin folder

```make d.up``` to run app in docker container

```make d.down``` to stop app in docker container

```make d.up.build``` to build container
