package domain

type UploadedResponse struct {
	Upload       []StorageCreatedFile `json:"upload"`
	InvalidFiles []InvalidFile        `json:"invalidFile"`
}

type StorageCreatedFile struct {
	FullPath string `json:"fullPath"`
	Path     string `json:"path"`
	Folder   string `json:"folder"`
	Name     string `json:"name"`
	Mime     string `json:"mime"`
}

type InvalidFile struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
