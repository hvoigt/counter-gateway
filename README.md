# Counter Gateway

A REST API that allows to increment Prometheus counter metrics which are
exposed on the /metrics endpoint.

## Usage

### Start the server locally

```bash
make run
```

### Create/Increment a counter metric

```bash
curl 'http://localhost:8080/increment?counter=some_counter&label=bla%3Dblub&label=hallo%3Dwelt'
```

### Get the metrics

```bash
curl localhost:8080/metrics
```

### Build the Docker image

```bash
make docker
```
