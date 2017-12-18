package dynamodb

import (
	"gitlab.com/fanky5g/paysenger/server/types"
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetItemsFilterWithPaginate(t *testing.T) {
	var lEvaluated interface{}
	more := true

	for more {
		items, lastEvaluatedKey, d, err := GetItemFilterWithPaginate("Transactions", "ReferenceIndex", "reference", "SEND_AIRTIME", "", struct{}{}, lEvaluated, 2, types.Transaction{})
		assert.NoError(t, err)
		for _, item := range items {
			t.Log(item.(*types.Transaction))
		}

		if lastEvaluatedKey != nil {
			lEvaluated = lastEvaluatedKey.(*types.Transaction)
		}
		if d == true {
			more = false
		}
	}
	t.Log("done listing all transactions")
}
