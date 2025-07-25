---
title: "Steampipe Table: hudsonrock_username_search"
description: "Query Hudson Rock infostealer and credential data by username with SQL."
folder: "User"
---

# Table: hudsonrock_username_search - Query Hudson Rock Username Intelligence using SQL

The `hudsonrock_username_search` table allows you to query compromised credentials and infostealer data by username using the Hudson Rock API. This table provides a detailed view of infostealer infections, password and login exposures, stealer malware families, and more for a given username.

## Table Usage Guide

The `hudsonrock_username_search` table provides insights about compromised credentials, infostealer malware, and related data for a given username, including compromise date, stealer family, computer and OS details, IP address, and exposed passwords and logins.

**Important Notes**

- You must specify the `username` in the `where` or join clause (`where username=`, `join hudsonrock_username_search s on s.username=`) in order to query this table.

## Examples

### Basic username intelligence
Retrieve essential username intelligence data to understand the overall compromise status and exposure details for a specific username. This query helps in identifying the scope of credential theft, infostealer infections, and the types of sensitive information that may have been exposed from accounts associated with this username.

```sql+postgres
select
  username,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

### High value credential detection
Identify usernames with significant credential exposure by filtering for those with more than 3 compromised passwords. This query helps security teams prioritize response efforts by focusing on accounts with the most severely compromised credentials that may represent higher risk targets.

```sql+postgres
select
  username,
  top_passwords,
  jsonb_array_length(top_passwords) as num_passwords
from
  hudsonrock_username_search
where
  username = 'johndoe'
  and jsonb_array_length(top_passwords) > 3;
```

```sql+sqlite
select
  username,
  top_passwords,
  json_array_length(top_passwords) as num_passwords
from
  hudsonrock_username_search
where
  username = 'johndoe'
  and json_array_length(top_passwords) > 3;
```

### Antivirus software analysis
Search for specific antivirus software installations across compromised usernames to understand the effectiveness of endpoint protection solutions. This query helps in identifying patterns in security tool deployments and can guide improvements in antivirus coverage and configuration.

```sql+postgres
select
  username,
  antiviruses
from
  hudsonrock_username_search
where
  username = 'johndoe'
  and antiviruses::text ilike '%Kaspersky%';
```

```sql+sqlite
select
  username,
  antiviruses
from
  hudsonrock_username_search,
  json_each(hudsonrock_username_search.antiviruses)
where
  username = 'johndoe'
  and lower(json_each.value) = 'kaspersky';
```

### Credential exposure ranking
Rank usernames by the number of compromised passwords to identify the most severely affected accounts. This query helps security teams prioritize incident response efforts and can be used to allocate resources based on the severity of credential exposure.

```sql+postgres
select
  username,
  top_passwords,
  jsonb_array_length(top_passwords) as num_passwords
from
  hudsonrock_username_search
where
  username = 'johndoe'
order by
  num_passwords desc;
```

```sql+sqlite
select
  username,
  top_passwords,
  json_array_length(top_passwords) as num_passwords
from
  hudsonrock_username_search
where
  username = 'johndoe'
order by num_passwords desc;
```

### Comprehensive username compromise assessment
Perform a complete analysis of all available intelligence data for a compromised username. This query provides a holistic view of the compromise, including technical details, exposed credentials, system information, and threat context, enabling comprehensive incident response and remediation planning.

```sql+postgres
select
  username,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path,
  stealer_family,
  computer_name,
  os,
  ip_address,
  country,
  city,
  isp,
  compromise_date
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

```sql+sqlite
select
  username,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path,
  stealer_family,
  computer_name,
  os,
  ip_address,
  country,
  city,
  isp,
  compromise_date
from
  hudsonrock_username_search
where
  username = 'johndoe';
```

### Recent compromise detection
Identify recently compromised usernames to prioritize incident response efforts. This query helps security teams focus on the most recent threats and can be used to trigger immediate response procedures for newly discovered compromises.

```sql+postgres
select
  username,
  compromise_date,
  stealer_family,
  message,
  top_passwords,
  top_logins
from
  hudsonrock_username_search
where
  compromise_date > current_date - interval '30 days'
order by
  compromise_date desc;
```

```sql+sqlite
select
  username,
  compromise_date,
  stealer_family,
  message,
  top_passwords,
  top_logins
from
  hudsonrock_username_search
where
  compromise_date > date('now', '-30 days')
order by
  compromise_date desc;
```

### Stealer family analysis
Analyze the distribution of infostealer malware families across compromised usernames to understand the threat landscape. This query helps in identifying the most prevalent malware families and can guide security investments in specific threat detection and prevention capabilities.

```sql+postgres
select
  stealer_family,
  count(*) as compromise_count,
  count(distinct username) as unique_usernames
from
  hudsonrock_username_search
where
  stealer_family is not null
group by
  stealer_family
order by
  compromise_count desc;
```

```sql+sqlite
select
  stealer_family,
  count(*) as compromise_count,
  count(distinct username) as unique_usernames
from
  hudsonrock_username_search
where
  stealer_family is not null
group by
  stealer_family
order by
  compromise_count desc;
```
