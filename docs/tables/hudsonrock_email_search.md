---
title: "Steampipe Table: hudsonrock_email_search"
description: "Query Hudson Rock infostealer and credential data by email with SQL."
---

# Table: hudsonrock_email_search

The `hudsonrock_email_search` table allows you to query compromised credentials and infostealer data by email using the Hudson Rock API.

## Table Usage Guide

The `hudsonrock_email_search` table provides insights about compromised credentials, infostealer malware, and related data for a given email address.

**Important Notes**
- You must provide an `email` qualifier in the `where` clause for all queries.

## Examples

### Basic email intelligence

```sql
select
  email,
  message,
  date_compromised,
  stealer_family,
  computer_name,
  operating_system,
  ip
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Get all data for an email

```sql
select
  *
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### List top passwords and logins for an email

```sql
select
  email,
  top_passwords,
  top_logins
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

## Columns

| Name                    | Type    | Description                                      |
|-------------------------|---------|--------------------------------------------------|
| email                   | string  | Email searched.                                  |
| message                 | string  | API message about the email.                     |
| date_compromised        | timestamp | Date the credentials were compromised.           |
| stealer_family          | string  | Infostealer malware family.                      |
| computer_name           | string  | Name of the infected computer.                   |
| operating_system        | string  | Operating system of the infected computer.       |
| malware_path            | string  | Path to the malware on the system.               |
| antiviruses             | json    | Antivirus software detected.                     |
| ip                      | string  | IP address of the infected machine.              |
| top_passwords           | json    | Top passwords found.                             |
| top_logins              | json    | Top logins found.                                |
| total_corporate_services| int     | Total corporate services found.                  |
| total_user_services     | int     | Total user services found.                       |
| data                    | json    | Raw data from the API response.                  |
