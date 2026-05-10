# Connecting Data Sources

Grafana supports a wide variety of data sources. Here is how to connect the most common ones.

## Prometheus (Metrics)

Prometheus is the most common data source for metrics.

1.  Navigate to **Connections** -> **Data Sources**.
2.  Click **Add data source**.
3.  Select **Prometheus**.
4.  Set the **URL** (e.g., `http://prometheus:9090`).
5.  Click **Save & test**.

## Loki (Logs)

Loki is used for log aggregation.

1.  Navigate to **Connections** -> **Data Sources**.
2.  Click **Add data source**.
3.  Select **Loki**.
4.  Set the **URL** (e.g., `http://loki:3100`).
5.  Click **Save & test**.

## Relational Databases (PostgreSQL / MySQL)

Grafana can also visualize data from relational databases.

1.  Navigate to **Connections** -> **Data Sources**.
2.  Click **Add data source**.
3.  Select **PostgreSQL** or **MySQL**.
4.  Enter the connection details:
    - **Host**: IP or hostname and port.
    - **Database**: Database name.
    - **User/Password**: Database credentials.
5.  Configure SSL settings if necessary.
6.  Click **Save & test**.

### SQL Querying
When using SQL databases, you can write standard SQL queries. Grafana also provides macros to handle time-series data, such as `$__timeFilter(column)`.
