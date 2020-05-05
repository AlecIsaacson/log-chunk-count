# synthetics-failover

This utility iterates through all locations for all synthetics monitors and for each test found to be running at location nrSource, it replaces it with location nrTarget.

It expects three arguments:

  * -apikey : A New Relic administrative API key for an account.
  * -nrSource : The name of a New Relic synthetics location.
  * -nrTarget : The name of another New Relic synthetics locations.

It also supports two optional arguments - one for debugging:

  * -verbose=true

and a simulation mode, where the app runs, but no change are actually made:

  * -simulate=true

by default, simulation mode is disabled.

The intent is for this to drive a failover process from one New Relic private location to another, but it can be used for other purposes.

Example:
```
synthetics-failover -apikey=12345 -nrSource=MY_BROKEN_LOCATION -nrTarget=MY_WORKING_LOCATION
```
