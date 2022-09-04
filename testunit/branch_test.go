package testunit_test

import (
	"testing"

	"github.com/tapvanvn/go-dashboard/runtime"
)

func TestBranchFinding(t *testing.T) {

	if branch, err := runtime.GetBranch("test"); err == nil {

		if subBranch, err := runtime.GetBranch("test.test2"); err == nil {

			subBranch2 := branch.Find([]string{"test2"})

			if subBranch == subBranch2 {

				return
			}
		}
	}
	t.Error()
}
