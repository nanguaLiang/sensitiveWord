package server

import (
	"Exercise/RPC/sensitiveWord/trie"
	pb "Exercise/RPC/sensitiveWord/wordfilter"
	"golang.org/x/net/context"
)

type Router struct {
}

func (r *Router) ContainsSensitiveWord(cxt context.Context, request *pb.ContainSensitiveWordRequest) (*pb.ContainSensitiveWordResponse, error) {
	ok, _, _ := trie.BlackTrie().Query(request.Text)
	return &pb.ContainSensitiveWordResponse{IsSensitiveWord: ok}, nil
}
