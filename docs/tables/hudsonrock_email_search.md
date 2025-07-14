---
title: "Steampipe Table: hudsonrock_email_search"
description: "Query Hudson Rock infostealer and credential data by email with SQL."
---

# Table: hudsonrock_email_search - Query Hudson Rock Email Intelligence using SQL

The `hudsonrock_email_search` table allows you to query compromised credentials and infostealer data by email using the Hudson Rock API. This table provides a detailed view of infostealer infections, password and login exposures, stealer malware families, and more for a given email address.

## Table Usage Guide

The `hudsonrock_email_search` table provides insights about compromised credentials, infostealer malware, and related data for a given email address, including compromise date, stealer family, computer and OS details, IP address, and exposed passwords and logins.

**Important Notes**
- You must provide an `email` qualifier in the `where` clause for all queries.
- This table is not intended for listing all emails, but for targeted intelligence on specific email addresses.

## Examples

### Basic email intelligence

```sql+postgres
select
  email,
  message,
  stealers
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
  message,
  stealers
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Unnest stealer details 

```sql+postgres
select
  email,
  jsonb_array_elements(stealers) as stealer_detail
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
  json_each(stealers) as stealer_detail
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Get top passwords and logins from the first stealer

```sql+postgres
select
  email,
  stealers->0->'top_passwords' as top_passwords,
  stealers->0->'top_logins' as top_logins
from
  hudsonrock_email_search
where
  email = 'user@example.com';;
```

```sql+sqlite
select
  email,
  json_extract(stealers, '$[0].top_passwords') as top_passwords,
  json_extract(stealers, '$[0].top_logins') as top_logins
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```
