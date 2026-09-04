# Epoch Time in Databases

## What is epoch time?

Unix epoch time represents an instant as the number of seconds since
`1970-01-01 00:00:00 UTC`. Values before that instant are negative.

```text
-1           = 1969-12-31 23:59:59 UTC
0            = 1970-01-01 00:00:00 UTC
1700000000   = 2023-11-14 22:13:20 UTC
```

An epoch value does not contain a time zone. Time zones only matter when a
date-time is converted to an epoch or an epoch is formatted for display.
Unix time also does not encode leap seconds.

## Units

The unit is part of the data contract; the number alone is ambiguous.

| Unit | Values per second | Example near 2023 |
| --- | ---: | ---: |
| Seconds | `1` | `1700000000` |
| Milliseconds | `1,000` | `1700000000000` |
| Microseconds | `1,000,000` | `1700000000000000` |
| Nanoseconds | `1,000,000,000` | `1700000000000000000` |

Do not guess the unit from the number of digits in production code. Document
it explicitly and use names such as `created_at_epoch_ms`. Use `BIGINT` for
integer epoch values; a signed 32-bit integer overflows on
`2038-01-19 03:14:08 UTC`.

## Prefer native date-time types

Store an instant using the database's native, time-zone-aware type when one is
available. For example, PostgreSQL's `timestamptz` preserves the instant and
supports date arithmetic, validation, formatting, and indexes directly.

```sql
CREATE TABLE events (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    occurred_at timestamptz NOT NULL DEFAULT now()
);
```

Convert to epoch at an API or system boundary instead of storing a duplicate
epoch column. Store epoch integers when an external protocol requires them or
when the source data is already defined that way.

## PostgreSQL

### Convert a timestamp to epoch seconds

```sql
SELECT EXTRACT(EPOCH FROM TIMESTAMPTZ '2023-11-14 22:13:20+00');
-- 1700000000.000000

SELECT EXTRACT(EPOCH FROM now());
```

For whole milliseconds, multiply before converting to `bigint`:

```sql
SELECT floor(EXTRACT(EPOCH FROM occurred_at) * 1000)::bigint AS occurred_at_epoch_ms
FROM events;
```

### Convert epoch seconds to a timestamp

`to_timestamp()` accepts seconds, including a fractional part:

```sql
SET TIME ZONE 'UTC';

SELECT to_timestamp(1700000000);
-- 2023-11-14 22:13:20+00

SELECT to_timestamp(1700000000123 / 1000.0);
-- 2023-11-14 22:13:20.123+00
```

`to_timestamp()` returns `timestamptz`; its displayed offset follows the
session time zone.

### Timestamp without a time zone

PostgreSQL treats `EXTRACT(EPOCH FROM timestamp)` as nominal seconds and does
not apply time-zone or daylight-saving rules. If a timezone-less value really
represents local time, attach its known zone first:

```sql
SELECT EXTRACT(
    EPOCH FROM (local_created_at AT TIME ZONE 'Asia/Tehran')
)
FROM events;
```

Do not attach the server's current time zone implicitly; the result could
change between environments.

### Keep timestamp indexes usable

Apply conversion to the input parameter, not to every value in an indexed
column:

```sql
-- Uses a normal index on occurred_at.
SELECT *
FROM events
WHERE occurred_at >= to_timestamp($1);

-- Usually cannot use that normal index directly.
SELECT *
FROM events
WHERE EXTRACT(EPOCH FROM occurred_at) >= $1;
```

## MySQL

MySQL uses the session time zone when converting between date-time values and
Unix time. Set it explicitly when UTC behavior is required.

```sql
SET time_zone = '+00:00';

SELECT UNIX_TIMESTAMP('2023-11-14 22:13:20');
-- 1700000000

SELECT FROM_UNIXTIME(1700000000);
-- 2023-11-14 22:13:20
```

## SQLite

SQLite can store date-time values as text, Julian day numbers, or Unix
timestamps. Use the `unixepoch` modifier when reading an epoch value.

```sql
SELECT unixepoch('2023-11-14T22:13:20Z');
-- 1700000000

SELECT datetime(1700000000, 'unixepoch');
-- 2023-11-14 22:13:20
```

## Common mistakes

- Mixing seconds and milliseconds.
- Parsing a local date-time without specifying its time zone.
- Storing epoch values in a signed 32-bit integer.
- Using floating-point numbers for precise microsecond or nanosecond values.
- Applying a conversion function to an indexed timestamp column in every query.
- Treating epoch seconds as a calendar duration; use native date/interval
  operations for months, daylight-saving transitions, and calendar rules.

At application boundaries, validate the unit and an acceptable date range
before saving an epoch supplied by a client.

## References

- [PostgreSQL date/time functions](https://www.postgresql.org/docs/current/functions-datetime.html)
- [MySQL date and time functions](https://dev.mysql.com/doc/refman/8.4/en/date-and-time-functions.html)
- [SQLite date and time functions](https://www.sqlite.org/lang_datefunc.html)
