package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

//ErrorResponse stores the error message.
type ErrorResponse struct {
	Message string `json:"message"`
}

//Handler handles rest api events for Storage.
type Handler struct {
	storage Storage
}

//NewHandler constructs the handler object.
func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

//CheckImage returns true if uploading file is an image.
func CheckImage(file *multipart.FileHeader) (bool, error) {
	fileContent, _ := file.Open()
	byteContainer, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return false, err
	}
	return strings.Contains(http.DetectContentType(byteContainer), "image"), nil
}

//FileNameWithoutExtension returns filename without extension.
func FileNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}

//GeneratePreview creates the thumbnail of uploaded file.
func GeneratePreview(filename string, ch chan string) error {
	imagePath, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer imagePath.Close()
	srcImage, _, err := image.Decode(imagePath)
	if err != nil {
		return err
	}
	dstImage := imaging.Thumbnail(srcImage, 80, 80, imaging.Lanczos)

	newFileName := FileNameWithoutExtension(filename) + "thumbnail.jpg"
	ch <- newFileName
	newImage, err := os.Create(newFileName)
	if err != nil {
		return err
	}
	defer newImage.Close()
	jpeg.Encode(newImage, dstImage, &jpeg.Options{jpeg.DefaultQuality})
	return nil
}

//AddPhoto describes rest api POST method.
//It saves photo at local machine and calls gallery's insert method.
func (h *Handler) AddPhoto(c *gin.Context) {
	var photo Photo

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("get form err: %s", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	isImage, err := CheckImage(file)
	if err != nil {
		fmt.Printf("Error while reading file: %s", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if !isImage {
		fmt.Println("The type of the downloaded file does not match the image")
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "The type of the downloaded file does not match the image",
		})
		return
	}

	path := Configuration.GalleryPath() + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		fmt.Printf("Error while saving file: %s", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	ch := make(chan string)
	go GeneratePreview(path, ch)

	photo.Path = path
	photo.Preview = &Photo{Path: <-ch}
	photo.SetName()
	photo.Preview.SetName()

	if err != nil {
		fmt.Printf("Error while saving file: %s", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Insert(&photo)

	c.JSON(http.StatusOK, map[string]string{
		"Message": fmt.Sprintf("File %s uploaded successfully.", file.Filename),
	})
}

//GetPhotos describes rest api GET method.
//It returns the slice of all photos.
func (h *Handler) GetPhotos(c *gin.Context) {
	photos := h.storage.Get()
	resPhotos := make([]map[int]string, 0)
	//Cheks does the information in the database corresponds to reality.
	for _, photo := range photos {
		if _, err := os.Stat(photo.Path); os.IsNotExist(err) {
			fmt.Println("File ", photo.Name, " is not exist in ", photo.Path)
			continue
		}
		photoMap := map[int]string{photo.ID: photo.Name}
		resPhotos = append(resPhotos, photoMap)
	}
	c.JSON(http.StatusOK, resPhotos)
}

//DeletePhoto describes rest api DELETE method.
//It deletes photo from db and local storage by its id.
func (h *Handler) DeletePhoto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	path := h.storage.Delete(id)

	//Check and delete photo from local storage.
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err = os.Remove(path)
		if err != nil {
			fmt.Printf("failed to delete file: %s\n", err.Error())
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})
			return
		}
	}

	previewPath := FileNameWithoutExtension(path) + "thumbnail.jpg"

	//Check and delete preview from local storage.
	if _, err := os.Stat(previewPath); !os.IsNotExist(err) {
		err = os.Remove(previewPath)
		if err != nil {
			fmt.Printf("failed to delete file: %s\n", err.Error())
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})
			return
		}
	}

	DeletePreview(previewPath)

	c.JSON(http.StatusOK, map[string]string{
		"Message": fmt.Sprint("photo deleted"),
	})
}
