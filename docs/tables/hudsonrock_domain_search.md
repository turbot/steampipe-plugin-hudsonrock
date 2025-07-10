---
title: "Steampipe Table: hudsonrock_domain_search"
description: "Query Hudson Rock domain intelligence data with SQL."
---

# Table: hudsonrock_domain_search

The `hudsonrock_domain_search` table allows you to query domain-related cybercrime and infostealer intelligence from the Hudson Rock API.

## Table Usage Guide

The `hudsonrock_domain_search` table provides insights about a domain, including the number of employees, users, third parties, total records, and more, as discovered by Hudson Rock's intelligence platform.

**Important Notes**
- You must provide a `domain` qualifier in the `where` clause for all queries.

## Examples

### Basic domain intelligence

```sql
select
  domain,
  total,
  total_stealers,
  employees,
  users,
  third_parties,
  logo
from
  hudsonrock_domain_search
where
  domain = 'tesla.com';
```

### Get all data for a domain

```sql
select
  *
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```

### Get the logo URL for a domain

```sql
select
  domain,
  logo
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```

## Columns

| Name                | Type    | Description                                      |
|---------------------|---------|--------------------------------------------------|
| domain              | string  | Domain searched.                                 |
| total               | int     | Total records found.                             |
| total_stealers      | int     | Total stealers found.                            |
| employees           | int     | Number of employees.                             |
| users               | int     | Number of users.                                 |
| third_parties       | int     | Number of third parties.                         |
| logo                | string  | Logo URL.                                        |
| data                | json    | Raw data from the API response.                  |
