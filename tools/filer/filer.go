package filer

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Filer struct {
	err      *error
	path     string
	openType int
	openMode os.FileMode
	handle   *os.File
}

func File(path string) *Filer {
	var f *Filer = new(Filer)
	f.path = path
	return f
}

func (f *Filer) GetHandle() *os.File {
	if f.openType == 0 {
		f.TypeRead()
	}
	f.initHandle()
	return f.handle
}

func (f *Filer) Err(err *error) *Filer {
	f.err = err
	return f
}

func (f *Filer) TypeOpen() *Filer {
	//var f *file.File = file.NewFile(f.path)
	//f.Open()
	f.openType = os.O_CREATE | os.O_APPEND | os.O_RDWR
	return f
	//var info string = f.GetInfo()
	//f.Close()
}

func (f *Filer) TypeRead() *Filer {
	f.openType = os.O_RDONLY
	return f
}

func (f *Filer) TypeClear() *Filer {
	f.openType = os.O_TRUNC | os.O_CREATE | os.O_RDWR
	return f
}

func (f *Filer) TypeMode(mode os.FileMode) *Filer {
	f.openMode = mode
	return f
}

func (f *Filer) initHandle() *Filer {
	if f.openMode == 0 {
		f.openMode = 0777
	}
	var err error
	f.handle, err = os.OpenFile(f.path, f.openType, os.FileMode(f.openMode))
	f.err = &err
	return f
}

func (f *Filer) PutJson(v interface{}) (write int, err error) {
	data, err := json.Marshal(v)
	if err != nil {
		return 0, err
	}
	return f.Put(data)
}

func (f *Filer) Put(content []byte) (write int, err error) {
	if f.openType == 0 {
		f.TypeOpen()
	}
	//如果文件夹不存在，尝试创建文件夹
	//if folder not exist, try to create folder
	var pos int = strings.LastIndex(f.path, "/")
	if pos > 0 {
		if f.openMode == 0 {
			f.openMode = 0777
		}
		os.MkdirAll(f.path[0:pos], f.openMode)
	}
	f.initHandle()
	if *f.err != nil {
		f.err = &err
		return 0, err
	}
	//写入\n
	content = append(content, 10)
	write, err = f.handle.Write(content)
	f.handle.Close()
	return write, err
}

func (f *Filer) Remove() (err error) {
	err = os.Remove(f.path)
	f.err = &err
	return err
}

func (f *Filer) Rename(filename string) (err error) {
	err = os.Rename(f.path, filename)
	f.err = &err
	return err
}

func (f *Filer) Exist() bool {
	var err error
	var fileInfo os.FileInfo
	f.TypeRead()
	f.initHandle()
	fileInfo, err = f.handle.Stat()
	if err == nil {
		if fileInfo.Mode().IsRegular() {
			return true
		}
	}
	return false
}

func (f *Filer) Get() (content []byte, err error) {
	if f.openType == 0 {
		f.TypeRead()
	}
	f.initHandle()
	if *f.err != nil {
		return nil, *f.err
	}
	content, err = ioutil.ReadAll(f.handle)
	if err != nil {
		f.err = &err
	}
	f.handle.Close()
	return content, err
}

func (f *Filer) ReadLine(handler func(string) error) (err error) {

	if f.openType == 0 {
		f.TypeRead()
	}
	f.initHandle()
	if *f.err != nil {
		return *f.err
	}

	buf := bufio.NewReader(f.handle)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			break
		}
		if err = handler(line); err != nil {
			break
		}

	}

	f.handle.Close()
	return err

}

func (f *Filer) IsExist() (exist bool) {
	if _, err := os.Stat(f.path); err == nil {
		return true
	}
	return false
}

func (f *Filer) GetAllFilenames() (filenames []string) {
	if !f.IsExist() {
		return
	}
	err := filepath.Walk(f.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filenames = append(filenames, info.Name())
		}
		return nil
	})
	if err != nil {
		return
	}

	return
}
