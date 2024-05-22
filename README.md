# ranked-pick-api

## Introduction

`ranked-pick-api` is the API source for [Ranked Pick](http://rankedpick.com), a web application designed for creating and responding to ranked choice style surveys.

**Features**

- Create and share ranked choice surveys
- Respond to surveys that have been shared with you
- **TODO**: Share surveys publicly, by link, or via invite
- **TODO**: Allow participants to nominate their own options

## Development

See the [database doc](./docs/database.md) for DB management.

### Go

`$GOPATH` and `$GOCACHE` are used by the container to speed up builds and avoid reinstallation.

### Docker

`$GOPATH`

```bash
docker compose run --service-ports rp-api
```

### Database Management


## Testing

TODO
