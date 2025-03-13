package models

// FileType represents the type of file being uploaded
type FileType int

const (
	// FileTypeCourses represents a courses file
	FileTypeCourses FileType = iota
	// FileTypeEquivalencies represents an equivalencies file
	FileTypeEquivalencies
	// FileTypeStudents represents a students file
	FileTypeStudents
	// FileTypeStudentCourses represents a student courses file
	FileTypeStudentCourses
)

// String returns the string representation of a FileType
func (ft FileType) String() string {
	switch ft {
	case FileTypeCourses:
		return "courses"
	case FileTypeEquivalencies:
		return "equivalencies"
	case FileTypeStudents:
		return "students"
	case FileTypeStudentCourses:
		return "studentcourses"
	default:
		return "unknown"
	}
}

// UploadFile represents a file to be uploaded
type UploadFile struct {
	Type     FileType
	FilePath string
}

// UploadRequest represents a request to upload files
type UploadRequest struct {
	APIKey string
	Files  []UploadFile
}

// UploadResponse represents the response from an upload request
type UploadResponse struct {
	Success bool
	Message string
	Code    int
}
