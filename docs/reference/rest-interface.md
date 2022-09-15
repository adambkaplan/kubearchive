# Kubearchive REST Interface

All resources are subresources under `/archive/{cluster}/{namespace}`, which partitions resources by cluster name and namespace.

## `{group}`/`{version}`/`{kind}`

Archived objects use a dynamic path under the resource's group, version, and kind. The group name must be URL-encoded. For example, kubernetes `Job` objects would be archived under `/batch/v1/job`.

| Path | Method | Query Params | Description |
| ---- | ------ | ------------ | ----------- |
| `/` | GET | `label`, `name` | List all archived resources of the given kind. Option to filter by name and labels. Multiple labels can be specified, and these specify either the `key` or `key=value`. |
| `/` | POST | | Create an archive record.
| `/{archive_meta_id}`| GET | | Get the full JSON of an archived record.
| `/{archive_meta_id}` | PUT | | Update the archived record. |
| `/{archive_meta_id}` | DELETE | | Delete the archived record. |
| `/{archive_meta_id}/logs` | GET | | Get all log records associated with an archived record.
| `/{archive_meta_id}/logs` | POST | | Append a log to a resource record. |

## `logs`

Top-level resource for logs.
All logs belong to a namespace.
From there, a log can be associated with one or more resource records.

| Path | Method | Query Params | Description |
| ---- | ------ | ------------ | ----------- |
| `/` | GET | `pod_name`, `container_type`, `container_name` | List all log records for the namespace. Option to filter by pod name, container type, and container name. |
| `/` | POST | | Add a log record for a container. |
| `/{archive_log_id}` | GET | | Fetch the log for the given identifier. This streams the actual contents of the log. |
| `/{archive_log_id}` | POST | | Append the actual container log to a log record. |
| `/{archive_log_id}` | DELETE | | Delete the log record. |
