package types

const (
	// ModuleName defines the module name
	ModuleName = "sportiverse"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sportiverse"
)

var (
	ParamsKey = []byte("p_sportiverse")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
