package data

import (
	"kratos-realworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewOSSStore, NewUserRepo, NewProfileRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	oss *minio.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, oss *minio.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, oss: oss}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:dangerous@tcp(127.0.0.1:3306)/realworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
	return db
}

func NewOSSStore(c *conf.Data) *minio.Client {
	// ctx := context.Background()
	endpoint := "localhost:19000"
	accessKeyID := "admin"
	secretAccessKey := "admin123456"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Debug(err)
		return nil
	}

	log.Debug("%#v\n", minioClient) // minioClient is now set up
	log.Debug("SUCCESSFULY CREATED minioClient")
	return minioClient
}
