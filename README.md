# Prducer-Consumer with Kafka

Prerequisite: Docker and Golang compiler  
Ensure to have docker up && running.

Exposed .env config for the sake of project submission.

1. To start working with the repository, start the zookeeper and kafka-server containers

```bash
make kafka
```
In-case `broker` container fails to start, do `sudo docker container rm zookeeper broker` to remove pre-existing container images and build again.

2. Build api and consumer bin files

```bash
make build
```

3. Start the api-server by

```bash
make deploy
```

4. Start consumer-server in a new terminal instance by

```bash
make consumer
```
This can result in error if the topic or partition in kafka does not exists. Please POST atleast 1 entry to /recieve endpoint before starting consumer.

## API Endpoints

GET `/ping`: Healthcheck endpoint, to check if server is online.  
POST `/recieve`: To recieves the product and user details.

Body:
```
{
  "user_id": 131,
  "product_name": "Sony WH1000XM4",
  "product_description": "Headphones",
  "product_images": ["https://source.unsplash.com/random"]
  "product_price": 196423
}
```
