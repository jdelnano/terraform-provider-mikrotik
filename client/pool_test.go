package client

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestAddPoolAndDeletePool(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	name := "testacc"
	ranges := "172.16.0.1-172.16.0.8,172.16.0.10"
	comment := "terraform-acc-test-pool"
	expectedPool := &Pool{
		Name:    name,
		Ranges:  ranges,
		Comment: comment,
	}
	pool, err := c.AddPool(
		name,
		ranges,
		comment,
	)

	fmt.Sprintf("expected pool:  %v", expectedPool)
	fmt.Sprintf("created pool:  %v", pool)
	if err != nil {
		t.Errorf("Error creating a pool with: %v", err)
	}

	if len(pool.Id) < 1 {
		t.Errorf("The created pool does not have an Id: %v", pool)
	}

	if strings.Compare(pool.Name, expectedPool.Name) != 0 {
		t.Errorf("The pool Name fields do not match. actual: %v expected: %v", pool.Name, expectedPool.Name)
	}

	if strings.Compare(pool.Ranges, expectedPool.Ranges) != 0 {
		t.Errorf("The pool Ranges fields do not match. actual: %v expected: %v", pool.Ranges, expectedPool.Ranges)
	}

	if strings.Compare(pool.Comment, expectedPool.Comment) != 0 {
		t.Errorf("The pool Comment fields do not match. actual: %v expected: %v", pool.Comment, expectedPool.Comment)
	}

	foundPool, err := c.FindPool(pool.Id)

	if err != nil {
		t.Errorf("Error getting pool with: %v", err)
	}

	if !reflect.DeepEqual(pool, foundPool) {
		t.Errorf("Created pool and found pool do not match. actual: %v expected: %v", foundPool, pool)
	}

	err = c.DeletePool(pool.Id)

	if err != nil {
		t.Errorf("Error deleting pool with: %v", err)
	}
}

func TestFindPool_forNonExistingPool(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	poolId := "Invalid id"
	_, err := c.FindPool(poolId)

	expectedErrStr := fmt.Sprintf("pool `%s` not found", poolId)
	if err == nil || err.Error() != expectedErrStr {
		t.Errorf("client should have received error indicating the following pool with id `%s` was not found. Instead error was nil", poolId)
	}
}
