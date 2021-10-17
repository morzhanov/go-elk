# Go ELK example

Go ELK stack example with filebeat, metricbeat and ElasticSearch API.

<img src="https://i.ibb.co/w6pFWFH/IMG-0148.png" alt="arch"/>

## Features

- ELK stack setup
- Filebeat logs collector setup
- Metricbeat collector setup
- ElasticSearch API usage

## Structure

- `/cmd` - application setup
- `/config` - .env file with environment variables
- `/deploy`
  - `/filebeat` - filebeat.yml file for Filebeat setup 
  - `/metricbeat` - metricbeat.yml file for Metricbeat setup and Metricbeat Dockerfile
  - `docker-compose.yml` - docker-compose file with ELK stack and beats
- `/internal`
  - `/config` - config module based on viper package
  - `/doc` - ES Document interface
  - `/es` - ElasticSearch module for ES API implementation usage
  - `/generator` - generates Documents and sends them to ES using `es` module
  - `/logger` - application logger, creates file transport (for filebeat) and console transport
  - `/metrics` - setups application metrics handling via `expvar` package
  - `/rest` - application REST module

## How it works

1. ES API
   - application generate data with generator and saves them to ES
   - application registers public REST API for ES data manipulation
      - Update doc
      - Delete doc
      - Search docs
2. Filebeat
   - logger sends logs to `/logs` directory
   - filebeat crawl the logs and sends them to ES
3. Metricbeat
   - metrics package registers application metric data
     - go general metric data (memstat, cpu stats, etc.)
     - request count
     - error count
     - docs count
   - metricbeat consumes metrics from `localhost:8080/debug/vars` endpoint and sends them to the ES
4. Kibana - uses for data visualization

### Postman

Application contains `go-elk-example.postman_collection` file which you can import to your Postman app and perform requests to the application REST API.

## Kibana setup

### Create index patterns

Firstly you should create index patterns for ES data:
- `goelkmetr` - Go app metrics data
- `goelklog` - Go app logs data
- `docs` - ES Docs API data

Creating Index patterns:

<img src="https://i.ibb.co/QNp9PZz/1634497962209.png" alt="kibana-setup"/>

Index patterns created:

<img src="https://i.ibb.co/6HrCqDM/1634498016511.png" alt="kibana-setup"/>

Docs index example data:

<img src="https://i.ibb.co/FDKp4yC/1634498033442.png" alt="kibana-setup"/>

### Discover

After index patterns created we could search and filter data on the `Discover` page:

<img src="https://i.ibb.co/SXKzVfK/1634498970258.png" alt="kibana-setup"/>

Metrics data on `Discover` page:

<img src="https://i.ibb.co/B60nh52/1634498076514.png" alt="kibana-setup"/>

### Dashboard

Also, we could create a Kibana Dashboard and visualize ES data:
- error count
- request count
- documents count
- some go stats
- some docs stats

<img src="https://i.ibb.co/Yk8SBQg/1634498860795.png" alt="kibana-setup"/>
