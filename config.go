package epplet

/*
#include <X11/Xlib.h>

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

#include <stdlib.h>

static inline void free_string_array(char **arr, int count) {
    if (!arr) return;
    for (int i = 0; i < count; i++) {
        if (arr[i]) free(arr[i]);
    }
    free(arr);
}
*/
import "C"

import "unsafe"

// LoadConfig loads the default epplet configuration file.
func LoadConfig() {
	C.Epplet_load_config()
}

// LoadConfigFile loads configuration data from a specific file path.
func LoadConfigFile(filename string) {
	cFile := C.CString(filename)
	defer C.free(unsafe.Pointer(cFile))
	C.Epplet_load_config_file(cFile)
}

// GetInstance returns the instance ID of this epplet.
func GetInstance() int {
	return int(C.Epplet_get_instance())
}

// QueryConfig returns the configuration value for a key, or empty string if not found.
func QueryConfig(key string) string {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cVal := C.Epplet_query_config(cKey)
	if cVal == nil {
		return ""
	}
	return C.GoString(cVal)
}

// QueryConfigDef returns the configuration value for a key, or defaultValue if not found.
func QueryConfigDef(key, defaultValue string) string {
	cKey := C.CString(key)
	cDef := C.CString(defaultValue)
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cDef))
	cVal := C.Epplet_query_config_def(cKey, cDef)
	if cVal == nil {
		return defaultValue
	}
	return C.GoString(cVal)
}

// ModifyConfig updates or sets a configuration key-value pair.
func ModifyConfig(key, value string) {
	cKey := C.CString(key)
	cVal := C.CString(value)
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cVal))
	C.Epplet_modify_config(cKey, cVal)
}

// AddConfig adds a new configuration key-value pair. Use when the key is known to be missing.
func AddConfig(key, value string) {
	cKey := C.CString(key)
	cVal := C.CString(value)
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cVal))
	C.Epplet_add_config(cKey, cVal)
}

// ModifyMultiConfig sets multiple string values for a configuration key.
func ModifyMultiConfig(key string, values []string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	count := len(values)
	if count == 0 {
		C.Epplet_modify_multi_config(cKey, nil, 0)
		return
	}

	cValues := make([]*C.char, count)
	for i, s := range values {
		cStr := C.CString(s)
		defer C.free(unsafe.Pointer(cStr))
		cValues[i] = cStr
	}

	C.Epplet_modify_multi_config(cKey, (**C.char)(unsafe.Pointer(&cValues[0])), C.int(count))
}

// QueryMultiConfig queries multiple string values for a configuration key.
func QueryMultiConfig(key string) []string {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	var count C.int
	cArr := C.Epplet_query_multi_config(cKey, &count)
	if cArr == nil || count <= 0 {
		return nil
	}

	n := int(count)
	defer C.free_string_array(cArr, count)

	cSlice := unsafe.Slice(cArr, n)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		if cSlice[i] != nil {
			res[i] = C.GoString(cSlice[i])
		}
	}
	return res
}

// SaveConfig flushes and saves current configuration settings to disk.
func SaveConfig() {
	C.Epplet_save_config()
}

// ClearConfig deletes current in-memory configuration settings.
func ClearConfig() {
	C.Epplet_clear_config()
}

// DataDir returns the path to the epplet data directory.
func DataDir() string {
	cStr := C.Epplet_data_dir()
	if cStr == nil {
		return ""
	}
	return C.GoString(cStr)
}

// E16UserDir returns the path to the e16 user configuration directory.
func E16UserDir() string {
	cStr := C.Epplet_e16_user_dir()
	if cStr == nil {
		return ""
	}
	return C.GoString(cStr)
}
