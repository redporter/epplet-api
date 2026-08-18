# Epplet C API vs Go Bindings (`epplet-api`) Mapping

This document provides a comprehensive inventory of all data structures, functions, and symbols declared in the C Epplet API header ([epplet.h](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.h)) and indicates whether and how each element is bridged in the Go wrapper ([epplet-api](file:///home/nemanja/projects/123-Go/dev/epplet-api)).

---

## Summary Matrix

| Category | Total C Symbols | Bridged to Go | Bridging Status |
| :--- | :---: | :---: | :---: |
| **Data Structures & Enums** | 5 | 3 | 60% |
| **Core Lifecycle & Windowing** | 7 | 7 | 100% |
| **Window Creation & Management** | 9 | 9 | 100% |
| **IPC & Communications** | 4 | 4 | 100% |
| **Imageclass & Textclass** | 5 | 5 | 100% |
| **Event & Callback Handlers** | 13 | 13 | 100% |
| **Timers** | 4 | 3 | 75% |
| **Gadget Operations & Constructors** | 35 | 33 | 94% |
| **2D Drawing & RGB Buffers** | 11 | 11 | 100% |
| **OpenGL / GLX Contexts** | 4 | 4 | 100% |
| **Command Execution & Dialogs** | 10 | 10 | 100% |
| **Config File System** | 13 | 13 | 100% |

---

## Detailed Symbol Mapping

### 1. Data Structures & Primitives

| C Type / Symbol in `epplet.h` | Go Type / Symbol in `epplet-api` | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `typedef void *Epplet_gadget;` | `Gadget` interface, `baseGadget`, `*Button`, `*Textbox`, `*Slider`, `*Popup`, etc. | **Bridged** | [gadget.go:L130-L225](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L130-L225) |
| `typedef enum gad_type GadType;` | `type GadgetType int` (`GadgetButton`, `GadgetTextbox`, etc.) | **Bridged** | [gadget.go:L73-L88](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L73-L88) |
| `typedef struct _rgb_buf *RGB_buf;` | `type RGBBuf struct` | **Bridged** | [drawing.go:L114](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L114) |
| `typedef struct _configitem ConfigItem;` | *None* | Unbridged | Config items not exposed as struct |
| `typedef struct _etimer ETimer;` | *Internal Go tracking map* | Indirect | Managed via `Timer()` tracking map in Go |

---

### 2. Core Lifecycle & Main Windowing

| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_Init(...)` | `Init(name, version, info string, w, h int, vertical bool)` | **Bridged** | [epplet.go:L32](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L32) |
| `Epplet_cleanup()` | `Cleanup()` | **Bridged** | [epplet.go:L51](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L51) |
| `Epplet_show()` | `Show()` | **Bridged** | [epplet.go:L55](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L55) |
| `Epplet_remember()` | `Remember()` | **Bridged** | [epplet.go:L59](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L59) |
| `Epplet_unremember()` | `Unremember()` | **Bridged** | [epplet.go:L63](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L63) |
| `Epplet_Loop()` | `Loop()` | **Bridged** | [epplet.go:L67](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L67) |
| `Epplet_get_main_window()` | `GetMainWindow() Window` | **Bridged** | [window.go:L39](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L39) |
| `Epplet_get_display()` | `GetDisplay() *Display` | **Bridged** | [epplet.go:L80](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L80) |

---

### 3. Auxiliary Windows & Context Stack

| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_create_window` | `CreateWindow(w, h int, title string, vert bool) Window` | **Bridged** | [window.go:L34](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L34) |
| `Epplet_create_window_borderless` | `CreateWindowBorderless(w, h int, title string, vert bool) Window` | **Bridged** | [window.go:L48](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L48) |
| `Epplet_create_window_config` | `CreateWindowConfig(w, h int, title string, ok, apply, cancel Cb) Window` | **Bridged** | [window.go:L62](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L62) |
| `Epplet_window_push_context` | `WindowPushContext(win Window)` / `Window.PushContext()` | **Bridged** | [window.go:L91](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L91) |
| `Epplet_window_pop_context` | `WindowPopContext() Window` | **Bridged** | [window.go:L96](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L96) |
| `Epplet_window_show` | `Window.Show()` | **Bridged** | [window.go:L106](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L106) |
| `Epplet_window_hide` | `Window.Hide()` | **Bridged** | [window.go:L111](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L111) |
| `Epplet_window_destroy` | `Window.Destroy()` | **Bridged** | [window.go:L116](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L116) |
| `Epplet_clear_window` | `Window.Clear()` | **Bridged** | [window.go:L121](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L121) |

---

### 4. IPC & Communications

| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_send_ipc(char *s)` | `SendIPC(msg string)` | **Bridged** | [epplet.go:L91](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L91) |
| `Epplet_wait_for_ipc()` | `BlockForIPC() string` / `WaitForIPCAsync()` | **Bridged** | [epplet.go:L99](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L99) |
| `Epplet_register_comms_handler(...)` | `HandleCommand(...)`, `HandleUnknownCommand(...)` | **Bridged** | [comms.go:L68-L82](file:///home/nemanja/projects/123-Go/dev/epplet-api/comms.go#L68-L82) |

---

### 5. Imageclass & Textclass Rendering

| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_imageclass_apply(...)` | `Window.ImageclassApply(iclass, state string)` | **Bridged** | [window.go:L127](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L127) |
| `Epplet_imageclass_paste(...)` | `Window.ImageclassPaste(iclass, state string, x, y, h, w int)` | **Bridged** | [window.go:L137](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L137) |
| `Epplet_imageclass_get_pixmaps(...)` | `ImageclassGetPixmaps(...) ImageclassPixmaps` | **Bridged** | [epplet.go:L186](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L186) |
| `Epplet_textclass_draw(...)` | `Window.TextclassDraw(tclass, state string, x, y int, txt string)` | **Bridged** | [window.go:L147](file:///home/nemanja/projects/123-Go/dev/epplet-api/window.go#L147) |
| `Epplet_textclass_get_size(...)` | `TextclassGetSize(...) Size` | **Bridged** | [epplet.go:L210](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L210) |

---

### 6. 2D Primitive Drawing & RGB Buffers

| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_draw_line` | `DrawLine(...)` / `Window.DrawLine(...)` | **Bridged** | [drawing.go:L22](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L22) |
| `Epplet_draw_box` | `DrawBox(...)` / `Window.DrawBox(...)` | **Bridged** | [drawing.go:L32](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L32) |
| `Epplet_draw_outline` | `DrawOutline(...)` / `Window.DrawOutline(...)` | **Bridged** | [drawing.go:L42](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L42) |
| `Epplet_get_color` | `GetColor(r, g, b int) int` | **Bridged** | [drawing.go:L52](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L52) |
| `Epplet_paste_image` | `PasteImage(...)` / `Window.PasteImage(...)` | **Bridged** | [drawing.go:L57](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L57) |
| `Epplet_paste_image_size` | `PasteImageSize(...)` / `Window.PasteImageSize(...)` | **Bridged** | [drawing.go:L64](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L64) |
| `Esync` | `Sync()` | **Bridged** | [drawing.go:L71](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L71) |
| `Epplet_make_rgb_buf` | `MakeRGBBuf(w, h int) *RGBBuf` | **Bridged** | [drawing.go:L121](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L121) |
| `Epplet_get_rgb_pointer` | `RGBBuf.Data() []byte` | **Bridged** | [drawing.go:L145](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L145) |
| `Epplet_paste_buf` | `RGBBuf.Paste(win Window, x, y int)` / `Window.PasteBuf` | **Bridged** | [drawing.go:L156](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L156) |
| `Epplet_free_rgb_buf` | `RGBBuf.Free()` | **Bridged** | [drawing.go:L164](file:///home/nemanja/projects/123-Go/dev/epplet-api/drawing.go#L164) |

---

### 7. OpenGL / GLX Contexts

| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_bind_double_GL` | `DrawingArea.BindDoubleGL(...) GLXContext` | **Bridged** | [gl.go:L97](file:///home/nemanja/projects/123-Go/dev/epplet-api/gl.go#L97) |
| `Epplet_bind_single_GL` | `DrawingArea.BindSingleGL(...) GLXContext` | **Bridged** | [gl.go:L110](file:///home/nemanja/projects/123-Go/dev/epplet-api/gl.go#L110) |
| `Epplet_default_bind_GL` | `DrawingArea.DefaultBindGL() GLXContext` | **Bridged** | [gl.go:L123](file:///home/nemanja/projects/123-Go/dev/epplet-api/gl.go#L123) |
| `Epplet_unbind_GL` | `GLXContext.Unbind()` / `UnbindGL(...)` | **Bridged** | [gl.go:L154](file:///home/nemanja/projects/123-Go/dev/epplet-api/gl.go#L154) |

---

### 8. Command Execution & Dialog Helpers

| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_run_command` | `RunCommand(cmd string) int` | **Bridged** | [process.go:L34](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L34) |
| `Epplet_read_run_command` | `ReadRunCommand(cmd string) string` | **Bridged** | [process.go:L41](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L41) |
| `Epplet_spawn_command` | `SpawnCommand(cmd string) int` | **Bridged** | [process.go:L51](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L51) |
| `Epplet_pause_spawned_command` | `PauseSpawnedCommand(pid int)` | **Bridged** | [process.go:L58](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L58) |
| `Epplet_unpause_spawned_command` | `UnpauseSpawnedCommand(pid int)` | **Bridged** | [process.go:L63](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L63) |
| `Epplet_kill_spawned_command` | `KillSpawnedCommand(pid int)` | **Bridged** | [process.go:L68](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L68) |
| `Epplet_destroy_spawned_command` | `DestroySpawnedCommand(pid int)` | **Bridged** | [process.go:L73](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L73) |
| `Epplet_register_child_handler` | `RegisterChildHandler(handler ChildHandler)` | **Bridged** | [process.go:L91](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L91) |
| `Epplet_show_about` | `ShowAbout(name string)` | **Bridged** | [process.go:L104](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L104) |
| `Epplet_dialog_ok` | `DialogOk(text string)` | **Bridged** | [process.go:L111](file:///home/nemanja/projects/123-Go/dev/epplet-api/process.go#L111) |

---

### 9. Config Subsystem

| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_load_config` | `LoadConfig()` | **Bridged** | [config.go:L25](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L25) |
| `Epplet_load_config_file` | `LoadConfigFile(filename string)` | **Bridged** | [config.go:L30](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L30) |
| `Epplet_get_instance` | `GetInstance() int` | **Bridged** | [config.go:L37](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L37) |
| `Epplet_query_config` | `QueryConfig(key string) string` | **Bridged** | [config.go:L42](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L42) |
| `Epplet_query_config_def` | `QueryConfigDef(key, defaultVal string) string` | **Bridged** | [config.go:L53](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L53) |
| `Epplet_modify_config` | `ModifyConfig(key, value string)` | **Bridged** | [config.go:L66](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L66) |
| `Epplet_add_config` | `AddConfig(key, value string)` | **Bridged** | [config.go:L75](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L75) |
| `Epplet_modify_multi_config` | `ModifyMultiConfig(key string, values []string)` | **Bridged** | [config.go:L84](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L84) |
| `Epplet_query_multi_config` | `QueryMultiConfig(key string) []string` | **Bridged** | [config.go:L106](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L106) |
| `Epplet_save_config` | `SaveConfig()` | **Bridged** | [config.go:L131](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L131) |
| `Epplet_clear_config` | `ClearConfig()` | **Bridged** | [config.go:L136](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L136) |
| `Epplet_data_dir` | `DataDir() string` | **Bridged** | [config.go:L141](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L141) |
| `Epplet_e16_user_dir` | `E16UserDir() string` | **Bridged** | [config.go:L150](file:///home/nemanja/projects/123-Go/dev/epplet-api/config.go#L150) |

---

### 10. Event & Callback Handlers

| C Function in `epplet.h` | Go Function / Type | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_register_expose_handler` | `RegisterExposeHandler(handler ExposeHandler)` | **Bridged** | [callbacks.go:L147](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L147) |
| `Epplet_register_move_resize_handler` | `RegisterMoveResizeHandler(handler MoveResizeHandler)` | **Bridged** | [callbacks.go:L173](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L173) |
| `Epplet_register_button_press_handler` | `RegisterButtonPressHandler(handler ButtonPressHandler)` | **Bridged** | [callbacks.go:L198](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L198) |
| `Epplet_register_button_release_handler` | `RegisterButtonReleaseHandler(handler ButtonReleaseHandler)` | **Bridged** | [callbacks.go:L221](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L221) |
| `Epplet_register_key_press_handler` | `RegisterKeyPressHandler(handler KeyPressHandler)` | **Bridged** | [callbacks.go:L243](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L243) |
| `Epplet_register_key_release_handler` | `RegisterKeyReleaseHandler(handler KeyReleaseHandler)` | **Bridged** | [callbacks.go:L266](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L266) |
| `Epplet_register_mouse_motion_handler` | `RegisterMouseMotionHandler(handler MouseMotionHandler)` | **Bridged** | [callbacks.go:L289](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L289) |
| `Epplet_register_mouse_enter_handler` | `RegisterMouseEnterHandler(handler MouseEnterHandler)` | **Bridged** | [callbacks.go:L313](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L313) |
| `Epplet_register_mouse_leave_handler` | `RegisterMouseLeaveHandler(handler MouseLeaveHandler)` | **Bridged** | [callbacks.go:L335](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L335) |
| `Epplet_register_focus_in_handler` | `RegisterFocusInHandler(handler FocusInHandler)` | **Bridged** | [callbacks.go:L358](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L358) |
| `Epplet_register_focus_out_handler` | `RegisterFocusOutHandler(handler FocusOutHandler)` | **Bridged** | [callbacks.go:L381](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L381) |
| `Epplet_register_event_handler` | `RegisterEventHandler(handler XEventHandler)` | **Bridged** | [callbacks.go:L434](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L434) |
| `Epplet_register_delete_event_handler` | `RegisterDeleteEventHandler(handler DeleteEventHandler)` | **Bridged** | [callbacks.go:L459](file:///home/nemanja/projects/123-Go/dev/epplet-api/callbacks.go#L459) |

---

### 11. Timers

| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_timer(...)` | `Timer(cb TimerCallback, d time.Duration, name string)` | **Bridged** | [epplet.go:L266](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L266) |
| `Epplet_remove_timer(char *name)` | `RemoveTimer(name string)` | **Bridged** | [epplet.go:L285](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L285) |
| `Epplet_timer_get_data(char *name)` | `TimerGetData(name string) TimerCallback` | **Bridged** | [epplet.go:L299](file:///home/nemanja/projects/123-Go/dev/epplet-api/epplet.go#L299) |
| `Epplet_get_time()` | *None* | Unbridged | Replaced by Go standard `time` library |

---

### 12. Gadgets & Widgets

#### Base Gadget Attributes & Controls
| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_gadget_get_x` | `Gadget.GetX() int` | **Bridged** | [gadget.go:L149](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L149) |
| `Epplet_gadget_get_y` | `Gadget.GetY() int` | **Bridged** | [gadget.go:L153](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L153) |
| `Epplet_gadget_get_width` | `Gadget.GetWidth() int` | **Bridged** | [gadget.go:L157](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L157) |
| `Epplet_gadget_get_height` | `Gadget.GetHeight() int` | **Bridged** | [gadget.go:L161](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L161) |
| `Epplet_gadget_get_type` | `Gadget.GetType() int` | **Bridged** | [gadget.go:L165](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L165) |
| `Epplet_gadget_show` | `Gadget.Show()` | **Bridged** | [gadget.go:L173](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L173) |
| `Epplet_gadget_hide` | `Gadget.Hide()` | **Bridged** | [gadget.go:L177](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L177) |
| `Epplet_gadget_move` | `Gadget.Move(x, y int)` | **Bridged** | [gadget.go:L181](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L181) |
| `Epplet_gadget_destroy` | `Gadget.Destroy()` | **Bridged** | [gadget.go:L185](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L185) |
| `Epplet_gadget_data_changed` | `Gadget.DataChanged()` | **Bridged** | [gadget.go:L189](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L189) |
| `Epplet_gadget_draw` | `Gadget.Draw(unOnly, force bool)` | **Bridged** | [gadget.go:L193](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L193) |
| `Epplet_redraw` | `Redraw()` | **Bridged** | [gadget.go:L124](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L124) |
| `Epplet_gadget_get_data` | *None* | Unbridged | N/A in Go (callbacks managed via cgo.Handle) |

#### Buttons
| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_create_button` | `CreateButton(...) *Button` | **Bridged** | [gadget.go:L217](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L217) |
| `Epplet_create_text_button` | `CreateTextButton(...) *Button` | **Bridged** | [gadget.go:L260](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L260) |
| `Epplet_create_std_button` | `CreateStdButton(...) *Button` | **Bridged** | [gadget.go:L287](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L287) |
| `Epplet_create_image_button` | `CreateImageButton(...) *Button` | **Bridged** | [gadget.go:L312](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L312) |
| `Epplet_change_button_label` | `Button.ChangeLabel(label string)` | **Bridged** | [gadget.go:L339](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L339) |
| `Epplet_change_button_image` | `Button.ChangeImage(image string)` | **Bridged** | [gadget.go:L347](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L347) |

#### Textbox
| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_create_textbox` | `CreateTextbox(...) *Textbox` | **Bridged** | [gadget.go:L359](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L359) |
| `Epplet_textbox_contents` | `Textbox.Contents() string` | **Bridged** | [gadget.go:L393](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L393) |
| `Epplet_reset_textbox` | `Textbox.Reset()` | **Bridged** | [gadget.go:L403](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L403) |
| `Epplet_change_textbox` | `Textbox.ChangeContents(string)` | **Bridged** | [gadget.go:L410](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L410) |
| `Epplet_textbox_insert` | `Textbox.InsertContents(string)` | **Bridged** | [gadget.go:L419](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L419) |

#### DrawingArea
| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_create_drawingarea` | `CreateDrawingArea(...) *DrawingArea` | **Bridged** | [gadget.go:L432](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L432) |
| `Epplet_get_drawingarea_window` | `DrawingArea.DrawingAreaWindow() Window` | **Bridged** | [gadget.go:L440](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L440) |

#### Sliders & Toggle Buttons
| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_create_hslider` | `CreateHSlider(...) *Slider` | **Bridged** | [gadget.go:L451](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L451) |
| `Epplet_create_vslider` | `CreateVSlider(...) *Slider` | **Bridged** | [gadget.go:L470](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L470) |
| `Epplet_get_hslider_clicked` | `Slider.HSliderClicked() bool` | **Bridged** | [gadget.go:L489](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L489) |
| `Epplet_get_vslider_clicked` | `Slider.VSliderClicked() bool` | **Bridged** | [gadget.go:L496](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L496) |
| `Epplet_create_togglebutton` | `CreateToggleButton(...) *ToggleButton` | **Bridged** | [gadget.go:L507](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L507) |

#### Popups & PopupButtons
| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_create_popup` | `CreatePopup() *Popup` | **Bridged** | [gadget.go:L540](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L540) |
| `Epplet_add_popup_entry` | `Popup.AddEntry(label, pixmap, cb)` | **Bridged** | [gadget.go:L548](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L548) |
| `Epplet_add_sized_popup_entry` | `Popup.AddSizedEntry(label, pixmap, w, h, cb)` | **Bridged** | [gadget.go:L573](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L573) |
| `Epplet_remove_popup_entry` | `Popup.RemoveEntry(entryNum int)` | **Bridged** | [gadget.go:L598](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L598) |
| `Epplet_popup_entry_num` | `Popup.EntryNum() int` | **Bridged** | [gadget.go:L605](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L605) |
| `Epplet_pop_popup` | `Popup.Pop(win Window)` | **Bridged** | [gadget.go:L612](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L612) |
| `Epplet_popup_entry_get_data` | *None* | Unbridged | N/A in Go (callbacks managed via cgo.Handle) |
| `Epplet_create_popupbutton` | `CreatePopupButton(...) *PopupButton` | **Bridged** | [gadget.go:L619](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L619) |
| `Epplet_change_popbutton_popup` | `PopupButton.ChangePopup(popup *Popup)` | **Bridged** | [gadget.go:L649](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L649) |
| `Epplet_change_popbutton_label` | `PopupButton.ChangeLabel(label string)` | **Bridged** | [gadget.go:L659](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L659) |

#### Image & Label Gadgets
| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_create_image` | `CreateImage(...) *ImageGadget` | **Bridged** | [gadget.go:L672](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L672) |
| `Epplet_change_image` | `ImageGadget.Change(w, h int, image string)` | **Bridged** | [gadget.go:L684](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L684) |
| `Epplet_move_change_image` | `ImageGadget.MoveChange(x, y, w, h int, image string)` | **Bridged** | [gadget.go:L695](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L695) |
| `Epplet_create_label` | `CreateLabel(x, y int, label string, size int) *Label` | **Bridged** | [gadget.go:L710](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L710) |
| `Epplet_change_label` | `Label.Change(label string)` | **Bridged** | [gadget.go:L722](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L722) |
| `Epplet_move_change_label` | `Label.MoveChange(x, y int, label string)` | **Bridged** | [gadget.go:L731](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L731) |

#### Progress Bars
| C Function in `epplet.h` | Go Function / Method | Status | Notes / Location |
| :--- | :--- | :---: | :--- |
| `Epplet_create_hbar` | `CreateHBar(...) *ProgressBar` | **Bridged** | [gadget.go:L744](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L744) |
| `Epplet_create_vbar` | `CreateVBar(...) *ProgressBar` | **Bridged** | [gadget.go:L754](file:///home/nemanja/projects/123-Go/dev/epplet-api/gadget.go#L754) |

---

### 13. Unbridged Features (C API Only)

The following functions in `epplet.h` do not currently have direct Go wrappers in `epplet-api`:

#### Callback Data Pointers (N/A in Go)
- `Epplet_gadget_get_data` (Callbacks managed in Go using `cgo.Handle`)
- `Epplet_popup_entry_get_data` (Callbacks managed in Go using `cgo.Handle`)
