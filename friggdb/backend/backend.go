package backend

import (
	"context"

	"github.com/google/uuid"
)

type Writer interface {
	Write(ctx context.Context, blockID uuid.UUID, tenantID string, bMeta []byte, bBloom []byte, bIndex []byte, objectFilePath string) error
}

type Reader interface {
	Tenants() ([]string, error)
	Blocks(tenantID string) ([]uuid.UUID, error)
	BlockMeta(blockID uuid.UUID, tenantID string) ([]byte, error)
	Bloom(blockID uuid.UUID, tenantID string) ([]byte, error)
	Index(blockID uuid.UUID, tenantID string) ([]byte, error)
	Object(blockID uuid.UUID, tenantID string, start uint64, length uint32) ([]byte, error)

	Shutdown()
}

type Compactor interface {
	MarkBlockCompacted(blockID uuid.UUID, tenantID string) error
	ClearBlock(blockID uuid.UUID, tenantID string) error
	CompactedBlocks(tenantID string) ([]uuid.UUID, error)
}
