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
  date_compromised,
  stealer_family,
  computer_name,
  operating_system,
  ip
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  message,
  date_compromised,
  stealer_family,
  computer_name,
  operating_system,
  ip
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

### Get all data for a username

```sql+postgres
select
  *
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

```sql+sqlite
select
  *
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

### List top passwords and logins for a username

```sql+postgres
select
  username,
  top_passwords,
  top_logins
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  top_passwords,
  top_logins
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

### Get antivirus and malware path details for a username

```sql+postgres
select
  username,
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
  antiviruses,
  malware_path
from
  hudsonrock_username_search
where
  username = 'johndoe';
```
