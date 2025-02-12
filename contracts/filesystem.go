package contracts

import (
	"bufio"
	"io/fs"
	"time"
)

type FileVisibility int

type FileSystemProvider func(config Fields) FileSystem

type File interface {
	fs.FileInfo
	Get() string
	Disk() string
}

type FileSystemFactory interface {
	Disk(disk string) FileSystem
	Extend(driver string, provider FileSystemProvider)

	FileSystem
}

type FileSystem interface {
	Name() string

	Exists(path string) bool

	Get(path string) (string, error)
	ReadStream(path string) (*bufio.Reader, error)

	Put(path, contents string) error
	WriteStream(path string, contents string) error

	GetVisibility(path string) FileVisibility
	SetVisibility(path string, perm fs.FileMode) error

	Prepend(path, contents string) error

	Append(path, contents string) error

	Delete(path string) error

	Copy(from, to string) error

	Move(from, to string) error

	Size(path string) (int64, error)

	LastModified(path string) (time.Time, error)

	Files(directory string) []File

	AllFiles(directory string) []File

	Directories(directory string) []string

	AllDirectories(directory string) []string

	MakeDirectory(path string) error

	DeleteDirectory(directory string) error
}
