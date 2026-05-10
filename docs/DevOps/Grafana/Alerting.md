# Grafana Alerting

Grafana Alerting (Unified Alerting) allows you to learn about problems in your systems moments after they occur.

## Alert Rules

Alert rules define the conditions that trigger an alert.

1.  Navigate to **Alerting** -> **Alert rules**.
2.  Click **Create alert rule**.
3.  Define the query and the condition (e.g., `avg() of query A is above 80`).
4.  Set the evaluation interval and the "For" duration (how long the condition must be met before firing).
5.  Add metadata like labels and annotations (Summary, Description).

## Contact Points

Contact points define *where* the notifications are sent.

1.  Navigate to **Alerting** -> **Contact points**.
2.  Click **Add contact point**.
3.  Choose an integration type:
    - **Slack**: Provide a Webhook URL or use a Bot Token.
    - **Discord**: Provide a Webhook URL.
    - **Email**: Configure SMTP settings in Grafana configuration.
    - **PagerDuty**, **Opsgenie**, etc.
4.  Test the contact point to ensure it's working.

## Notification Policies

Notification policies define *how* and *when* alerts are routed to contact points.

-   **Root Policy**: The default policy for all alerts.
-   **Specific Policies**: Use label matchers to route specific alerts to different contact points (e.g., route "critical" alerts to PagerDuty and "warning" alerts to Slack).
-   **Muting**: You can create "Mute timings" to suppress notifications during scheduled maintenance.
