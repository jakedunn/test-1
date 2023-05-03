# Health Check

<kbd>[UNIT:HEALTH](https://docs.preludesecurity.com/docs/security-policy#health)</kbd>
<kbd>[ALERT:-](#the-url)</kbd>

The Prelude Health Check is a basic test that checks if a device is capable of running future Verified Security Tests.

Endpoint protection should not interfere with this test.

## How

> Safety: no code exists in the test or clean functions of this test

The health check has no business logic but instead just exits with a successful code (100) if it is runnable. 

## Resolution

If this test fails:

* Ensure the device is online and running a probe
