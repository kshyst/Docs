# Introduction to Grafana

Grafana is a multi-platform open-source analytics and interactive visualization web application. It provides charts, graphs, and alerts for the web when connected to supported data sources.

## What is Grafana?

At its core, Grafana is a visualization tool that allows you to query, visualize, alert on, and understand your metrics no matter where they are stored. It enables you to create, explore, and share dashboards with your team and foster a data-driven culture.

## Core Concepts

### Dashboards
A dashboard is a set of one or more panels organized and arranged into one or more rows. Grafana ships with a variety of panels. Grafana makes it easy to construct the right queries and customize the display settings so that you can create the dashboard you need. Each panel can interact with data from any configured Grafana data source.

### Panels
Panels are the basic visualization building blocks of a dashboard. Each panel has a query editor specific to the data source selected in the panel. The query editor allows you to extract the visualization you want.

## The LGTM Stack

Grafana is a central piece of the **LGTM** stack, a set of open-source tools designed to provide a complete observability solution:

-   **Loki**: For logs.
-   **Grafana**: For visualization and alerting.
-   **Tempo**: For traces.
-   **Mimir**: For metrics.

Together, these tools provide a unified experience for monitoring and troubleshooting complex systems.
