---
title: "Steampipe Table: hudsonrock_ip_search"
description: "Query Hudson Rock infostealer and credential data by IP address with SQL."
folder: "IP"
---

# Table: hudsonrock_ip_search - Query Hudson Rock IP Intelligence using SQL

The `hudsonrock_ip_search` table allows you to query info-stealer and credential data by IP address using the Hudson Rock API. This table provides a detailed view of infostealer infections, password and login exposures, stealer malware details, and more for a given IP address.

## Table Usage Guide

The `hudsonrock_ip_search` table provides insights about compromised credentials, infostealer malware, and related data for a given IP address, including compromise date, computer and OS details, malware path, antivirus products, and exposed passwords and logins.

**Important Notes**

- You must specify the `ip` in the `where` or join clause (`where ip=`, `join hudsonrock_ip_search s on s.ip=`) in order to query this table.

## Examples

### Basic IP intelligence
Retrieve essential IP intelligence data to understand the overall compromise status and exposure details for a specific IP address. This query helps in identifying the scope of credential theft, infostealer infections, and the types of sensitive information that may have been exposed from devices associated with this IP.

```sql+postgres
select
  ip,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  ip,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

### Get detailed stealer infection analysis
Extract comprehensive details about each infostealer infection associated with an IP address, including compromise timeline, system information, and exposed credentials. This query helps security teams understand the full scope of infections and can guide targeted remediation efforts for specific compromised systems.

```sql+postgres
select
  ip,
  date_compromised,
  computer_name,
  operating_system,
  malware_path,
  antiviruses,
  top_passwords,
  top_logins
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  ip,
  date_compromised,
  computer_name,
  operating_system,
  malware_path,
  antiviruses,
  top_passwords,
  top_logins
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

### High value credential detection
Identify IP addresses with significant credential exposure by filtering for those with more than 3 compromised passwords. This query helps security teams prioritize response efforts by focusing on the most severely compromised systems that may represent higher risk targets.

```sql+postgres
select
  ip,
  top_passwords,
  jsonb_array_length(top_passwords) as num_passwords
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8'
  and jsonb_array_length(top_passwords) > 3;
```

```sql+sqlite
select
  ip,
  top_passwords,
  json_array_length(top_passwords) AS num_passwords
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8'
  and json_array_length(top_passwords) > 3;
```

### Antivirus software analysis
Search for specific antivirus software installations across compromised IP addresses to understand the effectiveness of endpoint protection solutions. This query helps in identifying patterns in security tool deployments and can guide improvements in antivirus coverage and configuration.

```sql+postgres
select
  ip,
  antiviruses
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8'
  and antiviruses::text ilike '%Kaspersky%';
```

```sql+sqlite
select
  ip,
  antiviruses
from
  hudsonrock_ip_search,
  json_each(hudsonrock_ip_search.antiviruses)
where
  ip = '8.8.8.8'
  and lower(json_each.value) = 'kaspersky';
```

### Credential exposure ranking
Rank IP addresses by the number of compromised passwords to identify the most severely affected systems. This query helps security teams prioritize incident response efforts and can be used to allocate resources based on the severity of credential exposure.

```sql+postgres
select
  ip,
  top_passwords,
  json_array_length(top_passwords) AS num_passwords
from
  hudsonrock_ip_search
order by
  num_passwords desc;
```

```sql+sqlite
select
  ip,
  top_passwords,
  json_array_length(top_passwords) AS num_passwords
from
  hudsonrock_ip_search
order by
  num_passwords desc;
```

### Comprehensive IP compromise assessment
Perform a complete analysis of all available intelligence data for a compromised IP address. This query provides a holistic view of the compromise, including technical details, exposed credentials, system information, and threat context, enabling comprehensive incident response and remediation planning.

```sql+postgres
select
  ip,
  message,
  date_compromised,
  computer_name,
  operating_system,
  malware_path,
  antiviruses,
  top_passwords,
  top_logins,
  stealer_family,
  country,
  city,
  isp
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  ip,
  message,
  date_compromised,
  computer_name,
  operating_system,
  malware_path,
  antiviruses,
  top_passwords,
  top_logins,
  stealer_family,
  country,
  city,
  isp
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

### Recent compromise detection
Identify recently compromised IP addresses to prioritize incident response efforts. This query helps security teams focus on the most recent threats and can be used to trigger immediate response procedures for newly discovered compromises.

```sql+postgres
select
  ip,
  date_compromised,
  stealer_family,
  message,
  top_passwords,
  top_logins
from
  hudsonrock_ip_search
where
  date_compromised > current_date - interval '30 days'
order by
  date_compromised desc;
```

```sql+sqlite
select
  ip,
  date_compromised,
  stealer_family,
  message,
  top_passwords,
  top_logins
from
  hudsonrock_ip_search
where
  date_compromised > date('now', '-30 days')
order by
  date_compromised desc;
```

### Geographic threat intelligence
Extract geographic and network information about compromised IP addresses to understand the attack context and potential threat actor locations. This query helps in threat intelligence analysis and can assist in identifying patterns across multiple compromises.

```sql+postgres
select
  ip,
  country,
  city,
  isp,
  computer_name,
  operating_system,
  date_compromised
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```

```sql+sqlite
select
  ip,
  country,
  city,
  isp,
  computer_name,
  operating_system,
  date_compromised
from
  hudsonrock_ip_search
where
  ip = '8.8.8.8';
```
