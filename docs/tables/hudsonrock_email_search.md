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
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Get top passwords and logins from the first stealer

```sql+postgres
select
  email,
  top_passwords,
  top_logins
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
  top_passwords,
  top_logins
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Get antiviruses and malware path from the first stealer

```sql+postgres
select
  email,
  antiviruses,
  malware_path
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
  antiviruses,
  malware_path
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```
