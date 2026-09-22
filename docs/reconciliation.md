# Session reconciliation

Reconciliation starts after Nakama initializes the adapter and runs immediately, then every I3D_RECONCILE_INTERVAL (default 1m). Each pass has I3D_RECONCILE_TIMEOUT (30s). Startup failures are logged and retried on the next interval. Shutdown cancels the worker and stops its ticker. Set the interval to 0s to disable it.

The worker first captures every storage page and its versions, then scans every provider page for I3D_APPLICATION_ID and optional I3D_FLEET_ID. A failed page, invalid response or repeated cursor aborts the pass without applying that scan. It does not infer absence from an individual page.

Only records with known matching application/fleet ownership and older than I3D_RECONCILE_GRACE_PERIOD (default 2m) can be removed. A record must be absent or no longer ALLOCATED in two consecutive complete scans, with the same storage version. Failed scans reset that confirmation. Legacy records with unknown ownership are retained when absent. Provider-visible allocated records establish scope when refreshed; applications selected with overwriteApplicationId need a separately scoped worker to reconcile their absence.

Refreshes and deletions use the version captured before the scan. Concurrent allocation, admission or update wins a conflict; reconciliation waits for a later pass instead of retrying against the newer record. Newly discovered allocated records are inserted only if absent. Provider refresh preserves locally owned capacity and admissions; a changed provider creation timestamp resets stale ownership. Discovered records without local capacity cannot issue Join admissions until an application establishes capacity.

This is eventual consistency, not proof that a server has terminated. Two complete scans plus the grace period reduce false absence but cannot eliminate a provider outage that returns successful, incomplete results. Tune the grace period to the fleet's observed propagation delay, or disable absence reconciliation when that assumption cannot be met. Multiple Nakama nodes are safe through conditional storage operations; they may perform redundant scans.

Reconciliation only repairs Nakama storage. It never calls the provider restart operation. The public Delete/lifecycle RPC remains the explicit request to restart a game server.

Get also captures the cache version before its provider request; conflicts restart the entire read up to four attempts. Delete captures the version before restarting and never removes a replacement allocation. Nakama 3.41 returns a plain error for conditional-delete conflicts: the adapter confirms a changed/missing version with a new storage read instead of depending on an error string or write-only sentinel.
