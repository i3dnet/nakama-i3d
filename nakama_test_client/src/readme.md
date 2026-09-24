# Local Nakama clients

Use sh scripts/smoke.sh from the repository root for the bounded, automated two-client check. It opens real Nakama sockets and tests allocation notifications, lifecycle authorization, native storage concurrency, pagination and recovery.

For the interactive development client, start the development Compose stack and run go run ./cmd/main from this directory with Go 1.27.1. NAKAMA_URL defaults to http://127.0.0.1:7350. Start two clients for one match, or use the clients Compose profile with --scale nakama-test-client=2. Each connection submits one ticket.

The local fixtures use defaultkey, local-test and test-token. They are test credentials. The player client must not hold One API credentials or the headless server's Nakama HTTP key.

SessionId in the example notification is empty because the adapter does not issue provider reservation tokens. Connect using IpAddress and Port and implement player authentication in your game transport.
