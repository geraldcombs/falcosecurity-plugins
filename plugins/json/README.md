# Falcosecurity Json Plugin

This directory contains the json extractor plugin, which can extract values from any json payload. It is used to extract information from json payloads like [k8s_audit](https://falco.org/docs/event-sources/kubernetes-audit/) events or from event payloads generated by source plugins like [cloudtrail](../cloudtrail/README.md), which happen to represent their event payload as json.

## Event Source

The Json plugin is an extractor plugin, and as a result does not have an event source.

## Supported Fields

Here is the current set of supported fields:

| Name | Type | Description |
| ---- | ---- | ----------- |
| json.value | string | Extracts a value from a json-encoded input. Syntax is json.value[<json pointer>], where `<json pointer>` is a [json pointer](https://datatracker.ietf.org/doc/html/rfc6901).
| json.obj | string | The full json message as a text string.
| json.rawtime | string | The time of the event, identical to evt.rawtime.
| jevt.value | string | Alias for json.value, provided for backwards compatibility.
| jevt.obj | string | Alias for json.obj, provided for backwards compatibility.
| jevt.rawtime" | string | Alias for json.rawtime, provided for backwards compatibility.

## Configuration

This plugin does not have any configuration. Any initialization value passed to `plugin_init()` is ignored.

### `falco.yaml` Example

Here is a complete `falco.yaml` snippet showing valid configurations for the dummy plugin:

```yaml
plugins:
  - name: json
    library_path: libjson.so
    init_config: ""
    open_params: ""

# Optional. If not specified the first entry in plugins is used.
load_plugins: [json]
```
