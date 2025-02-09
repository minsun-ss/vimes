CREATE database vimes;

CREATE TABLE vimes.logs (
    event_source String,
    timestamp_micros UInt64,
    message String
) ENGINE = MergeTree
ORDER BY
    (timestamp_micros, event_source) TTL toDateTime (timestamp_micros / 1000000) + INTERVAL 1 YEAR SETTINGS index_granularity = 8192
