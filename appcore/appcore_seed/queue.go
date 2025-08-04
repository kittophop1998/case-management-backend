package appcore_seed

import (
	"case-management/model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedQueue(db *gorm.DB) map[string]uuid.UUID {
	queueMap := make(map[string]uuid.UUID)

	queues := []model.Queue{
		{Name: "Inbound Agent Supervisor Big C Queue"},
		{Name: "Inbound Agent Supervisor Queue"},
		{Name: "Case Handling Team Queue"},
		{Name: "CCAgent"},
		{Name: "Case Handling Team Supervisor Queue"},
		{Name: "Fraud Operation"},
		{Name: "Fraud Supervisor"},
		{Name: "Manager Up"},
		{Name: "Chargeback Supervisor"},
		{Name: "Inbound Agent Supervisor BKK Queue"},
		{Name: "Inbound Agent Supervisor CM Queue"},
		{Name: "Convince Supervisor"},
		{Name: "Chargeback Team"},
		{Name: "EDP Team Group"},
		{Name: "Convince Team"},
		{Name: "Inbound Agent Supervisor HY Queue"},
		{Name: "Authorize Supervisor"},
		{Name: "Fraud Manager"},
		{Name: "Tele Motor Team"},
		{Name: "Authorize Operation"},
		{Name: "Inbound Agent Supervisor KK Queue"},
		{Name: "Agent Telemarketing"},
		{Name: "Supervisor Telemarketing"},
	}

	for _, q := range queues {
		var queue model.Queue
		err := db.Where("name = ?", q.Name).FirstOrCreate(&queue, q).Error
		if err != nil {
			log.Printf("failed to seed queue %s: %v", q.Name, err)
		} else {
			queueMap[q.Name] = queue.ID
		}
	}

	return queueMap
}
