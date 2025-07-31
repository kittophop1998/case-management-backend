package appcore_storage

import (
	"case-management/appcore/appcore_config"

	"github.com/minio/minio-go"
)

var Storage *minio.Client

func InitStorage() {
	var err error
	Storage, err = minio.New(appcore_config.Config.MinioURL, appcore_config.Config.MinioAccessKey, appcore_config.Config.MinioSecretKey, appcore_config.Config.MinioSSL)
	if err != nil {
		panic("cannot connect storage (minio) >> " + err.Error())
	}
}
