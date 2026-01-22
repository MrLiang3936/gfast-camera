package model

import "github.com/tiger1103/gfast/v3/internal/app/system/model/entity"

type CameraByAlgorithm struct {
	*entity.SysCamera
	AlgorithmList []*entity.SysAlgorithm `json:"algorithmList"`
}
