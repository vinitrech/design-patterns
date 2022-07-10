/*
Used to define a runtime interface for creating an object.
It’s called a factory because it creates various types of objects
without necessarily knowing what kind of object it creates or how to create it.

The abstract factory provides a way to encapsulate a group of individual factories
that have a common theme without specifying their concrete classes.
*/

package factory

import "fmt"

type (
	mongoDB struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}

	file struct {
		name    string
		content string
	}

	ntfs struct {
		files map[string]file
	}

	ext4 struct {
		files map[string]file
	}

	FileSystem interface {
		CreateFile(string)
		FindFile(string) file
	}

	Database interface {
		GetData(string) string
		PutData(string, string)
	}

	Factory func(string) interface{}
)

func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}

	fmt.Println("MongoDB")
	return mdb.database[query]
}

func (sql sqlite) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}

	fmt.Println("Sqlite")
	return sql.database[query]
}

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sql sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

func (ntfs ntfs) CreateFile(path string) {
	file := file{content: "NTFS file", name: path}
	ntfs.files[path] = file
	fmt.Println("NTFS")
}

func (ext ext4) CreateFile(path string) {
	file := file{content: "EXT4 file", name: path}
	ext.files[path] = file
	fmt.Println("EXT4")
}

func (ntfs ntfs) FindFile(path string) file {
	if _, ok := ntfs.files[path]; !ok {
		return file{}
	}

	return ntfs.files[path]
}

func (ext ext4) FindFile(path string) file {
	if _, ok := ext.files[path]; !ok {
		return file{}
	}

	return ext.files[path]
}

func FilesystemFactory(env string) interface{} {
	switch env {
	case "production":
		return ntfs{
			files: make(map[string]file),
		}
	case "development":
		return ext4{
			files: make(map[string]file),
		}
	default:
		return nil
	}
}

func DatabaseFactory(env string) interface{} {
	switch env {
	case "production":
		return mongoDB{
			database: make(map[string]string),
		}
	case "development":
		return sqlite{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}

func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FilesystemFactory
	default:
		return nil
	}
}
