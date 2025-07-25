---
title: "Steampipe Table: hudsonrock_domain_search"
description: "Query Hudson Rock domain intelligence data with SQL."
folder: "Domain"
---

# Table: hudsonrock_domain_search - Query Hudson Rock Domain Intelligence using SQL

The `hudsonrock_domain_search` table allows you to query domain-related cybercrime and infostealer intelligence from the Hudson Rock API. This table provides a comprehensive view of a domain's exposure, including employee and user compromise statistics, stealer malware families, password strength, third-party associations, and more.

## Table Usage Guide

The `hudsonrock_domain_search` table provides insights about a domain, including the number of employees, users, third parties, total records, stealer families, password statistics, and more, as discovered by Hudson Rock's intelligence platform.

**Important Notes**

- You must specify the `domain` in the `where` or join clause (`where domain=`, `join hudsonrock_domain_search s on s.domain=`) in order to query this table.

## Examples

### Basic domain intelligence
Retrieve essential domain intelligence metrics to understand the overall exposure and compromise statistics for a specific domain. This query helps in identifying the scale of potential security issues, including the number of affected employees, users, and third parties, as well as the total volume of compromised records.

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

### Get domain logo and branding information
Extract the official logo URL for a domain to verify brand authenticity and support visual identification in reports or dashboards. This query is useful for creating branded security reports or integrating domain intelligence into existing security tools.

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

### Get stealer malware family details
Analyze the types of infostealer malware families that have compromised credentials for a specific domain. This query helps security teams understand the threat landscape and prioritize remediation efforts based on the sophistication and capabilities of the malware families involved.

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

### Get password security assessment
Evaluate password strength statistics for both employees and users associated with a domain to identify potential security weaknesses. This query helps in understanding the overall password hygiene and can guide security awareness training priorities and password policy improvements.

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

### Get third-Party domain exposure
Identify and analyze all third-party domains associated with a primary domain to understand the extended attack surface and supply chain risks. This query helps in mapping the broader ecosystem of connected services and identifying potential lateral movement opportunities for attackers.

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

### Get comprehensive domain risk assessment
Perform a complete risk assessment by combining multiple intelligence factors for a domain. This query provides a holistic view of the domain's security posture, including exposure metrics, threat actor activity, and potential attack vectors, enabling informed decision-making for security investments and incident response planning.

```sql+postgres
select
  domain,
  total,
  total_stealers,
  employees,
  users,
  third_parties,
  stealer_families,
  employee_passwords,
  user_passwords,
  logo,
  third_party_domains
from
  hudsonrock_domain_search
where
  domain = 'microsoft.com';
```

```sql+sqlite
select
  domain,
  total,
  total_stealers,
  employees,
  users,
  third_parties,
  stealer_families,
  employee_passwords,
  user_passwords,
  logo,
  third_party_domains
from
  hudsonrock_domain_search
where
  domain = 'microsoft.com';
```

### High risk domain identification
Identify domains with significant exposure by filtering for high compromise counts. This query helps security teams prioritize their response efforts by focusing on domains with the most severe exposure levels, enabling efficient resource allocation for incident response and remediation activities.

```sql+postgres
select
  domain,
  total,
  employees,
  users,
  total_stealers
from
  hudsonrock_domain_search
where
  total > 1000
  and employees > 100
order by
  total desc;
```

```sql+sqlite
select
  domain,
  total,
  employees,
  users,
  total_stealers
from
  hudsonrock_domain_search
where
  total > 1000
  and employees > 100
order by
  total desc;
```
