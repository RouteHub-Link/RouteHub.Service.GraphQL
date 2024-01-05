# RouteHUB GraphQL Backend Service

## About
This project...

## Configuration Guide

This guide provides an overview of the settings in the `config/config.yaml` file.

### Database Configuration

- `host`: The hostname of your database server.
- `port`: The port number on which your database server is listening.
- `user`: The username for your database.
- `password`: The password for your database.
- `database`: The name of your database.
- `application_name`: The name of your application.
- `type`: Configuration for database migration and seeding.
  - `migrate`: If set to true, the database schema will be migrated upon application startup.
  - `seed`: If set to true, the database will be seeded with initial data upon application startup.
  - `provider`: The database provider. Set this to `embed` for an embedded database.

### GraphQL Configuration

- `port`: The port number on which your GraphQL server will run.
- `playground`: If set to true, the GraphQL playground will be enabled. This is useful for development and testing.
- `dataloader`: Configuration for the GraphQL data loader.
  - `wait`: The maximum duration to wait before dispatching a batch load.
  - `cache`: If set to true, the data loader will cache data.
  - `lrue`: Configuration for the Least Recently Used (LRU) cache.
    - `size`: The maximum number of items that can be stored in the cache.
    - `expire`: The duration after which an item in the cache will expire.

Please adjust these settings according to your needs and environment.

## Project Structure

Brief project description goes here.

- `/auth`: Contains authentication related code.
  - `user_context.go`: Handles user context.
  - `jwt.go`: Handles JWT authentication.
  - `auth_middleware.go`: Middleware for authentication.

- `/clients`: Contains client related code.
  - `colly_client.go`: Handles interactions with the Colly web scraping framework.

- `/config`: Contains configuration files for the service.
  - `configuration.go`: Handles loading and parsing of configuration.
  - `config.yaml`: YAML configuration file.
  - `provider_enum.go`: Enumerations for provider configurations.

- `/database`: Contains database related code.
  - `/configure`: Database configuration related code.
  - `/enums`: Enumerations used in the database.
  - `/models`: Database models.
  - `/relations`: Database relations.
  - `/types`: Database Json Object Types.
  - `connection.go`: Handles database connection.
  - `embeded_postgre.go`: Handles embedded PostgreSQL database.
  - `mock.go`: Mock database for testing.
  - `migrate.go`: Handles database migrations.

- `/directives`: Contains directive related code.
  - `domain_url_check.go`: Directive for domain URL check .
  - `organization_permission.go`: Directive for organization permission.
  - `auth_directive.go`: Directive for authentication.
  - `assign.go`: Directive for assignment.
  - `platform_permission.go`: Directive for platform permission.

- `/graph`: Contains GraphQL related code.
  - `/model`: GraphQL models.
    - `/connections`: Connection related models.
      - `relay.go`: Relay connection model.
    - `/inputs`: Input related models.
    - `models_gen.go`: Generated models.
  - `/resolvers`: GraphQL resolvers.
  - `/schemas`: GraphQL schemas.
    - `sgorm-gqlgen-relay-generated.gql`: Generated schema for relay.
    - `schema.gql`: Base schema with some directives and enums.
  - `generated.go`: Generated GraphQL code.
- `/loaders`: Contains data loader related code.
  - `user_loader.go`: User data loader.
  - `loaders.go`: Base data loader.
- `/services`: Contains business logic related code.
  - `/domain`: Domain related services.
  - `/link`: Link related services.
  - `/organization`: Organization related services.
  - `/platform`: Platform related services.
  - `/user`: User related services.
  - `/utils`: Utility services.
    - `lru_cache.go`: Provides a Least Recently Used (LRU) cache.
    - `hash_service.go`: Provides hashing services.
    - `lru_expireable_cache.go`: Provides an LRU cache with expiration.
    - `url_validator.go`: Provides URL validation services.
  - `service_container.go`: Contains the service container that manages all services.

- `generate.go`: Script to generate necessary code.
- `go.mod` and `go.sum`: Define the module's dependencies.
- `main.go`: The entry point for the service.
- `readme.md`: Documentation for the project.
- `server.go`: Handles server setup and configuration.
- `tools.go`: Contains code for tools used in the project.
- `gqlgen.yml`: Configuration file for gqlgen, a Go library for building GraphQL servers.
- `.gitignore`: Specifies intentionally untracked files to ignore when using Git.

Please note that after any changes to the GraphQL schema, you must run `go run -mod=mod github.com/99designs/gqlgen` to regenerate the necessary code.
