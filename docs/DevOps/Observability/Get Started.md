# Observability

## ELK Stack

Elasticsearch, Logstash, and Kibana

### ElasticSearch (The Search & Storage Engine)

-  A highly scalable, distributed, RESTful search and analytics engine built on Apache Lucene.
- It acts as the heart of the ELK stack. It stores the data and provides lightning-fast search capabilities. It is particularly well-suited for indexing large volumes of structured and unstructured data (like logs).

### Logstash (The Data Processing Pipeline)

- A server-side data processing pipeline.
- It ingests data from a multitude of sources simultaneously, transforms it, and then sends it to a “stash” like Elasticsearch. It can filter logs, parse complex data (like IP addresses or timestamps), and format the data so it is uniformly structured before it gets stored.

### Kibana (The Visualization Tool)

- A data visualization and exploration dashboard.
-  It serves as the user interface for the ELK stack. Once the data is stored in Elasticsearch, Kibana allows users to search, view, and interact with the data through beautifully designed charts, graphs, maps, and tables.

## Prometheus + Grafana + Loki + Tempo

