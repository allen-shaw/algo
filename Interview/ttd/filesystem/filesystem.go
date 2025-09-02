package filesystem

import (
	"errors"
	"sort"
	"strings"
)

// Level 1: Add file with file name and file size, then delete file by name.
// Level 2: Return largest n files starting with given prefix in descending order.
// Level 3: Add user which capacity limit, then user can add file but should not exceed the capacity limit.
// Also merge 2 users, for example to transfer all files from user2 to user1 and give all remaining capacity

type File struct {
	Name string
	Size int
}

func newFile(name string, size int) *File {
	return &File{
		Name: name,
		Size: size,
	}
}

type User struct {
	name     string
	capacity int
	files    map[string]*File
	fileSize int
}

func newUser(name string, cap int) *User {
	return &User{
		name:     name,
		capacity: cap,
		files:    make(map[string]*File),
		fileSize: 0,
	}
}

type FileSystem struct {
	files map[string]*File
	users map[string]*User
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		files: make(map[string]*File),
		users: make(map[string]*User),
	}
}

func (fs *FileSystem) AddFile(name string, size int) error {
	_, ok := fs.files[name]
	if ok {
		return errors.New("file existed")
	}
	f := newFile(name, size)
	fs.files[name] = f
	return nil
}

func (fs *FileSystem) DeleteFile(name string) error {
	_, ok := fs.files[name]
	if !ok {
		return errors.New("file not existed")
	}
	delete(fs.files, name)
	return nil
}

func (fs *FileSystem) GetLargestFiles(prefix string, n int) []*File {
	files := make([]*File, 0)
	for _, f := range fs.files {
		if strings.HasPrefix(f.Name, prefix) {
			files = append(files, f)
		}
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].Size == files[j].Size {
			return files[i].Name < files[j].Name
		}
		return files[i].Size > files[j].Size
	})

	if len(files) <= n {
		return files
	}
	return files[:n]
}

func (fs *FileSystem) AddUser(name string, cap int) error {
	if cap < 0 {
		return errors.New("invalid cap")
	}
	_, ok := fs.users[name]
	if ok {
		return errors.New("user existed")
	}

	user := newUser(name, cap)
	fs.users[name] = user
	return nil
}

func (fs *FileSystem) AddFileToUser(file *File, user *User) error {
	_, ok := user.files[file.Name]
	if ok {
		return errors.New("file existed")
	}

	if file.Size+user.fileSize > user.capacity {
		return errors.New("file size too large")
	}

	user.files[file.Name] = file
	user.fileSize += file.Size
	return nil
}

// from user2 to user1
func (fs *FileSystem) MergeUsers(user1, user2 *User) error {
	if user1 == user2 {
		return errors.New("user1 is user2")
	}

	user1.capacity += user2.capacity

	for _, f := range user2.files {
		if _, ok := user1.files[f.Name]; ok {
			continue
		}
		fs.AddFileToUser(f, user1)
	}

	delete(fs.users, user2.name)
	return nil
}
