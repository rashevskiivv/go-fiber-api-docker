# API to work with products
This project is a pet project to try fiber. Here's a CRUD for 1 entity.

In project were used the following libraries:
* [Fiber](https://docs.gofiber.io/) - Routing
* [Viper](https://github.com/spf13/viper) - Managing the environment variables
* [GORM](https://gorm.io/) - ORM
* Postgres - DB
* Docker with volumes + Docker-compose - Containerization
* [Zap](https://github.com/uber-go/zap) - Logging

## Docker
1. Clone this repository.
2. Run: ```docker compose up```
3. Stop in different window: ```docker compose down``` OR CTRL+C to terminate the process in the same window.