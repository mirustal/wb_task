package set




func SetValue(strings []string) (map[string]bool) {
	set := make(map[string]bool)
	for _, s := range strings {
		set[s] = true
	}
	return set
}
