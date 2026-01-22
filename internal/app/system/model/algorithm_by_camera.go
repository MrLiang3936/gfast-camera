package model

import "github.com/tiger1103/gfast/v3/internal/app/system/model/entity"

type AlgorithmByCamera struct {
	*entity.SysAlgorithm
	CameraList []*entity.SysCamera `json:"cameraList"`
}
