package core

func Info(query string) App {
	object := SearchPkg(query)

	if len(object.AppDetail) == 0 {
		return App{}
	}

	for _, v := range object.AppDetail {
		if v.PkgName == query {
			return v
		}
	}

	return App{}
}
