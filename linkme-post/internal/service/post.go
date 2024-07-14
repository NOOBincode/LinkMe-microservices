package service

import (
	"context"
	userpb "github.com/GoSimplicity/LinkMe-monorepo/api/user/v1"
	"github.com/golang/protobuf/ptypes/empty"
	pb "linkme-post/api/post/v1"
	"linkme-post/domain"
	"linkme-post/internal/biz"
)

type PostService struct {
	pb.UnimplementedPostServer
	userClient userpb.UserClient
	biz        *biz.PostBiz
}

func NewPostService(biz *biz.PostBiz, userClient userpb.UserClient) *PostService {
	return &PostService{
		biz:        biz,
		userClient: userClient,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.CreatePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	postId, err := s.biz.CreatePost(ctx, domain.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userId,
		Plate:   domain.Plate{ID: req.PlateId},
	})
	if err != nil {
		return &pb.CreatePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.CreatePostReply{
		Code: 0,
		Msg:  "create post success",
		Data: postId,
	}, nil
}
func (s *PostService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.UpdatePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	err = s.biz.UpdatePost(ctx, domain.Post{
		ID:      req.PostId,
		Title:   req.Title,
		Content: req.Content,
		UserID:  userId,
		Plate:   domain.Plate{ID: req.PlateId},
	})
	if err != nil {
		return &pb.UpdatePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.UpdatePostReply{
		Code: 0,
		Msg:  "update post success",
	}, nil
}
func (s *PostService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.DeletePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	err = s.biz.DeletePost(ctx, req.PostId, userId)
	if err != nil {
		return &pb.DeletePostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.DeletePostReply{
		Code: 0,
		Msg:  "delete post success",
	}, nil
}
func (s *PostService) PublishPost(ctx context.Context, req *pb.PublishPostRequest) (*pb.PublishPostReply, error) {
	// TODO 暂时保留
	return &pb.PublishPostReply{
		Code: 0,
		Msg:  "publish post success",
	}, nil
}
func (s *PostService) WithdrawPost(ctx context.Context, req *pb.WithdrawPostRequest) (*pb.WithdrawPostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.WithdrawPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	dp := domain.Post{
		ID:     req.PostId,
		UserID: userId,
		Status: domain.Withdrawn,
	}
	err = s.biz.UpdatePostStatus(ctx, dp)
	if err != nil {
		return &pb.WithdrawPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.WithdrawPostReply{
		Code: 0,
		Msg:  "withdraw post success",
	}, nil
}
func (s *PostService) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.ListPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	pagination := domain.Pagination{
		Page: int(req.Page),
		Size: &req.Size,
		Uid:  userId,
	}
	posts, err := s.biz.ListPost(ctx, pagination)
	if err != nil {
		return &pb.ListPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	pbPosts := make([]*pb.ListPost, len(posts))
	for i, post := range posts {
		pbPosts[i] = &pb.ListPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreateAt,
			UpdatedAt: post.UpdatedAt,
			UserId:    post.UserID,
			PlateId:   post.Plate.ID,
		}
	}
	return &pb.ListPostReply{
		Code: 0,
		Msg:  "list post success",
		Data: pbPosts,
	}, nil
}
func (s *PostService) ListPubPost(ctx context.Context, req *pb.ListPubPostRequest) (*pb.ListPubPostReply, error) {
	pagination := domain.Pagination{
		Page: int(req.Page),
		Size: &req.Size,
	}
	posts, err := s.biz.ListPubPost(ctx, pagination)
	if err != nil {
		return &pb.ListPubPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	pbPosts := make([]*pb.ListPost, len(posts))
	for i, post := range posts {
		pbPosts[i] = &pb.ListPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreateAt,
			UpdatedAt: post.UpdatedAt,
			UserId:    post.UserID,
			PlateId:   post.Plate.ID,
		}
	}
	return &pb.ListPubPostReply{
		Code: 0,
		Msg:  "list post success",
		Data: pbPosts,
	}, nil
}

func (s *PostService) ListAdminPost(ctx context.Context, req *pb.ListAdminPostRequest) (*pb.ListAdminPostReply, error) {
	// TODO 暂时保留
	return &pb.ListAdminPostReply{}, nil
}
func (s *PostService) DetailPost(ctx context.Context, req *pb.DetailPostRequest) (*pb.DetailPostReply, error) {
	userId, err := s.getUserId(ctx)
	if err != nil {
		return &pb.DetailPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	post, err := s.biz.GetPost(ctx, req.PostId, userId)
	if err != nil {
		return &pb.DetailPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.DetailPostReply{
		Code: 0,
		Msg:  "get post success",
		Data: &pb.DetailPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreateAt,
			UpdatedAt: post.UpdatedAt,
			UserId:    post.UserID,
			PlateId:   post.Plate.ID,
		},
	}, nil
}
func (s *PostService) DetailPubPost(ctx context.Context, req *pb.DetailPubPostRequest) (*pb.DetailPubPostReply, error) {
	post, err := s.biz.GetPubPost(ctx, req.PostId)
	if err != nil {
		return &pb.DetailPubPostReply{
			Code: 1,
			Msg:  err.Error(),
		}, err
	}
	return &pb.DetailPubPostReply{
		Code: 0,
		Msg:  "get pub post success",
		Data: &pb.DetailPost{
			Id:           post.ID,
			Title:        post.Title,
			Content:      post.Content,
			CreatedAt:    post.CreateAt,
			UpdatedAt:    post.UpdatedAt,
			UserId:       post.UserID,
			PlateId:      post.Plate.ID,
			LikeCount:    post.LikeNum,
			CollectCount: post.CollectNum,
			ViewCount:    post.ViewNum,
		},
	}, nil
}
func (s *PostService) DetailAdminPost(ctx context.Context, req *pb.DetailAdminPostRequest) (*pb.DetailAdminPostReply, error) {
	// TODO 暂时保留
	return &pb.DetailAdminPostReply{}, nil
}
func (s *PostService) GetPostStats(ctx context.Context, req *empty.Empty) (*pb.GetPostStatsReply, error) {
	// TODO 暂时保留
	return &pb.GetPostStatsReply{}, nil
}
func (s *PostService) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostReply, error) {
	// TODO 暂时保留
	return &pb.LikePostReply{}, nil
}
func (s *PostService) CollectPost(ctx context.Context, req *pb.CollectPostRequest) (*pb.CollectPostReply, error) {
	// TODO 暂时保留
	return &pb.CollectPostReply{}, nil
}

// 通过grpc调用linkme-user模块方法，获取userId
func (s *PostService) getUserId(ctx context.Context) (int64, error) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0NFpHOWJuQ2RUaTlCWGFVZHpoNmhpcFUxQll4S0daZiIsImV4cCI6MTcyMDk2MTQ3MCwiVWlkIjoyMTE0NDE5NDIzMTgzMjU3NiwiU3NpZCI6IjYyNTYyMzcyLThlZjMtNDAwYS05ZTYwLTllYWIzZTY4ZjQ3OSIsIlVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgU2FmYXJpLzUzNy4zNiIsIkNvbnRlbnRUeXBlIjoiYXBwbGljYXRpb24vanNvbiJ9.pHGkc8EzucCq2E6gjS3Q8u2eKwPxyd0ORPvgSP9DvTlRtNjnDMrAUvUkYfFpOiwvV9qnbZnN6bNFJPlX3YKovg"
	info, err := s.userClient.GetUserInfo(ctx, &userpb.GetUserInfoRequest{Token: token})
	if err != nil {
		return -1, err
	}
	return info.UserId, nil
}
