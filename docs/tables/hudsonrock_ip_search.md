---
title: "Steampipe Table: hudsonrock_ip_search"
description: "Query Hudson Rock infostealer and credential data by IP address with SQL."
---

# Table: hudsonrock_ip_search

The `hudsonrock_ip_search` table allows you to query info-stealer and credential data by IP address using the Hudson Rock API.

## Table Usage Guide

The `hudsonrock_ip_search` table provides insights about compromised credentials, infostealer malware, and related data for a given IP address.

**Important Notes**
- You must provide an `ip` qualifier in the `where` clause for all queries.

## Examples

### Basic IP intelligence

```sql
select
  ip,
  message,
  total_corporate_services,
  total_user_services
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

### Get all data for an IP

```sql
select
  *
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

### List stealer data for an IP

```sql
select
  ip,
  stealers
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```
