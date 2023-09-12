package types

const (
	// ModuleName defines the module name
	ModuleName = "sync"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sync"
)

var (
	KeyPrefixHeader            = []byte{0x01}
	KeyPrefixHeaderCount       = []byte{0x01}
	KeyPrefixAdmin             = []byte{0x02}
	KeyPrefixHeaderHashMapping = []byte{0x02}
)
