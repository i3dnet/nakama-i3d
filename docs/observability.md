# Observability

Nakama metrics expose i3d_allocation_total / i3d_allocation_duration for accepted asynchronous allocations, and i3d_reconciliation_total / i3d_reconciliation_duration for reconciliation passes. The only tag is result: success, error, timeout or canceled. An allocation records one terminal outcome before its callback, including when no callback was supplied. Validation failures before acceptance are returned synchronously and are not counted as allocations.

Alert on repeated reconciliation errors or timeouts and rising allocation error rates. A process-local callback is not durable; a crash can prevent a terminal metric or notification. Healthcheck is handler liveness, not provider readiness.

The OpenAPI client's existing OpenTelemetry instrumentation is retained with bounded route templates and method/status labels. Debug logs and spans exclude authorization headers, response headers, request/response bodies, instance IDs and filter queries. HTTP errors retain a status-only span description. Keep these customizations when regenerating client.go; do not restore generated raw request/response dumps. Transport errors returned to callers may still include the endpoint, so configure endpoints without embedded credentials.
