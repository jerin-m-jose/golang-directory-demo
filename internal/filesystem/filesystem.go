package filesystem

// FileSystem interface to be used by in memory file system
type FileSystem interface {
	Create(name string) error
	List(sort Sort)
	Move(src string, dest string) error
	Delete(name string) error
	GetChildren() interface{}
}
