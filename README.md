# Experimenting with SSE

# TODO
- setup webex
- setup webook
- test sending requests to grafana to display
- set up how to check out events from
- handle webhook post events

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
