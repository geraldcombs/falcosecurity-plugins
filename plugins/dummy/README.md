# Falcosecurity Dummy Plugin

This directory contains the dummy plugin, which is an example plugin written in Go. It's referenced in the [developer's guide](https://falco.org/docs/plugins/developers_guide/) to walk through the implementation of a source plugin.

It generates synthetic events and doesn't serve any purpose other than for documentation.

It is a companion to the `dummy_c` plugin which is written in C++ but supports the same fields, event source, and configuration.

## Event Source

The event source for dummy events is `dummy`.

## Supported Fields

Here is the current set of supported fields:

| Name | Type | Description |
| ---- | ---- | ----------- |
| dummy.divisible | uint64 | Return 1 if the value is divisible by the provided divisor, 0 otherwise.
| dummy.value | uint64 | The sample value in the event. 
| dummy.strvalue | string | The sample value in the event, as a string.

## Configuration

### Plugin Initialization

The format of the initialization string is a json object. Here's an example:

```json
{"jitter": 10}
```

The json object has the following properties:

* `jitter`: Controls the random value that is added to each event returned in next().

The init string can be the empty string, which is treated identically to `{}`.

### Plugin Open Params

The format of the open params string is a json object. Here's an example:

```json
{"start": 1, "maxEvents": 100}
```

The json object has the following properties:
* `start`: denotes the initial value of the sample
* `maxEvents`: denotes the number of events to return before returning EOF.

The open params string can be the empty string, which is treated identically to `{}`.

### `falco.yaml` Example

Here is a complete `falco.yaml` snippet showing valid configurations for the dummy plugin:

```yaml
plugins:
  - name: dummy
    library_path: libdummy.so
    init_config: '{"jitter": 10}'
    open_params: '{"start": 1, "maxEvents": 100}'

# Optional. If not specified the first entry in plugins is used.
load_plugins: [dummy]
```
