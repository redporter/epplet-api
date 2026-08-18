# Config Subsystem Documentation

This document describes the persistent configuration management system in the `epplet-api` library used to read, write, and manage Epplet settings on disk.

---

## 1. Overview

Enlightenment 16 Epplets feature a key-value configuration engine stored in the user's e16 configuration directory (`~/.e16/epplets/configs/`).

- **Automatic Persistence**: Configurations automatically load during `epplet.Init()` and flush to disk during `epplet.Cleanup()`.
- **Instance Identification**: Multiple instances of the same Epplet maintain isolated configuration sets indexed by instance ID (`GetInstance()`).

---

## 2. Single Key-Value Operations

### Reading Settings

1. **`epplet.QueryConfig`**:
   ```go
   func QueryConfig(key string) string
   ```
   Returns the string value associated with `key`, or `""` if missing.

2. **`epplet.QueryConfigDef`**:
   ```go
   func QueryConfigDef(key, defaultValue string) string
   ```
   Returns the string value associated with `key`, or `defaultValue` if the key is not found.

---

### Writing & Updating Settings

1. **`epplet.ModifyConfig`**:
   ```go
   func ModifyConfig(key, value string)
   ```
   Updates an existing configuration key or adds it if missing.

2. **`epplet.AddConfig`**:
   ```go
   func AddConfig(key, value string)
   ```
   Adds a key-value pair. Faster than `ModifyConfig` when a key is known to be missing.

---

## 3. Multi-Value Array Configurations

For settings consisting of variable-length string arrays (e.g. bookmark lists, monitored network interfaces, server lists):

### `epplet.ModifyMultiConfig`
```go
func ModifyMultiConfig(key string, values []string)
```
- **Description**: Associates an array of string values with a single key.

---

### `epplet.QueryMultiConfig`
```go
func QueryMultiConfig(key string) []string
```
- **Description**: Retrieves the array of string values associated with `key`. Returns `nil` if missing.

---

## 4. File, Instance & Directory Helpers

| Function | Description |
| :--- | :--- |
| `LoadConfig()` | Loads default epplet configuration settings from disk. |
| `LoadConfigFile(filename string)` | Loads configuration data from a specific file path. |
| `SaveConfig()` | Immediately flushes current in-memory configurations to disk. |
| `ClearConfig()` | Deletes all in-memory configuration settings for the epplet. |
| `GetInstance() int` | Returns the instance integer ID of the running epplet. |
| `DataDir() string` | Returns the path to the system epplet data directory (e.g. `/usr/local/share/e16/epplet_data`). |
| `E16UserDir() string` | Returns the path to the user's e16 configuration directory (e.g. `/home/user/.e16`). |

---

## 5. Code Example: Configuration Management

```go
package main

import (
	"fmt"

	"github.com/redporter/epplet-api"
)

func main() {
	epplet.Init("ConfigDemo", "1.0", "Configuration Management Example", 4, 4, false)

	// Load existing configuration (or use defaults)
	username := epplet.QueryConfigDef("username", "Guest")
	refreshRate := epplet.QueryConfigDef("refresh_rate", "5")
	fmt.Printf("Loaded User: %s, Refresh Rate: %s sec\n", username, refreshRate)

	// Read multi-value array config
	servers := epplet.QueryMultiConfig("monitored_hosts")
	if len(servers) == 0 {
		servers = []string{"192.168.1.1", "10.0.0.1"}
		epplet.ModifyMultiConfig("monitored_hosts", servers)
	}
	fmt.Println("Monitored Hosts:", servers)

	// Modify settings
	epplet.ModifyConfig("username", "Alice")
	epplet.ModifyConfig("refresh_rate", "10")

	// Force save to disk (also done automatically on Cleanup)
	epplet.SaveConfig()

	epplet.Show()
	epplet.Loop()
}
```
