package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FaceService struct {
	BaseURL string
}

type EmbeddingResponse struct {
	Success   bool      `json:"success"`
	Error     string    `json:"error,omitempty"`
	Embedding []float32 `json:"embedding"`
}

type DetectionResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"error,omitempty"`
	Faces   []FaceData `json:"faces"`
}

type FaceData struct {
	BBox       [4]float32 `json:"bbox"`
	Confidence float32    `json:"confidence"`
	Embedding  []float32  `json:"embedding"`
}

func NewFaceService(baseURL string) *FaceService {
	return &FaceService{BaseURL: baseURL}
}

func (fs *FaceService) GenerateEmbedding(imagePath string) ([]float32, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(imagePath))
	if err != nil {
		return nil, err
	}

	io.Copy(part, file)
	writer.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/embedding", fs.BaseURL), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result EmbeddingResponse
	json.NewDecoder(resp.Body).Decode(&result)

	if !result.Success {
		return nil, fmt.Errorf("embedding generation failed: %s", result.Error)
	}

	return result.Embedding, nil
}

func (fs *FaceService) DetectFaces(imagePath string) ([]FaceData, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(imagePath))
	if err != nil {
		return nil, err
	}

	io.Copy(part, file)
	writer.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/detect", fs.BaseURL), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result DetectionResponse
	json.NewDecoder(resp.Body).Decode(&result)

	if !result.Success {
		return nil, fmt.Errorf("face detection failed: %s", result.Error)
	}

	return result.Faces, nil
}

func (fs *FaceService) CosineSimilarity(emb1, emb2 []float32) float32 {
	if len(emb1) != len(emb2) {
		return 0
	}

	var dotProduct, norm1, norm2 float32
	for i := 0; i < len(emb1); i++ {
		dotProduct += emb1[i] * emb2[i]
		norm1 += emb1[i] * emb1[i]
		norm2 += emb2[i] * emb2[i]
	}

	if norm1 == 0 || norm2 == 0 {
		return 0
	}

	if norm1*norm2 == 0 {
		return 0
	}

	return dotProduct / (norm1 * norm2)
}
