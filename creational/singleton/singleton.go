package singleton

type singleton struct {}

var instance *singleton

// lazy loading
func getInstanceLazy() *singleton {
	if instance == nil {
		instance = &singleton{}
		return instance
	}
	return instance
}


