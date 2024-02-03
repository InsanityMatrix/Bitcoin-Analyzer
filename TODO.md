Current Process

Analyzer will have a set monitor threshold,

A Monitored TX will be monitored for the next 2016 blocks (2 Weeks)
If UTXO from monitored TX is sent, gets added to chain and timer restarts
Max of 3 chains until it hits another exchange address (unless otherwise configured) or monitored transaction gets dropped

If a Monitored TX hits all criteria, the Address gets Saved along with transactions


Exchange addresses are fetched using:

You know a transaction came from exchange if:


