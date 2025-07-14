---
title: "Steampipe Table: hudsonrock_domain_search"
description: "Query Hudson Rock domain intelligence data with SQL."
---

# Table: hudsonrock_domain_search - Query Hudson Rock Domain Intelligence using SQL

The `hudsonrock_domain_search` table allows you to query domain-related cybercrime and infostealer intelligence from the Hudson Rock API. This table provides a comprehensive view of a domain's exposure, including employee and user compromise statistics, stealer malware families, password strength, third-party associations, and more.

## Table Usage Guide

The `hudsonrock_domain_search` table provides insights about a domain, including the number of employees, users, third parties, total records, stealer families, password statistics, and more, as discovered by Hudson Rock's intelligence platform.

**Important Notes**
- You must provide a `domain` qualifier in the `where` clause for all queries.
- This table is not intended for listing all domains, but for targeted intelligence on specific domains.

## Examples

### Basic domain intelligence

```sql+postgres
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

```sql+sqlite
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

### Get the logo URL for a domain

```sql+postgres
select
  domain,
  logo
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  logo
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```

### Get stealer family breakdown for a domain

```sql+postgres
select
  domain,
  stealer_families
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  stealer_families
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```

### Get password strength statistics for employees and users

```sql+postgres
select
  domain,
  employee_passwords,
  user_passwords
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  employee_passwords,
  user_passwords
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```

### List third-party domains associated with a domain

```sql+postgres
select
  domain,
  jsonb_array_elements(third_party_domains) as third_party_domain
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  json_each(third_party_domains) as third_party_domain
from
  hudsonrock_domain_search
where
  domain = 'hp.com';
```
