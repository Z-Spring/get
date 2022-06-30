package args

func ConvertSliceToMap(s []string) map[string]struct{} {
	set := make(map[string]struct{}, len(s))
	for _, v := range s {
		set[v] = struct{}{}
	}
	return set
}

func IsContain(m map[string]struct{}, target []string) bool {
	if len(target) > 1 {
		for _, v := range target {
			_, ok := m[v]
			return ok
		}
	} else if len(target) == 1 {
		_, ok := m[target[0]]
		return ok
	}
	return false
}

func IsContain2(target string, m map[string]struct{}) bool {
	_, ok := m[target]
	return ok
}
