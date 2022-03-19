package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) UnFollowUser(ctx context.Context, req *v1.UnFollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}


func (s *RealWorldService) ListArticle(ctx context.Context, req *v1.ListArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) FeedArticle(ctx context.Context, req *v1.FeedArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, req *v1.GetArticlesRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (reply *v1.SingleCommentReply, err error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *RealWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (reply *v1.SingleCommentReply, err error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, req *v1.UnFavoriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (reply *v1.TagListReply, err error) {
	return &v1.TagListReply{}, nil
}