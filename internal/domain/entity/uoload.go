package entity

type UpLoadFile struct {
	FileTag string `json:"file_tag" binding:"required"`
}

type FileInfo struct {
	FileURl   string `json:"file_url"`
	ObjectKey string `json:"object_key"`
	FileName  string `json:"file_name"`
	Size      int64  `json:"size"`
}
