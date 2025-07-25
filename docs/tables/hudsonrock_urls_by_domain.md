---
title: "Steampipe Table: hudsonrock_urls_by_domain"
description: "Query Hudson Rock for URLs identified by infostealer infections for a given domain with SQL."
folder: "URL"
---

# Table: hudsonrock_urls_by_domain - Query URLs Identified by Infostealer Infections for a Domain using SQL

The `hudsonrock_urls_by_domain` table allows you to query URLs identified by infostealer infections for a given domain using the Hudson Rock API. This table provides lists of employee, client, and all URLs associated with infostealer infections for a specified domain, as well as additional metadata and API messages.

## Table Usage Guide

The `hudsonrock_urls_by_domain` table provides lists of employee and client URLs, as well as all URLs, associated with infostealer infections for a specified domain. It also includes any API messages and raw data returned by the Hudson Rock API.

**Important Notes**

- You must specify the `domain` in the `where` or join clause (`where domain=`, `join hudsonrock_urls_by_domain s on s.domain=`) in order to query this table.

## Examples

### Basic URL intelligence
Retrieve essential URL intelligence data to understand the scope of compromised web services and applications for a specific domain. This query helps in identifying the types of services that have been accessed by infostealer malware, including both employee and client-facing applications.

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

### Get employee URL analysis
Extract and analyze individual employee URLs to understand which internal services and applications have been compromised. This query helps security teams identify specific applications that may need additional security controls or monitoring to prevent future credential theft.

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

### High exposure domain detection
Identify domains with significant URL exposure by filtering for those with more than 5 compromised employee URLs. This query helps security teams prioritize response efforts by focusing on organizations with the most extensive compromise footprint that may represent higher risk targets.

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

### Authentication service detection
Search for authentication-related URLs within compromised employee URLs to identify login portals and credential management systems that may have been targeted. This query helps in understanding the attack vectors and can guide improvements in authentication security.

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

### URL exposure ranking
Rank domains by the number of compromised employee URLs to identify the most severely affected organizations. This query helps security teams prioritize incident response efforts and can be used to allocate resources based on the severity of URL exposure.

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

### Comprehensive URL compromise assessment
Perform a complete analysis of all URL intelligence data for a compromised domain. This query provides a holistic view of the URL exposure, including both employee and client-facing services, enabling comprehensive incident response and remediation planning.

```sql+postgres
select
  domain,
  employees_urls,
  clients_urls,
  all_urls,
  message,
  jsonb_array_length(employees_urls) as num_employee_urls,
  jsonb_array_length(clients_urls) as num_client_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  employees_urls,
  clients_urls,
  all_urls,
  message,
  json_array_length(employees_urls) as num_employee_urls,
  json_array_length(clients_urls) as num_client_urls
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

### Client URL analysis
Extract and analyze individual client URLs to understand which customer-facing services have been compromised. This query helps security teams identify potential data breach scope and can guide customer notification and remediation efforts.

```sql+postgres
select
  domain,
  jsonb_array_elements(clients_urls) as client_url
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

```sql+sqlite
select
  domain,
  json_each(clients_urls) as client_url
from
  hudsonrock_urls_by_domain
where
  domain = 'hp.com';
```

