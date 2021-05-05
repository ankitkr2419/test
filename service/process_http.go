package service

import (
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

// This Handler will be used when we need entry inside only processes table 
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

		processID, err := parseUUID(vars["process_id"])
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ProcessIDInvalidError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ProcessIDInvalidError.Error()})
			return
		}

		process, err := deps.Store.DuplicateProcess(req.Context(), processID)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ProcessDuplicationError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessDuplicationError.Error()})
			return
		}

		logger.Infoln(responses.ProcessDuplicationSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, process)
	})
}

func rearrangeProcessesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		recipeID, err := parseUUID(vars["recipe_id"])
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.RecipeIDInvalidError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RecipeIDInvalidError.Error()})
		}

		var sequenceArr []db.ProcessSequence

		err = json.NewDecoder(req.Body).Decode(&sequenceArr)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ProcessesDecodeSeqError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ProcessesDecodeSeqError.Error()})
			return
		}

		logger.Infoln("Sequence Array: ", sequenceArr)

		processes, err := deps.Store.RearrangeProcesses(req.Context(), recipeID, sequenceArr)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ProcessesRearrangeError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessesRearrangeError.Error()})
			return
		}

		logger.Infoln(responses.ProcessesRearrangeSuccess)
		responseCodeAndMsg(rw, http.StatusOK, processes)
	})
}
