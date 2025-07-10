---
title: "Steampipe Table: hudsonrock_username_search"
description: "Query Hudson Rock infostealer and credential data by username with SQL."
---

# Table: hudsonrock_username_search

The `hudsonrock_username_search` table allows you to query compromised credentials and infostealer data by username using the Hudson Rock API.

## Table Usage Guide

The `hudsonrock_username_search` table provides insights about compromised credentials, infostealer malware, and related data for a given username.

**Important Notes**
- You must provide a `username` qualifier in the `where` clause for all queries.

## Examples

### Basic username intelligence

```sql
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

```sql
select
  *
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

### List top passwords and logins for a username

```sql
select
  username,
  top_passwords,
  top_logins
from
  hudsonrock_username_search
where
  username = 'johndoe';
```
