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

### List all URLs (employee and client) for a domain in a single array

```sql+postgres
select
  domain,
  all_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  all_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

### Get the API message and all data for a domain

```sql+postgres
select
  domain,
  message,
  data
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  message,
  data
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```
