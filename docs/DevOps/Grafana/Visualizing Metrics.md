# Visualizing Metrics

Visualizing metrics is the primary use case for Grafana. This involves creating dashboards and panels that query data sources like Prometheus.

## Creating Panels

1.  Open a dashboard and click **Add** -> **Visualization**.
2.  Select your data source (e.g., Prometheus).
3.  Write your query in the query editor.
4.  Choose a visualization type (Graph, Stat, Gauge, Table, etc.) from the right sidebar.
5.  Configure panel options (Title, Description, Axis labels, Legend).

## PromQL Basics

When using Prometheus, you use PromQL (Prometheus Query Language) to retrieve data.

-   **Instant Vector**: `http_requests_total`
-   **Range Vector**: `http_requests_total[5m]`
-   **Rate Calculation**: `rate(http_requests_total[5m])`
-   **Filtering with Labels**: `http_requests_total{status="200", method="GET"}`
-   **Aggregation**: `sum(rate(http_requests_total[5m])) by (handler)`

## Organizing Dashboards

-   **Rows**: Use rows to group related panels. You can collapse rows to hide complexity.
-   **Library Panels**: If you have a panel you want to reuse across multiple dashboards, you can save it as a Library Panel.
-   **Annotations**: Add manual or automatic annotations (e.g., from deployment events) to provide context to your metric spikes.
