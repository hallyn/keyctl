package keyctl
//taken from https://github.com/jsipprell/keyctl/pull/17
const (
	syscall_keyctl   uintptr = 219
	syscall_add_key  uintptr = 217
	syscall_setfsgid uintptr = 152
)