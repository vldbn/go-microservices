# Go - Microservices

## A sample microservice architecture application with a messaging-based communication style.

### Setup

- Create .env file in project root directory:

  ```.env
  USER_SERVICE_PORT=8000
  USER_DATABASE_URL=mongodb://users_service_db:27017
  USER_DATABASE=userDatabase

  ADDRESS_SERVICE_PORT=8001
  ADDRESS_DATABASE_URL=mongodb://addresses_service_db:27017
  ADDRESS_DATABASE=addressDatabase

  RABBIT_USERNAME=admin
  RABBIT_PASSWORD=admin
  RABBIT_HOST=rabbitmq
  ```

- Run commands:

```shell
docker-compose build
docker-compose up
```
