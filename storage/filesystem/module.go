package filesystem

import (
	"github.com/cloudperson/go-git-ex.v4/plumbing/cache"
	"github.com/cloudperson/go-git-ex.v4/storage"
	"github.com/cloudperson/go-git-ex.v4/storage/filesystem/dotgit"
)

type ModuleStorage struct {
	dir *dotgit.DotGit
}

func (s *ModuleStorage) Module(name string) (storage.Storer, error) {
	fs, err := s.dir.Module(name)
	if err != nil {
		return nil, err
	}

	return NewStorage(fs, cache.NewObjectLRUDefault()), nil
}
