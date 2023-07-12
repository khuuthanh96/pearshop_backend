# pearshop_backend

## Start database
```
docker-compose up --build -d
```
## Run server
```
make up
```

## Run test
```
make test
```

## Description
Project has 3 main layer:
1. Layer delivery: contain all api handler & middleware, interact mainly with use-case layer
2. Layer usecase: contain all the bussiness logic of the app, querying data, connecting to 3rd services, ...
3. Layer domain: contain all the entity, repository interfaces, services interfaces.
   - entity: contain data object of the app, it maybe come from database, redis, 3rd service.
   - repository: contain interface that define methods to interact with entity objects.
   - services: contain interface that define methods to interact with 3rd services.
4. External: Implementation of repository & services

Libraries used:
* Gin: Http web framework.
* Sql-migrate: Migrate database version.
* Gomock: Generates mock database for testing api.
* Google/Wire: Generates dependency injection for use-cases.
* Swaggo: Generates api docs
* Go-validator: Validates request payload with custom error messages
