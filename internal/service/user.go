package service

import (
	"context"
	"fmt"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc  *biz.UserUsecase
	oss *minio.Client
	log *log.Helper
}

func (u *RealWorldService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 34)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	fmt.Println("filesize: ", handler.Size)
	timestamp := time.Unix(time.Now().Unix(), 0)

	_, err = u.oss.PutObject(context.Background(), "materials", timestamp.String()+handler.Filename,
		file, handler.Size, minio.PutObjectOptions{ContentType: "application/octet-stream", PartSize: 5 * 1024 * 1024})

	if err != nil {
		fmt.Println(err)
		return
	}

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func (u *RealWorldService) RegisterHTTP(ctx context.Context, endpoint string, minioClient *minio.Client) (http.Handler, error) {
	mux := http.NewServeMux()
	// Add Upload
	mux.Handle("/upload", u)
	return mux, nil
}

func (s *RealWorldService) GetStoreClient(ctx context.Context) *minio.Client {
	return s.uc.GetStoreClient(ctx)
}

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "ALEX",
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "ALEX",
		},
	}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}

func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}

func (s *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}
