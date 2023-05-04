# CVE-2019-14287

<kbd>[UNIT:NGAV](https://docs.preludesecurity.com/docs/security-policy#ngav)</kbd>
<kbd>[ALERT:CVE-2019-14287](https://nvd.nist.gov/vuln/detail/CVE-2019-14287)</kbd>

This VST tests whether a host is vulnerable to CVE-2019-14287. CVE-2019-14287 (applicable to Linux and Darwin systems) is caused by a flaw in the sudo command that allows a non-root user to bypass restrictions and run commands as root by specifying the user ID -1. Non-root users being able to run unauthorized commands as root can pose a significant security risk, potentially leading to compromise of the entire system.

## How

> Safety: as a proof-of-concept, this VST utilizes the identified vulnerability and only executes the "whoami" command

This test involves running the "whoami" command as a non-root user, which outputs the user ID of the user who runs it and can be used to determine if the machine is vulnerable to the mapped CVE. This test is safe because running this command makes no lasting change (it only outputs a text value), and no other commands are run.

Example Output:
```bash
[+] Starting test
[+] Testing CVE-2019-14287
[-] Test failed (unauthorized access to root was gained)
```

## Resolution

If this test fails:

* Ensure your system is up-to-date with the latest patches
* Ensure you have an antivirus program installed and running.
* If using an EDR, make sure the antivirus capability is enabled and turned up, appropriately
