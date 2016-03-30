package fswatch

type StringCache map[string]string

func NewStringCache() StringCache {
	return StringCache{}
}

func (c StringCache) Equals(cache StringCache) bool {
	mergedCache := Merge(c, cache)
	for key, value := range mergedCache {
		// updated value for key or a value does not exist for a pre-existing key
		if _, ok := cache[key]; (value != c[key] && value == cache[key]) || !ok {
			return false
		}
	}
	return true
}

func (c StringCache) Diff(cache StringCache) (diff []string) {
	mergedCache := Merge(c, cache)
	for key, value := range mergedCache {
		// updated value for key or a value does not exist for a pre-existing key
		if _, ok := cache[key]; (value != c[key] && value == cache[key]) || !ok {
			diff = append(diff, key)
		}
	}
	return
}

func (c StringCache) Keys() []string {
	keys := make([]string, len(c))
	i := 0
	for key := range c {
		keys[i] = key
		i++
	}
	return keys
}

func Merge(cache1, cache2 StringCache) StringCache {
	cache := StringCache{}
	for key, value := range cache1 {
		cache[key] = value
	}
	for key, value := range cache2 {
		cache[key] = value
	}

	return cache
}
