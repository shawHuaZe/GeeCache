package geecache

// A ByteView holds an immutable view of bytes.
type ByteView struct {
	b []byte
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
// 这里返回clone是为了防止缓存值被外部程序修改
// 高层次 (关心如何安全地暴露 ByteView 的数据)
// 目的 保证 ByteView 的不可变性，同时提供数据访问。
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

// 私有方法
// 低层次 (只关心如何复制字节)， 目的执行复制操作
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
