package set

func NewStringSet(keys ...string) StringSet {
	set := StringSet{}
	return set.Add(keys...)
}

type StringSet map[string]struct{}

// 交集
func (s StringSet) Intersect(ss StringSet) StringSet {
	set := StringSet{}
	for key := range s {
		_, ok := ss[key]
		if ok {
			set[key] = struct{}{}
		}
	}
	return set
}

// 并集
func (s StringSet) Union(ss StringSet) StringSet {
	set := StringSet{}
	for key := range s {
		set[key] = struct{}{}
	}
	for key := range ss {
		set[key] = struct{}{}
	}
	return set
}

// 差集
func (s StringSet) Expect(ss StringSet) StringSet {
	set := StringSet{}
	for key := range s {
		_, ok := ss[key]
		if !ok {
			set[key] = struct{}{}
		}
	}
	return set
}

// 集合是否相等
func (s StringSet) Equals(ss StringSet) bool {
	if s.Len() != ss.Len() {
		return false
	}
	for key := range ss {
		_, ok := s[key]
		if !ok {
			return false
		}
	}
	return true
}

// 集合是否包含
func (s StringSet) Contains(ss StringSet) bool {
	return s.Has(ss.Keys()...)
}

func (s StringSet) Add(keys ...string) StringSet {
	for _, key := range keys {
		s[key] = struct{}{}
	}
	return s
}

func (s StringSet) Del(keys ...string) StringSet {
	for _, key := range keys {
		delete(s, key)
	}
	return s
}

func (s StringSet) Has(keys ...string) bool {
	for _, key := range keys {
		_, ok := s[key]
		if !ok {
			return false
		}
	}
	return true
}

func (s StringSet) Any(keys ...string) bool {
	for _, key := range keys {
		_, ok := s[key]
		if ok {
			return true
		}
	}
	return false
}

func (s StringSet) Len() int {
	return len(s)
}

func (s StringSet) Keys() []string {
	keys := []string{}
	for key := range s {
		keys = append(keys, key)
	}
	return keys
}
