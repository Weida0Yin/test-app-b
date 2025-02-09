package upload

import (
	"github.com/gin-gonic/gin"
	"novel-app/internal/domain/entity"
	"novel-app/internal/svc"
	"novel-app/pkg/response"
)

type UploadHandler struct {
	service *svc.UploadService
}

func NewUploadHandler(e *svc.UploadService) *UploadHandler {
	return &UploadHandler{service: e}
}

// UploadFile 文件上传
func (h *UploadHandler) UploadFile(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		response.Fail(c, "Unauthorized")
		return
	}

	//req := entity.UpLoadFile{}
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	response.Fail(c, "Invalid input")
	//	return
	//}

	if userId, ok := userId.(int64); ok {
		// 2. 获取上传文件
		fileHeader, err := c.FormFile("file")
		if err != nil {
			response.Fail(c, "please input File")
			//ctx.JSON(http.StatusBadRequest, gin.H{"error": "请选择上传文件"})
			return
		}

		// 3. 验证文件类型
		//ext := filepath.Ext(fileHeader.Filename)
		//if !isValidFileType(ext) {
		//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "不支持的文件类型"})
		//	return
		//}

		// 4. 打开文件流
		file, err := fileHeader.Open()
		if err != nil {
			response.Fail(c, "read File error")
			//ctx.JSON(http.StatusInternalServerError, gin.H{"error": "文件读取失败"})
			return
		}
		defer file.Close()

		// 5. 生成OSS对象名称
		objectKey := h.service.GenerateObjectKey("avatar", userId, fileHeader.Filename)

		// 6. 上传到OSS
		fileURL, err := h.service.UploadFile(objectKey, file)
		if err != nil {
			response.Fail(c, "upload file error")
			//ctx.JSON(http.StatusInternalServerError, gin.H{"error": "文件上传失败"})
			return
		}

		resp := entity.FileInfo{
			FileName:  fileHeader.Filename,
			FileURl:   fileURL,
			ObjectKey: objectKey,
			Size:      fileHeader.Size,
		}

		// 7. 返回结果
		response.Success(c, resp, "success")
		return

	}

	response.Fail(c, "upload fail")
	return
}
