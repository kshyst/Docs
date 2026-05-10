# Viewing Logs with Loki

Grafana provides powerful tools for viewing and analyzing logs, primarily through integration with Loki.

## The Explore View

The **Explore** view is the best place for ad-hoc log querying and debugging.

1.  Click **Explore** in the sidebar.
2.  Select your **Loki** data source.
3.  Use the query builder or write LogQL directly.
4.  View results in the logs viewer, which supports filtering, highlighting, and context viewing.

## LogQL Basics

LogQL is Loki's query language, inspired by PromQL.

-   **Log Selectors**: `{job="varlogs"}`
-   **Line Filters**: `{job="varlogs"} |= "error"`
-   **JSON Parsing**: `{job="varlogs"} | json`
-   **Metric Queries from Logs**: `rate({job="varlogs"} |= "error" [5m])`

## Adding Log Panels to Dashboards

You can add a "Logs" panel to any dashboard to see real-time log data alongside your metrics.

1.  Add a new panel to your dashboard.
2.  Select **Loki** as the data source.
3.  Write your LogQL query.
4.  The visualization will automatically switch to the "Logs" view.
5.  Configure options like "Show labels", "Show timestamp", and "Enable log details".
