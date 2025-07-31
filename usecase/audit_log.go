package usecase

import (
	"case-management/model"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func (u *UseCase) CreateAuditLog(c *gin.Context, userID uuid.UUID, eventType, tableId, changeTable string, newStruct interface{}) error {

	changesJson, err := convertToDataType(newStruct)
	if err != nil {
		return err
	}

	log := model.AuditLog{
		EventType:   eventType,
		ChangeTable: changeTable,
		ChangeId:    tableId,
		Diff:        changesJson,
		UserID:      userID,
	}
	if err := u.caseManagementRepository.CreateAuditLog(c, log); err != nil {
		return err
	}

	return nil
}

func convertToDataType(value interface{}) (datatypes.JSON, error) {
	toJson, err := json.Marshal(value)
	if err != nil {
		return datatypes.JSON{}, err
	}

	toMap := map[string]interface{}{}
	json.Unmarshal([]byte(string(toJson)), &toMap)

	for key, value := range toMap {
		if value == nil || value == "" || key == "CreatedAt" || key == "UpdatedAt" || key == "DeletedAt" || key == "ID" {
			delete(toMap, key)
		} else {
			if arr, ok := value.([]interface{}); ok {
				for _, item := range arr {
					if itemMap, itemOk := item.(map[string]interface{}); itemOk {
						for k, v := range itemMap {
							if v == nil || v == "" || k == "CreatedAt" || k == "UpdatedAt" || k == "DeletedAt" {
								delete(itemMap, k)
							}
						}
					}
				}
			}
		}
	}

	jsonData, err := json.Marshal(toMap)
	if err != nil {
		return datatypes.JSON{}, err
	}

	var dataTypeJson datatypes.JSON
	err = json.Unmarshal(jsonData, &dataTypeJson)
	if err != nil {
		return datatypes.JSON{}, err
	}

	return dataTypeJson, nil
}
