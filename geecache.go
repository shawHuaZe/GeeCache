package geecache

import (
	"sync"
)

// A Getter loads data for a key.
// 回调函数
type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function.
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
/*
我们有一个 Getter 接口，它需要一个 Get 方法。
我们创建了一个 GetterFunc 类型，它是一个函数类型。
我们为 GetterFunc 这个类型实现了 Get 方法。
结论：因此，GetterFunc 类型现在满足了 Getter 接口。
这意味着，任何一个普通的、符合 func(key string) ([]byte, error) 签名的函数，
我们只需要通过一个简单的类型转换，把它变成 GetterFunc 类型，它就立刻可以被用在任何需要 Getter 接口的地方。
*/
// GetterFunc(getFromDB) 此方法用于类型转换
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// A Group is a cache namespace and associated data loaded spread over
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

// Go允许使用 var () 语法将多个变量声明合并，避免重复写 var 关键字
// 等价于分开声明：
var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// NewGroup create a new instance of Group
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter") //触发运行时错误，类似raise
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

// GetGroup returns the named group previously created with NewGroup, or
// nil if there's no such group.
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}
