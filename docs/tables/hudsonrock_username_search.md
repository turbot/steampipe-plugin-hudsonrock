---
title: "Steampipe Table: hudsonrock_username_search"
description: "Query Hudson Rock infostealer and credential data by username with SQL."
---

# Table: hudsonrock_username_search - Query Hudson Rock Username Intelligence using SQL

The `hudsonrock_username_search` table allows you to query compromised credentials and infostealer data by username using the Hudson Rock API. This table provides a detailed view of infostealer infections, password and login exposures, stealer malware families, and more for a given username.

## Table Usage Guide

The `hudsonrock_username_search` table provides insights about compromised credentials, infostealer malware, and related data for a given username, including compromise date, stealer family, computer and OS details, IP address, and exposed passwords and logins.

**Important Notes**
- You must provide a `username` qualifier in the `where` clause for all queries.
- This table is not intended for listing all usernames, but for targeted intelligence on specific usernames.

## Examples

### Basic username intelligence

```sql+postgres
select
  username,
  message,
  stealers
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  message,
  stealers
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

### Unnest stealer details

```sql+postgres
select
  username,
  jsonb_array_elements(stealers) as stealer_detail
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  json_each(stealers) as stealer_detail
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

### Get top passwords and logins from the first stealer

```sql+postgres
select
  username,
  stealers->0->'top_passwords' as top_passwords,
  stealers->0->'top_logins' as top_logins
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  json_extract(stealers, '$[0].top_passwords') as top_passwords,
  json_extract(stealers, '$[0].top_logins') as top_logins
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

### List top passwords and logins for a username

```sql+postgres
select
  username,
  stealer_detail->'top_passwords' as top_passwords,
  stealer_detail->'top_logins' as top_logins
from
  hudsonrock_username_search,
  lateral jsonb_array_elements(stealers) as stealer_detail
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  stealer_detail.value -> '$.top_passwords' as top_passwords,
  stealer_detail.value -> '$.top_logins' as top_logins
from
  hudsonrock_username_search,
  json_each(stealers) as stealer_detail
where
  username = 'johndoe';
```

### Get antivirus and malware path details for a username

```sql+postgres
select
  username,
  stealer_detail->'antiviruses' as antiviruses,
  stealer_detail->'malware_path' as malware_path
from
  hudsonrock_username_search,
  lateral jsonb_array_elements(stealers) as stealer_detail
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  json_extract(stealer_detail.value, '$.antiviruses') as antiviruses,
  json_extract(stealer_detail.value, '$.malware_path') as malware_path
from
  hudsonrock_username_search,
  json_each(stealers) as stealer_detail
where
  username = 'johndoe';
```
