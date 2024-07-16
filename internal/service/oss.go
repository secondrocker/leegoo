package service

import (
	"context"

	pb "leegoo/api/oss/v1"
	"leegoo/internal/data"
)

type OssService struct {
	pb.UnimplementedOssServer
	ossData *data.OssData
}

func NewOssService(ossData *data.OssData) *OssService {
	return &OssService{ossData: ossData}
}

func (s *OssService) GetUrl(ctx context.Context, req *pb.OssRequest) (*pb.OssReply, error) {
	url, err := s.ossData.GetUrl(req.GetBucket(), req.GetPath())
	var code int32 = 1
	if err != nil {
		code = 999
	}
	return &pb.OssReply{Code: code, Url: url}, err
}
func (s *OssService) Delete(ctx context.Context, req *pb.OssRequest) (*pb.OssReply, error) {
	err := s.ossData.Delete(req.GetBucket(), req.GetPath())
	var code int32 = 1
	if err != nil {
		code = 999
	}
	return &pb.OssReply{Code: code}, err
}
func (s *OssService) Restore(ctx context.Context, req *pb.OssRequest) (*pb.OssReply, error) {
	url, err := s.ossData.Restore(req.GetBucket(), req.GetPath())
	var code int32 = 1
	if err != nil {
		code = 999
	}
	return &pb.OssReply{Code: code, Url: url}, err
}
