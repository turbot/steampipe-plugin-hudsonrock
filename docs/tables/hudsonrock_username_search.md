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
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

#### Find usernames with more than 3 top passwords

```sql+postgres
select
  username,
  top_passwords,
  jsonb_array_length(top_passwords) as num_passwords
from
  hudsonrock_username_search
where
  username = 'johndoe'
  and jsonb_array_length(top_passwords) > 3;
```

```sql+sqlite
select
  username,
  top_passwords,
  json_array_length(top_passwords) as num_passwords
from
  hudsonrock_username_search
where
  username = 'johndoe'
  and json_array_length(top_passwords) > 3;
```

#### Search for a specific antivirus in the antiviruses array

```sql+postgres
select
  username,
  antiviruses
from
  hudsonrock_username_search
where
  username = 'johndoe'
  and antiviruses::text ilike '%Kaspersky%';
```

```sql+sqlite
select
  username,
  antiviruses
from
  hudsonrock_username_search,
  json_each(hudsonrock_username_search.antiviruses)
where
  ip = '8.8.8.8'
  and lower(json_each.value) = 'kaspersky';;
```

#### Order by number of top passwords (descending)

```sql+postgres
select
  username,
  top_passwords,
  jsonb_array_length(top_passwords) as num_passwords
from
  hudsonrock_username_search
where
  username = 'johndoe'
order by
  num_passwords desc;
```

```sql+sqlite
select
  username,
  top_passwords,
  json_array_length(top_passwords) as num_passwords
from
  hudsonrock_username_search
where
  username = 'johndoe'
order by num_passwords desc;
```
