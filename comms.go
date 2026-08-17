package epplet

/*

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

extern void goCommsGateway(void* data, char* msg);
static inline void register_comms_bridge() {
    Epplet_register_comms_handler((void (*)(void*, char*))goCommsGateway, NULL);
}

*/
import "C"

import "unsafe"
import "strings"
import "sync"


var (
	commsMu       sync.RWMutex
	commsHandlers = make(map[string]func(args string))
	commsFallback func(rawMsg string)
)

//export goCommsGateway
func goCommsGateway(_ unsafe.Pointer, cMsg *C.char) {
	if cMsg == nil {
		return
	}
	raw := C.GoString(cMsg)

	// Simple whitespace-based tokenizer: "<command> [arguments...]"
	parts := strings.SplitN(strings.TrimSpace(raw), " ", 2)
	if len(parts) == 0 || parts[0] == "" {
		return
	}

	cmd := parts[0]
	args := ""
	if len(parts) > 1 {
		args = parts[1]
	}

	commsMu.RLock()
	handler, exists := commsHandlers[cmd]
	fallback := commsFallback
	commsMu.RUnlock()

	if exists {
		handler(args)
	} else if fallback != nil {
		fallback(raw)
	}
}

var commsInitialized sync.Once

func initComms() {
	commsInitialized.Do(func() {
		C.register_comms_bridge()
	})
}

// HandleCommand registers a callback for a specific command prefix.
func HandleCommand(cmd string, fn func(args string)) {
	initComms()
	commsMu.Lock()
	commsHandlers[cmd] = fn
	commsMu.Unlock()
}

// HandleUnknownCommand registers a fallback for any unhandled raw messages.
func HandleUnknownCommand(fn func(rawMsg string)) {
	initComms()
	commsMu.Lock()
	commsFallback = fn
	commsMu.Unlock()
}
