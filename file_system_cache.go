package fswatch

type fileSystemCache struct {
	contents map[string]string
}

func newFileSystemCache() *fileSystemCache {
	return &fileSystemCache{contents: map[string]string{}}
}

func (f *fileSystemCache) Equals(values map[string]string) bool {
	for key, value := range f.contents {
		if values[key] != value {
			return false
		}
	}
	return true
}

func (f *fileSystemCache) Set(newCache map[string]string) {
	f.contents = newCache
}

func (f *fileSystemCache) Diff(values map[string]string) []string {
	changed := []string{}
	for key, value := range f.contents {
		if value != values[key] {
			changed = append(changed, key)
		}
	}
	return changed
}
