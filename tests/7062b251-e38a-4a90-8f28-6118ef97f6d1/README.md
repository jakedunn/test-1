# Is an endpoint control running?

Endpoint control solutions are designed to prevent malicious activity on an endpoint. Ensuring that your endpoint control mechanism is running is a crucial first step to securing your enterprise network. The following test attempts to do this by checking for specific processes that are commonly associated with endpoint control solutions.

Example output: 
```
[+] Starting test
[+] com.crowdstrike.falcon.Agent is running
[+] Completed with code: 100
```

An endpoint security control should be running.

## How

> Safety: this VST simply checks for processes and doesn't interfere with any of them.

The test works by first defining a list of endpoint control solution processes for Windows, macOS, and Linux. It then checks if any of these processes are running on the endpoint by running a command specific to the endpoint's operating system. If any of the specified processes are found to be running, the test will stop and return a 100 exit code. 