package api

import (
	"encoding/json"

	"go-pre-interview/model"
)

type apiResponse struct {
	OK        bool             `json:"ok"`
	ReplyTO   int              `json:"reply_to"`
	TestTable *model.TestTable `json:"test"`
}

type CreateRequest struct {
	Name         string `json:"name"`
	RandomString string `json:"random_string"`
}

type broadcastRequest struct {
	ID int `json:"test_id"`
}

type Message struct {
	Type      string `json:"type"`
	RequestId int    `json:"requestId"`
	Test      int    `json:"test"`
}

func (api *API) CreateHandler(session Session, requestID int, payload []byte) error {
	var request CreateRequest
	err := json.Unmarshal(payload, &request)
	if err != nil {
		return err
	}

	test := api.datastore.TestTableStore.Create(request.Name, request.RandomString)

	session.SetID(test.ID)

	response := apiResponse{true, requestID, test}
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}
	session.Send(data)

	return nil
}

type infoRequest struct {
	TestId int `json:"test_id"`
}

func (api *API) infoHandler(session Session, requestID int, payload []byte) error {
	var request infoRequest
	err := json.Unmarshal(payload, &request)
	if err != nil {
		return errInvalidRequest
	}

	if request.TestId == 0 {
		return errMissingID
	}

	test := api.datastore.TestTableStore.GetByID(request.TestId)
	if test == nil {
		return errTestNotFound
	}

	response := apiResponse{true, requestID, test}
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}
	session.Send(data)

	return nil
}

func newMessage(sessionId, testId int) []byte {
	msg := Message{"message", sessionId, testId}
	data, err := json.Marshal(msg)
	if err != nil {
		panic("unable to marshal json")
	}
	return data
}

func (api *API) broadcastHandler(session Session, requestID int, payload []byte) error {
	var request broadcastRequest
	err := json.Unmarshal(payload, &request)
	if err != nil {
		return errInvalidRequest
	}

	if requestID == 0 {
		return errMissingID
	}

	test := api.datastore.TestTableStore.GetByID(request.ID)
	if test == nil {
		return errTestNotFound
	}

	session.Subscribe(request.ID, session.ID())
	session.Broadcast(request.ID, newMessage(session.ID(), request.ID))

	response := apiResponse{OK: true, ReplyTO: requestID, TestTable: test}
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	session.Send(data)

	return nil
}
