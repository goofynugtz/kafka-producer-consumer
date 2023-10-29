# Prducer-Consumer with Kafka

Prerequisite: Docker and Golang compiler  
Ensure to have docker up && running.

1. To start working with the repository, start the zookeeper and kafka-server containers

```bash
make kafka
```

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
