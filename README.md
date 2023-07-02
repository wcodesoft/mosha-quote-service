# mosha-quote-service

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