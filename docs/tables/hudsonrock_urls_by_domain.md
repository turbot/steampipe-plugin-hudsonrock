---
title: "Steampipe Table: hudsonrock_urls_by_domain"
description: "Query Hudson Rock for URLs identified by infostealer infections for a given domain with SQL."
---

# Table: hudsonrock_urls_by_domain

The `hudsonrock_urls_by_domain` table allows you to query URLs identified by infostealer infections for a given domain using the Hudson Rock API.

## Table Usage Guide

The `hudsonrock_urls_by_domain` table provides lists of employee and client URLs associated with infostealer infections for a specified domain.

**Important Notes**
- You must provide a `domain` qualifier in the `where` clause for all queries.

## Examples

### List all employee and client URLs for a domain

```sql
select
  domain,
  employees_urls,
  clients_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

### Get the API message and all data for a domain

```sql
select
  domain,
  message,
  data
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```
