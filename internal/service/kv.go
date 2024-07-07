package service

import (
	"context"

	pb "leegoo/api/kvstore/v1"
	"leegoo/internal/data"
)

type KvService struct {
	pb.UnimplementedKvServer
	kvData *data.KvData
}

func NewKvService(kvData *data.KvData) *KvService {
	return &KvService{kvData: kvData}
}

func (s *KvService) Get(ctx context.Context, req *pb.KvRequest) (*pb.KvReply, error) {
	data, err := s.kvData.GetKv(req.GetKey())
	var code int32 = 1
	if err != nil {
		code = 999
	}
	return &pb.KvReply{Code: code, Data: data}, err
}
func (s *KvService) Del(ctx context.Context, req *pb.KvRequest) (*pb.KvReply, error) {
	err := s.kvData.DelKv(req.GetKey())
	var code int32 = 1
	if err != nil {
		code = 999
	}
	return &pb.KvReply{Code: code}, err
}

func (s *KvService) Set(ctx context.Context, req *pb.KvRequest) (*pb.KvReply, error) {
	err := s.kvData.SetKv(req.GetKey(), req.GetData())
	var code int32 = 1
	if err != nil {
		code = 999
	}
	return &pb.KvReply{Code: code}, err
}
