package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type CreditFile struct {
	ID                    string
	Owner                 string
	Academy               string
	DisciplinarySituation string
	GradePoint            float32
	AwardSituation        string
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	creditFile := CreditFile{
		ID:                    "2020131102",
		Owner:                 "hejiean",
		Academy:               "区块链产业学院",
		DisciplinarySituation: "无",
		GradePoint:            3.3,
		AwardSituation:        "学业优良奖",
	}

	fileJSON, err := json.Marshal(creditFile)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(creditFile.ID, fileJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	return nil
}

func (s *SmartContract) CreditFileExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	fileJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return fileJSON != nil, nil
}

func (s *SmartContract) CreateCreditFile(
	ctx contractapi.TransactionContextInterface,
	id string,
	owner string,
	academy string,
	disciplinarySituation string,
	gradePoint float32,
	awardSituation string) error {
	exists, err := s.CreditFileExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the creditFile %s already exists", id)
	}

	creditFile := CreditFile{
		ID:                    id,
		Owner:                 owner,
		Academy:               academy,
		DisciplinarySituation: disciplinarySituation,
		GradePoint:            gradePoint,
		AwardSituation:        awardSituation,
	}
	fileJSON, err := json.Marshal(creditFile)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, fileJSON)
}

func (s *SmartContract) GetCreditFile(ctx contractapi.TransactionContextInterface, id string) (*CreditFile, error) {
	fileJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if fileJSON == nil {
		return nil, fmt.Errorf("the creditFile %s does not exist", id)
	}

	var file CreditFile
	err = json.Unmarshal(fileJSON, &file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (s *SmartContract) UpdateCreditFile(
	ctx contractapi.TransactionContextInterface,
	id string,
	owner string,
	academy string,
	disciplinarySituation string,
	gradePoint float32,
	awardSituation string) error {
	exists, err := s.CreditFileExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the creditFile %s does not exist", id)
	}

	creditFile := CreditFile{
		ID:                    id,
		Owner:                 owner,
		Academy:               academy,
		DisciplinarySituation: disciplinarySituation,
		GradePoint:            gradePoint,
		AwardSituation:        awardSituation,
	}
	fileJSON, err := json.Marshal(creditFile)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, fileJSON)
}

func (s *SmartContract) DeleteCreditFile(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.CreditFileExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the creditFile %s does not exist", id)
	}
	return ctx.GetStub().DelState(id)
}
