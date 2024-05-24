# ranked-pick-api

## Introduction

`ranked-pick-api` is the backend for [Ranked Pick](http://rankedpick.com), a web application for creating and responding to ranked choice style surveys.

Web client source can be found [here](https://github.com/carterjackson/ranked-pick-web).

## Development

### Database Management

See the [database doc](./docs/database.md) for DB management.

### Docker

The service is run locally with [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/).

The `--service-ports` flag is required to open ports to localhost.

```bash
docker compose run --rm --build --service-ports rp-api
```

## Testing

TODO
