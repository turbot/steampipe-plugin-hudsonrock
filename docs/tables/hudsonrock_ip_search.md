---
title: "Steampipe Table: hudsonrock_ip_search"
description: "Query Hudson Rock infostealer and credential data by IP address with SQL."
---

# Table: hudsonrock_ip_search - Query Hudson Rock IP Intelligence using SQL

The `hudsonrock_ip_search` table allows you to query info-stealer and credential data by IP address using the Hudson Rock API. This table provides a detailed view of infostealer infections, password and login exposures, stealer malware details, and more for a given IP address.

## Table Usage Guide

The `hudsonrock_ip_search` table provides insights about compromised credentials, infostealer malware, and related data for a given IP address, including compromise date, computer and OS details, malware path, antivirus products, and exposed passwords and logins.

**Important Notes**
- You must provide an `ip` qualifier in the `where` clause for all queries.
- This table is not intended for listing all IPs, but for targeted intelligence on specific IP addresses.

## Examples

### Basic IP intelligence

```sql+postgres
select
  ip,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  ip,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

### Get details of each stealer infection for an IP

```sql+postgres
select
  ip,
  date_compromised,
  computer_name,
  operating_system,
  malware_path,
  antiviruses,
  top_passwords,
  top_logins
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  ip,
  date_compromised,
  computer_name,
  operating_system,
  malware_path,
  antiviruses,
  top_passwords,
  top_logins
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

#### Find IPs with more than 3 top passwords

```sql+postgres
select
  ip,
  top_passwords,
  jsonb_array_length(top_passwords) as num_passwords
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8'
  and jsonb_array_length(top_passwords) > 3;
```

```sql+sqlite
select
  ip,
  top_passwords,
  json_array_length(top_passwords) AS num_passwords
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8'
  and json_array_length(top_passwords) > 3;
```

#### Search for a specific antivirus in the antiviruses array

```sql+postgres
select
  ip,
  antiviruses
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8'
  and antiviruses::text ilike '%Kaspersky%';
```

```sql+sqlite
select
  ip,
  antiviruses
from
  hudsonrock_ip_search,
  json_each(hudsonrock_ip_search.antiviruses)
where
  ip = '8.8.8.8'
  and lower(json_each.value) = 'kaspersky';;
```

#### Order by number of top passwords (descending)

```sql+postgres
select
  ip,
  top_passwords,
  json_array_length(top_passwords) AS num_passwords
from
  hudsonrock_ip_search
order by
  num_passwords desc;
```

```sql+sqlite
select
  ip,
  top_passwords,
  json_array_length(top_passwords) AS num_passwords
from
  hudsonrock_ip_search
order by
  num_passwords desc;
```
