# Safety Check

<kbd>[UNIT:HEALTH](https://docs.preludesecurity.com/docs/security-policy#health)</kbd>
<kbd>[ALERT:-](#the-url)</kbd>

This test checks if a long-running Verified Security Test (VST) is stopped properly. A long-running VST that fails to stop properly can cause issues such as resource exhaustion or unexpected results. This test will monitor if the VST is stopped correctly after 10 seconds.

Example Output:
```
[+] Starting test
[+] Sleeping for 15 seconds
[-] VST did not stop properly, returned exit code 101
```

## How

> Safety: this test sleeps for a maximum duration of 15 seconds to simulate a long running VST.

The test waits for 15 seconds before exiting. If the exit code is 101, the VST has stopped properly within 10 seconds, and the test will pass. If the exit code is 102, the VST has not stopped properly, and the test will fail.

## Resolution

If this test fails:

* Contact Prelude Support for assistance
