# Malicious HTA Spyware

<kbd>[UNIT:ANTIVIRUS](https://docs.preludesecurity.com/docs/security-policy#antivirus)</kbd>
<kbd>[ALERT:-](#the-url)</kbd>

This test drops a defanged malicious HTA file. An HTA file is a type of file that contains HTML, JavaScript, and VBScript code that Microsoft's Internet Explorer browser engine can execute. Malicious actors may use HTA files to deliver malware to unsuspecting victims or to execute malicious code on a victim's computer. One common delivery method for malicious HTA files is through email attachments, such as in a spear-phishing campaign. The HTA file may be disguised as a legitimate document, such as an invoice or job application, to trick the victim into opening it. Once the malicious HTA file is executed, it may perform a number of malicious activities, such as downloading and executing additional malware, stealing sensitive data from the victim's computer, or providing remote access to the attacker. This test will monitor if any endpoint defense quarantines the malware.

## How

> Safety: the malware used has been defanged, so even if run, it will immediately exit.

The malware file is extracted from the test and written to a user-owned directory. The test then waits briefly before running a few checks to determine if it was responded to (not just detected) by any defensive products on the endpoint.

Example Output:
```bash
[+] Extracting file for quarantine test
[+] Pausing for 3 seconds to gauge defensive reaction
[-] Malicious file was not caught
```

## Resolution

If this test fails:

* Ensure you have an antivirus program installed and running.
* If using an EDR, make sure the antivirus capability is enabled and turned up, appropriately
