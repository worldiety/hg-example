package invoicing

import (
	"fmt"
	"github.com/worldiety/hg-example/internal/invoicing/domain/positions"
)

type InvoiceDraftError interface {
	error
	InvoiceDraft() bool
}

// CreateInvoiceDraft implements the named usecase.
func CreateInvoiceDraft(pos []positions.Position) (positions.CollectionID, InvoiceDraftError) {
	return "", positions.InfrastructureError{Cause: fmt.Errorf("just a scribble demo")}
}
