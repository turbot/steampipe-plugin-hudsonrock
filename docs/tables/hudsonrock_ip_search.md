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
  total_corporate_services,
  total_user_services
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

```sql+sqlite
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

### List stealer data for an IP

```sql+postgres
select
  ip,
  stealers
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  ip,
  stealers
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
