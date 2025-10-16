package services

import (
	"context"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type cloudinaryService struct {
	cld *cloudinary.Cloudinary
}

var cldSvc *cloudinaryService

func NewCloudinaryService() (*cloudinaryService, error) {
	if cldSvc != nil {
		return cldSvc, nil
	}

	cldUtil, err := cloudinary.New()
	if err != nil {
		return nil, err
	}

	cldSvc = &cloudinaryService{
		cld: cldUtil,
	}

	return cldSvc, nil

}

func (svc *cloudinaryService) UploadPdf(file *multipart.FileHeader) (string, string, error) {
	f, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	bytesData, err := ioutil.ReadAll(f)
	if err != nil {
		return "", "", err
	}

	resp, err := svc.cld.Upload.Upload(context.Background(), bytesData, uploader.UploadParams{
		ResourceType: "raw",
	})

	if err != nil {
		return "", "", err
	}

	return resp.SecureURL, resp.PublicID, nil
}

func (svc *cloudinaryService) GetFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (svc *cloudinaryService) DeleteFile(publicID string) error {
	_, err := svc.cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: "raw",
	})
	return err
}
