package server

import (
	"Exercise/RPC/sensitiveWord/trie"
	pb "Exercise/RPC/sensitiveWord/wordfilter"
	"golang.org/x/net/context"
)

var t =trie.InitTrie()//初始化敏感词trie树

type Router struct {
}

func (r *Router) ContainsSensitiveWord(cxt context.Context, request *pb.ContainSensitiveWordRequest) (*pb.ContainSensitiveWordResponse, error) {
	ok := t.Query(request.Text)
	return &pb.ContainSensitiveWordResponse{IsSensitiveWord: ok}, nil
}
