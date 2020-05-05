# log-chunk-count

This utility queries New Relic Insights for a count of log events with chunks.

It expects three arguments:

  * -apikey : A New Relic Insights query API key for an account.
  * -account : A New Relic account ID.
  * -chunks : The maximum number of chunks to query for.

It also supports an optional argument for debugging:

  * -verbose=true

This utility runs the following Insights query:

  SELECT count(\*) from Log where 'message-NN' is not null

It iterates from message-01 to message-NN where NN is equal to the value of the *chunks* command line argument.

The output is a CSV of the format chunk number, count.

This utility can be used to determine how many New Relic log messages are being chunked into separate fields due to exceeding the 4K limit.
