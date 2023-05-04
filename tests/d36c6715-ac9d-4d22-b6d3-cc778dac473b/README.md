# CVE-2016-0099

<kbd>[UNIT:NGAV](https://docs.preludesecurity.com/docs/security-policy#ngav)</kbd>
<kbd>[ALERT:CVE-2016-0099](https://nvd.nist.gov/vuln/detail/cve-2016-0099)</kbd>

This VST tests whether a host is vulnerable to CVE-2016-0099. CVE-2016-0099 (applicable to Windows systems) exploits a vulnerability in the Windows Secondary Logon Service that allows a non-administrative user to run arbitrary code with SYSTEM privileges. Non-root users being able to run unauthorized commands as root can pose a significant security risk, potentially leading to compromise of the entire system.

## How

> Safety: as a proof-of-concept, this VST utilizes the identified vulnerabilities and only executes the "whoami" command

This test involves running the "whoami" command as a non-root user, which outputs the user ID of the user who runs it and can be used to determine if the machine is vulnerable to the mapped CVE. This test is safe because running this command makes no lasting change (it only outputs a text value), and no other commands are run.

Example Output:
```bash
[+] Starting test
[+] Testing CVE-2016-0099
[+] The test was prevented or machine not vulnerable
```

## Resolution

If this test fails:

* Ensure your system is up-to-date with the latest patches
* Ensure you have an antivirus program installed and running.
* If using an EDR, make sure the antivirus capability is enabled and turned up, appropriately
