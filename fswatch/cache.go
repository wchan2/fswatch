package fswatch

type FSCache interface {
	Compare(map[string]string) bool
	Update(map[string]string)
	Diff(map[string]string) []string
}

func NewFileSystemCache() FSCache {
	return &cache{contents: map[string]string{}}
}

type cache struct {
	contents map[string]string
}

func (c *cache) Compare(values map[string]string) bool {
	for key, value := range c.contents {
		if values[key] != value {
			return false
		}
	}
	return true
}

func (c *cache) Update(newCache map[string]string) {
	c.contents = newCache
}

func (c *cache) Diff(values map[string]string) []string {
	changed := []string{}
	for key, value := range c.contents {
		if value != values[key] {
			changed = append(changed, key)
		}
	}
	return changed
}
