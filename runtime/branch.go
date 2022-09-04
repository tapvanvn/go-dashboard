package runtime

import (
	"strings"

	"github.com/tapvanvn/go-dashboard/common"
	"github.com/tapvanvn/go-dashboard/entity"
	"github.com/tapvanvn/goutil"
)

var __root *entity.Branch = entity.NewBranch()
var __registries map[string]*entity.Registry = make(map[string]*entity.Registry)

func GetBranch(path string) (*entity.Branch, error) {

	parts := strings.Split(path, ".")

	if len(parts) == 0 {

		return nil, common.ErrInvalidPath
	}
	rootName := parts[0]

	root, ok := __root.SubBranches[rootName]
	if !ok {

		root = entity.NewBranch()

		__root.SubBranches[rootName] = root
	}
	if len(parts) == 1 {

		return root, nil
	}

	hash := goutil.MD5Message(path)

	if _, ok := __registries[rootName]; !ok {

		__registries[rootName] = entity.NewRegistry(rootName)
	}

	if branch, ok := __registries[rootName].Hashes[hash]; ok {

		return branch, nil

	}

	sub := root.Find(parts[1:])
	__registries[rootName].Hashes[hash] = sub
	return sub, nil

}
