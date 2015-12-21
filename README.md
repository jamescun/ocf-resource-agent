OCF Resource Agent
==================

This package implements the [Resource Agent](https://raw.githubusercontent.com/ClusterLabs/resource-agents/master/doc/dev-guides/ra-dev-guide.txt) interface to Pacemaker for a Go program.

Advantages of using Go instead of (recommended) bash for resource agents include:

  - Faster startup and lower resource consumption
    - Useful for `monitor` operation which is invoked many times a minute
  - Static binary can use client libraries instead of shelling out to resources manager
  - Testability

Todo
----

  - Generate Metadata from configuration
