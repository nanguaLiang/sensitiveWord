package trie

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"errors"
)

//var blackTrie *Trie


// InitTrie 初始化Trie
func InitTrie() *Trie{
	projectRoot := fmt.Sprintf("%s/src/Exercise/RPC/sensitiveWord",os.Getenv("GOPATH"))
	dirPath := projectRoot+"/dicts"
	trie:=NewTrie()
	loadDict(trie,dirPath)
	return trie
}

// BlackTrie 返回黑名单Trie树
//func BlackTrie(dirPath string) *Trie {
//	if blackTrie == nil {
//		blackTrie = NewTrie()
//		loadDict(blackTrie,dirPath)
//	}
//	return blackTrie
//}

func loadDict(trieHandle *Trie, path string) {
	var loadAllDictWalk filepath.WalkFunc = func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		initTrie(trieHandle,path)

		return nil
	}

	err := filepath.Walk(path, loadAllDictWalk)
	if err != nil {
		panic(err)
	}
}

func initTrie(trieHandle *Trie,path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("fail to open file %s %s", path, err.Error()))
	}

	defer f.Close()

	log.Printf("%s Load dict: %s", time.Now().Local().Format("2006-01-02 15:04:05 -0700"), path)

	buf := bufio.NewReader(f)
	for {
		line, isPrefix, e := buf.ReadLine()
		if e != nil {
			if e != io.EOF {
				err = e
			}
			break
		}
		if isPrefix {
			continue
		}

		if word := strings.TrimSpace(string(line)); word != "" {
			tmp := strings.Split(word, " ")
			s := strings.Trim(tmp[0], " ")
			trieHandle.Add(s)
		}
	}

	return
}

//通过敏感词数组去构建trie树
func buildTrieByWords(trie *Trie,words []string) error{
	if len(words)==0{
		return errors.New("the words are nil")
	}
	for _,word:=range words{
		trie.Add(word)
	}
	return nil
}

//通过读文件去获取敏感词数组
func readFile(path string)(words []string,err error){
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("fail to open file %s %s", path, err.Error()))
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, isPrefix, e := buf.ReadLine()
		if e != nil {
			if e != io.EOF {
				err = e
			}
			break
		}
		if isPrefix {
			continue
		}

		if word := strings.TrimSpace(string(line)); word != "" {
			tmp := strings.Split(word, " ")
			s := strings.Trim(tmp[0], " ")
			words=append(words, s)
		}
	}
	return

}