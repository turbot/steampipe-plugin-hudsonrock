---
title: "Steampipe Table: hudsonrock_urls_by_domain"
description: "Query Hudson Rock for URLs identified by infostealer infections for a given domain with SQL."
---

# Table: hudsonrock_urls_by_domain - Query URLs Identified by Infostealer Infections for a Domain using SQL

The `hudsonrock_urls_by_domain` table allows you to query URLs identified by infostealer infections for a given domain using the Hudson Rock API. This table provides lists of employee, client, and all URLs associated with infostealer infections for a specified domain, as well as additional metadata and API messages.

## Table Usage Guide

The `hudsonrock_urls_by_domain` table provides lists of employee and client URLs, as well as all URLs, associated with infostealer infections for a specified domain. It also includes any API messages and raw data returned by the Hudson Rock API.

**Important Notes**
- You must provide a `domain` qualifier in the `where` clause for all queries.
- This table is not intended for listing all domains, but for targeted intelligence on specific domains.

## Examples

### List all employee and client URLs for a domain

```sql+postgres
select
  domain,
  employees_urls,
  clients_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  employees_urls,
  clients_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

### Unnest employee URLs

```sql+postgres
select
  domain,
  jsonb_array_elements(employees_urls) as employee_url
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  json_each(employees_urls) as employee_url
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

#### Find domains with more than 5 employee URLs

```sql+postgres
select
  domain,
  employees_urls,
  jsonb_array_length(employees_urls) as num_employee_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com'
  and jsonb_array_length(employees_urls) > 5;
```

```sql+sqlite
select
  domain,
  employees_urls,
  json_array_length(employees_urls) AS num_employee_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com'
  and json_array_length(employees_urls) > 5;
```

#### Search for a specific substring in any employee URL

```sql+postgres
select
  domain,
  employees_urls
from
  hudsonrock_urls_by_domain
where
  exists (
    select 1 from jsonb_array_elements_text(employees_urls) as url where url ilike '%login%'
  )
  and domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  employees_urls
from
  hudsonrock_urls_by_domain
where
  exists (
    select 1
    from jsonb_array_elements_text(employees_urls) AS url
    where url ILIKE '%login%'
  )
  and domain = 'hp.com';
```

#### Order by number of employee URLs (descending)

```sql+postgres
select
  domain,
  employees_urls,
  jsonb_array_length(employees_urls) as num_employee_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com'
order by
  num_employee_urls desc;
```

```sql+sqlite
select
  domain,
  employees_urls,
  json_array_length(employees_urls) AS num_employee_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com'
order by
  num_employee_urls desc;
```

