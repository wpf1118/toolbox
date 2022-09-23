package filer

import (
	"fmt"
	"os"
	"testing"
)

func TestFiler(t *testing.T) {
	var err error

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	currentDir += "/test/"

	t.Log("TestFiler")
	t.Log("test put----")
	t.Log(File(fmt.Sprintf("%s/0t1.txt", currentDir)).Err(&err).Put([]byte(`hello world`)))
	t.Log("test get----")
	t.Log(File(fmt.Sprintf("%s/0t1.txt", currentDir)).Err(&err).Get())
	t.Log("test open clear ----")
	t.Log(File(fmt.Sprintf("%s/0t1.txt", currentDir)).Err(&err).TypeClear().Put([]byte(`hello world`)))
	t.Log("test move ----")
	t.Log(File(fmt.Sprintf("%s/0t1.txt", currentDir)).Err(&err).Rename(fmt.Sprintf("%s/1t1.txt", currentDir)))
	t.Log("test remove ----")
	t.Log(File(fmt.Sprintf("%s/1t1.txt", currentDir)).Err(&err).Remove())

	//测试写入json
	type Config struct {
		IP string
	}
	var c Config = Config{
		IP: "s",
	}
	t.Log(File(fmt.Sprintf("%s/0t1.json", currentDir)).Err(&err).PutJson(c))

	t.Log("TestFiler end", err)
}
