# RouteHUB GraphQL Backend Service

![Design](https://github.com/RouteHub-Link/RouteHub.Service.GraphQL/assets/16222645/e747f0af-8e52-4309-9cf0-169f6c3cd750 "Routehub Banner")

## About the Project

## What does?

Route Hub is a specialized redirection solution designed with businesses in mind. In today's digital landscape, security, brand trust, and the use of redirection links can often be a concern. Many companies worry about the implications of sending users through a redirection page, fearing it may erode trust or seem unprofessional. But what if a redirection page could do more than just redirect? What if it could enhance your brand's image and provide valuable insights?

#### This is the core idea behind Route Hub

Route Hub transforms ordinary redirection pages into powerful brand link hubs. Instead of merely sending users from point A to point B, Route Hub allows you to manage your most frequently redirected links and discover new opportunities to engage with your audience. Each redirection becomes a touchpoint where you can reinforce your brand, gather analytics, and optimize your marketing strategies.

With Route Hub, you don't just redirect traffic—you direct it with purpose, ensuring that every click counts and every redirection page reflects your brand's integrity and professionalism. Whether it's for security, brand consistency, or uncovering new insights, Route Hub elevates the way businesses handle redirections.

- Business's can create their hub's with validated domains.
- Employee Invitation & permissions.
- Can create short link's with some options. (Scrape SEO, Custom SEO, Redirection Choices etc.)
- Link's validated automatically every link must be valid, reachable and not a redirection link.
- Hub's can list most populer links, pinned short link's, custom css designs and also business informations.

## Documentation

[Chek out wiki page.](https://github.com/RouteHub-Link/RouteHub.Service.GraphQL/wiki)

## Quick Start

```shell
## Host Domain Utility Asynq Task Manager & API (For validation)
git clone https://github.com/RouteHub-Link/DomainUtils
cd DomainUtils
docker compose up -d
cd ..

## Run GraphQL API
git clone https://github.com/RouteHub-Link/RouteHub.Service.GraphQL
cd RouteHub.Service.GraphQL
go run .
```

For deployment and more information, please check the [HELM Repository](https://github.com/RouteHub-Link/RouteHub.HELM "RouteHub.HELM") & [Artifacts Page](https://artifacthub.io/packages/search?repo=routehub-helm&sort=relevance&page=1 "Artifact Hub").

## Reminder Note

Please note that after any changes to the GraphQL schema, you must run `go run -mod=mod github.com/99designs/gqlgen` to regenerate the necessary code.
