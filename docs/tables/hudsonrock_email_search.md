---
title: "Steampipe Table: hudsonrock_email_search"
description: "Query Hudson Rock infostealer and credential data by email with SQL."
folder: "Email"
---

# Table: hudsonrock_email_search - Query Hudson Rock Email Intelligence using SQL

The `hudsonrock_email_search` table allows you to query compromised credentials and infostealer data by email using the Hudson Rock API. This table provides a detailed view of infostealer infections, password and login exposures, stealer malware families, and more for a given email address.

## Table Usage Guide

The `hudsonrock_email_search` table provides insights about compromised credentials, infostealer malware, and related data for a given email address, including compromise date, stealer family, computer and OS details, IP address, and exposed passwords and logins.

**Important Notes**

- You must specify the `email` in the `where` or join clause (`where email=`, `join hudsonrock_email_search s on s.email=`) in order to query this table.

## Examples

### Basic email intelligence
Retrieve essential email intelligence data to understand the overall compromise status and exposure details for a specific email address. This query helps in identifying the scope of credential theft, infostealer infections, and the types of sensitive information that may have been exposed.

```sql+postgres
select
  email,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
  message,
  top_passwords,
  top_logins,
  antiviruses,
  malware_path
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Get top passwords and logins from the first stealer
Extract the most commonly used passwords and login credentials that were stolen from the infected device. This query helps security teams understand the types of credentials that were compromised and can guide password policy enforcement and credential rotation strategies.

```sql+postgres
select
  email,
  top_passwords,
  top_logins
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
  top_passwords,
  top_logins
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Get antiviruses and malware path from the first stealer
Analyze the antivirus software installed on the compromised device and the specific malware infection path. This query helps in understanding the security posture of the infected system and can guide improvements in endpoint protection and malware detection capabilities.

```sql+postgres
select
  email,
  antiviruses,
  malware_path
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
  antiviruses,
  malware_path
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Comprehensive email compromise assessment
Perform a complete analysis of all available intelligence data for a compromised email address. This query provides a holistic view of the compromise, including technical details, exposed credentials, system information, and threat context, enabling comprehensive incident response and remediation planning.

```sql+postgres
select
  email,
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
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
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
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Geographic and network intelligence
Extract geographic and network information about the compromised device to understand the attack context and potential threat actor location. This query helps in threat intelligence analysis and can assist in identifying patterns across multiple compromises.

```sql+postgres
select
  email,
  ip_address,
  country,
  city,
  isp,
  computer_name,
  os
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

```sql+sqlite
select
  email,
  ip_address,
  country,
  city,
  isp,
  computer_name,
  os
from
  hudsonrock_email_search
where
  email = 'user@example.com';
```

### Recent compromise detection
Identify recently compromised email addresses to prioritize incident response efforts. This query helps security teams focus on the most recent threats and can be used to trigger immediate response procedures for newly discovered compromises.

```sql+postgres
select
  email,
  compromise_date,
  stealer_family,
  message,
  top_passwords,
  top_logins
from
  hudsonrock_email_search
where
  compromise_date > current_date - interval '30 days'
order by
  compromise_date desc;
```

```sql+sqlite
select
  email,
  compromise_date,
  stealer_family,
  message,
  top_passwords,
  top_logins
from
  hudsonrock_email_search
where
  compromise_date > date('now', '-30 days')
order by
  compromise_date desc;
```

### Get stealer family details
Analyze the distribution of infostealer malware families across compromised email addresses to understand the threat landscape. This query helps in identifying the most prevalent malware families and can guide security investments in specific threat detection and prevention capabilities.

```sql+postgres
select
  stealer_family,
  count(*) as compromise_count,
  count(distinct email) as unique_emails
from
  hudsonrock_email_search
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
  count(distinct email) as unique_emails
from
  hudsonrock_email_search
where
  stealer_family is not null
group by
  stealer_family
order by
  compromise_count desc;
```
