package registry

// Registry 定义注册表结构
type Registry struct {
	registry map[string]interface{}
}

// Register 方法用于向注册表中注册对象
func (r *Registry) Register(key string, value interface{}) {
	r.registry[key] = value
}

// Get 方法用于从注册表中检索对象
func (r *Registry) Get(key string) interface{} {
	return r.registry[key]
}
