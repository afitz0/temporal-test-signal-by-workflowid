# Test

A repro repo for testing setting a custom Workflow ID in a test environment using

```go
env.SetStartWorkflowOptions(client.StartWorkflowOptions{
    ID: "custom-workflow-id",
})
```

and then trying to `SignalWorkflowByID` using that ID.

Expected: `go test` passes.
Actual: `go test` fails with "Received unexpected error: Workflow custom-workflow-id not exists"
