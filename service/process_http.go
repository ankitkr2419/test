package service

import (
	"io/ioutil"
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func listProcessesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		list, err := deps.Store.ListProcesses(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		respBytes, err := json.Marshal(list)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling process data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func createProcessHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var process db.Process
		err := json.NewDecoder(req.Body).Decode(&process)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding process data")
			return
		}

		valid, respBytes := validate(process)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		var createdTemp db.Process
		createdTemp, err = deps.Store.CreateProcess(req.Context(), process)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create process")
			return
		}

		respBytes, err = json.Marshal(createdTemp)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling process data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}

func showProcessHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var latestT db.Process

		latestT, err = deps.Store.ShowProcess(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show process")
			return
		}

		respBytes, err := json.Marshal(latestT)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling process data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func deleteProcessHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = deps.Store.DeleteProcess(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while deleting process")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"process deleted successfully"}`))
	})
}

func updateProcessHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var process db.Process

		err = json.NewDecoder(req.Body).Decode(&process)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding process data")
			return
		}

		valid, respBytes := validate(process)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		process.ID = id
		err = deps.Store.UpdateProcess(req.Context(), process)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error update process")
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"process updated successfully"}`))
	})
}

func duplicateProcessHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		var process interface{}

		processID, err := parseUUID(vars["process_id"])
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ProcessIDInvalidError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ProcessIDInvalidError.Error()})
			return
		}

		processType := vars["process_type"]

		processBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ProcessDecodeError.Error()})
			logger.WithField("err", err.Error()).Error(responses.ProcessDecodeError.Error())
			return
		}

		switch processType{
		case "Piercing":
			process = &db.Piercing{}
		case "TipOperation":
			process = &db.TipOperation{}
		case "TipDocking":
			process = &db.TipDock{}
		case "AspireDispense":
			process = &db.AspireDispense{}
		case "Heating":
			process = &db.Heating{}
		case "Shaking":
			process = &db.Shaker{}
		case "AttachDetach":
			process = &db.AttachDetach{}
		case "Delay":
			process = &db.Delay{}
		default:
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ProcessTypeInvalid.Error()})
			logger.Errorln(responses.ProcessTypeInvalid.Error())
			return
		}

		err = json.Unmarshal(processBytes, &process)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ProcessDecodeError.Error()})
			logger.WithField("err", err.Error()).Error(responses.ProcessDecodeError.Error())
			return
		}

		valid, respBytes := validate(process)
		if !valid {
			logger.Errorln(responses.ProcessValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		process, err = deps.Store.DuplicateProcess(req.Context(), processID, process)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.DuplicateProcessCreationError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DuplicateProcessCreationError.Error()})
			return
		}

		logger.Infoln(responses.DuplicateProcessCreationSuccess)
		responseCodeAndMsg(rw, http.StatusOK, process)
	})
}
