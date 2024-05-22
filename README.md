# ranked-pick-api

## Introduction

`ranked-pick-api` is the backend for [Ranked Pick](http://rankedpick.com), a web application for creating and responding to ranked choice style surveys.

Web client source can be found [here](https://github.com/carterjackson/ranked-pick-web).

## Development

See the [database doc](./docs/database.md) for DB management.

### Docker

The service is run locally through [Docker](https://www.docker.com/) with [Docker Compose](https://docs.docker.com/compose/).

```bash
docker compose run --rm --build --service-ports rp-api
```

## Testing

TODO
