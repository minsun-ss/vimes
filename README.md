# Database Setup

In order to set up the database, run the logs.sql query. Hopefully migrations are not needed at this point. You can go about this a couple of ways:

- if you have clickhouse-client CLI, run the setup.sh in /clickhouse folder; or
- run logs.sql inside a SQL client of your choice (presuming you have the correct privileges)


# Experimenting with SSE

# TODO
- setup webex
- setup and test webhook
- test sending requests to grafana to display
- set up how to check out events from
- handle webhook post events
- viper configuration



# handling post events
- secured with hmac
- fire off emails
- quickly poll logs or generate stats
- generate charts of activity in last events
- fire off webex alerts
- check own logs of the application

# webex bot
- call some of these events
- generate stats
- check tempo hours

# website
- streaming logs
- statues
- embed grafanas if there
- should there be an accompaning database for this?
- google console logs TBD

# Other things to handle
- graceful degradation: e.g., I don't necessarily need opensearch or google working at all times
- webex bot failure/retry
- configmap setup
- prometheus alerting
- liveness and readiness checks

# What is vimes
I wanted a lightweight way to implement an interface that lets me monitor
logs of my own applications and send out alerts.

Turns out to be less lightweight than I thought...
