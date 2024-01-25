# RouteHUB GraphQL Backend Service

## About the Project

This project is a B2B Link Shortener platform designed to empower businesses by providing a customizable and feature-rich solution for URL shortening.
It offers a range of functionalities catering to the specific needs of organizations, allowing them to enhance brand visibility and control over their short links.

## About Open Source

This project showcases various examples of using GO & Gqlgen for GraphQL development.
You can easily adapt this project for your own needs or contribute to its improvement.
However, due to my current workload, I cannot implement all the features that I have planned for this project.
I also have a SSR example of this project that has similar functionality to this one. The idea is to use this GraphQL backend with two different frontend projects: a Dashboard and a Client.
The Dashboard allows you to create and manage hubs using this project, while the Client is deployed for each hub and provides analytics and template integration.
Currently, I have a Dashboard and a Client, but they are SSR go applications. The main goal is to make new applications with uses this GraphQL backend.

## Key Features

- **Custom Domains:** Users can add and manage custom domains for personalized shortening.
- **Custom Redirection Page:** Design tailored redirection pages for custom domains.
- **Homepage for Custom Shortener Domain:** Dedicated hub for managing links and settings.
- **Pinned Links:** Pin important or frequently accessed links for easy access.
- **Redirection Options:** Flexible options for link behavior customization.
- **Organization Authorities:** Robust permission system for managing team access.
- **Monitoring Requests:** Track and analyze traffic generated by shortened links.
- **SEO Crawling:** Shortened links are SEO-friendly and can be crawled by search engines.
- **Customizable SEO & OG Tags:** Users can customize SEO tags Open Graph (OG) images for enhanced branding. for each link.

## Key Handlers

- **Auth:** Customers can login, register and reset passwords.
- **Organization:** Customers can create, update and delete organizations. `M2M relation with users.`
- **Platform:** Customers can create, update and delete platforms. `M2M relation with organizations.`
- **Domain:** Customers can create, update and delete domains. `O2O relation with platforms.`
- **Link:** Customers can create, update and delete links. `O2O relation with Platform.`
- **Crawl:** Links can be crawled.
- **Permissions:** Customers can invite another person to their organization and assign them custom permissions.

## Implemented

- [x] **Auth:** Customers can login, register.
- [x] **Organization:** Customers can create, update organizations. `M2M relation with users.`
- [x] **User Invitation:** Customers can invite another person to their organization `M2M`.
- [x] **User Invitation Acceptance:** Invited persons can accept or decline invitations.
- [x] **Permissions;** Every platform and organization has special permissions. Querys and mutations are protected with directives.
- [x] **Platform:** Customers can create, update their platforms.
- [x] **Domain:** Every Platform must has a domain for deployment and link shortening with hub deployment process.
- [x] **Link:** Customers can create shorlink's with custom SEO tags and OG images.
- [x] **Crawl:** Links will be crawled automatically and if user want's to recrawl can create a crawl request and monitor process in link query. Crawling process is made with colly web scraping framework.
- [x] **Data loaders:** Data loaders are implemented for users.
- [x] **LRU Expirable Cache:** LRU Cache is implemented for data loaders.
- [x] **Application Configuration:** Application configuration is implemented with yaml file.
- [x] **Database Migration / Seeding / Multiple Provider:** Database migration and seeding is implemented with gorm and embeded postgre / postgresql or mysql database.
- [x] **GraphQL Playground:** GraphQL playground is implemented for development and testing is also configurable.
- [x] **GraphQL Directives:** GraphQL directives are implemented for authentication, domain url check, platform and organization permissions.
- [x] **GraphQL Relay:** GraphQL Relay is implemented for pagination and filtering purposes. `Link Conncetion.`
- [x] **Service Container:** Service container is implemented for accessing services in resolvers.
- [x] **Data Access Layer:** Data access layer is implemented for accessing database models in services.
- [x] **Data Loader Services;** Data loader services are implemented for accessing data loaders in services.
- [x] **Custom Url Validation with Configuration:** Custom url validation is implemented with configuration at utils/url_validator.go.

## Not Implemented

- [ ] **Analytics:** Link analytics are not implemented.
- [ ] **Deployments:** Platform Deployments are not implemented.
- [ ] **Link Pinning;** Link pinning for home page of platform is not implemented.
- [ ] **Queing:** Queing for crawling is not implemented.
- [ ] **User Verification:** User verification is not implemented.
- [ ] **Domain Ownership Verification:** Domain verification is not implemented.
- [ ] **Platform Redirection Templates:** Platform redirection page html templates are not implemented.
- [ ] **JWT RSA:** JWT RSA is not implemented.
- [ ] **Super Admin:** Super admin is not implemented. Some querys is not protected with directives.
- [ ] **File Upload:** File upload is not implemented for og image uploads.

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
  - `lrue`: Configuration for the Least Recently Used (LRU) cache. (Does not related to wait)
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
  - `provider_enum.go`: Enumerations for database provider configurations.

- `/database`: Contains database related code.
  - `/configure`: Database configuration related code.
  - `/enums`: Enumerations used in the database.
  - `/models`: Database models.
  - `/relations`: Database relations.
  - `/types`: Database JSON Object Types.
  - `connection.go`: Handles database connection.
  - `embeded_postgre.go`: Handles embedded PostgreSQL database.
  - `mock.go`: Mock data for testing.
  - `migrate.go`: Handles database migrations.

- `/directives`: Contains graphql directive related code.
  - `domain_url_check.go`: Directive for domain URL check. (Checks the domain from database, domain must be unique)
  - `organization_permission.go`: Directive for organization permission.
  - `auth_directive.go`: Directive for authentication.
  - `platform_permission.go`: Directive for platform permission.
  - `assign.go`: Directive middleware for assignment to graphql.

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
