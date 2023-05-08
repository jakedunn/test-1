# Common Ransomware

<kbd>[UNIT:NGAV](https://docs.preludesecurity.com/docs/security-policy#ngav)</kbd>
<kbd>[ALERT:-](#the-url)</kbd>

This test focuses on ransomware behaviors which are common across multiple strains of malware: creation of an encryption key and rapid pace encryption of local files using system libraries. Multiple types of ransomware attacks are prevelant nowadays. From run-of-the-mill brute-force encryption of all files, to targeted file selection to no-encryption ransomware - the defense has to be on the lookout for each. Recent ransomware attacks have targeted multiple industries including Healthcare and Education.

## How

> Safety: the test creates the files to encrypt, it doesn't use any existing ones.

This test creates 3 files in the local user's home directory, each with a different size and file type. Then, the test generates a new encryption key and attempts to encrypt each of the new files with the key. If it makes it through all steps without being stopped, the test reports the device as being unable to stop a common ransomware attack.

Example Output:
```bash
[*] Generating sample files...
[+] Generated a new encryption key
[+] Starting encryption process: one.txt
[+] Starting encryption process: two.xlsx
[+] Starting encryption process: three.pdf
```

## Resolution

If this test fails:

* Ensure you have an antivirus program installed and running.
* If using an EDR, make sure the antivirus capability is enabled and turned up, appropriately
