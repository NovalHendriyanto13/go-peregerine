package requests

type UploadCreateRequest struct {
	Resume string `json:"resume"`
}

type UploadResumeJob struct {
	Resume string `json:"file_path"`
	JobDesc string `json:"job_desc"`
}