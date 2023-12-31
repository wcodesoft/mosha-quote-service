# mosha-quote-service

[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/bb10938daac84a34a66c9a4be906720c)](https://app.codacy.com/gh/wcodesoft/mosha-quote-service/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/bb10938daac84a34a66c9a4be906720c)](https://app.codacy.com/gh/wcodesoft/mosha-quote-service/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Quote microservice used in Mosha

## Database

The main database used in the service is MongoDB. It's used to store the authors. To deploy it locally, run:

```bash
docker run --name mongo -p 27017:27017 -d mongodb/mongodb-community-server:latest 
```

## Docker

To build the container image, run:

```bash
docker build -t mosha-quote-service .
```

After that to run the container, run:

```bash
docker run --name mosha-quote-service -e MONGO_DB_HOST="mongodb://localhost:27017" --net=bridge -p 8180:8180 -d mosha-quote-service
```

## Sentry

This service uses Sentry to track errors. To test it locally set the environment variable `SENTRY_DSN` to your Sentry 
DSN before running the service.

```bash
export SENTRY_DSN="your-dsn"
```