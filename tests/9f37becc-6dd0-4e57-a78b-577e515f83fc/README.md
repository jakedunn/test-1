# IcedID

<kbd>[UNIT:ANTIVIRUS](https://docs.preludesecurity.com/docs/security-policy#antivirus)</kbd>
<kbd>[ALERT:-](#the-url)</kbd>

This test drops a defanged MSI file which contains a defanged IcedID dropper used for command and control and other post-exploitation activities. MSI files are a type of installer file used to distribute software applications. Malicious .msi files can be used to install malware, such as trojans or ransomware, on the victim's computer. First discovered in 2017, IcedID is classified as a banking trojan and remote access trojan (RAT). Having capabilities comparable to other sophisticated banking Trojans such as Zeus, Gozi, and Dridex. IcedID is a second-stage malware reliant on other first-stage malware, such as Emotet, to gain initial access and deploy it. In addition to stealing victims' financial information, IcedID often serves as a dropper for other second-stage malware, including ransomware, and has advanced capabilities to move laterally through a network. This test will monitor if any endpoint defense quarantines the malware.

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
