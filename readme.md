# RouteHUB GraphQL Backend Service
![Design](https://github.com/RouteHub-Link/RouteHub.Service.GraphQL/assets/16222645/e747f0af-8e52-4309-9cf0-169f6c3cd750)

## About the Project

This project aim's create a redirection reliable & trustfull.
This is a B2B Link Shortener platform designed to empower businesses by providing a customizable and feature-rich solution for URL shortening.
It offers a range of functionalities catering to the specific needs of organizations, allowing them to enhance brand visibility and control over their short links. 

## Whot does?
- Business's can create their hub's with validated domains.
- Employee Invitation & permissions.
- Can create short link's with some options. (Scrape SEO, Custom SEO, Redirection Choices etc.)
- Link's validated automatically every link must be valid, reachable and not a redirection link.
- Hub's can list most populer links, pinned short link's, custom css designs and also business informations.
 

## Documentation
[Chek out wiki page.](https://github.com/RouteHub-Link/RouteHub.Service.GraphQL/wiki)

## Quick Start
```shell
## Host Domain Utility Asynq Task Manager & API
git clone https://github.com/RouteHub-Link/DomainUtils
cd DomainUtils
docker compose up -d
cd ..

## Run GraphQL API
git clone https://github.com/RouteHub-Link/RouteHub.Service.GraphQL
cd RouteHub.Service.GraphQL
go run .
```

## Reminder Note
Please note that after any changes to the GraphQL schema, you must run `go run -mod=mod github.com/99designs/gqlgen` to regenerate the necessary code.
