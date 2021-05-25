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
			logger.WithField("err", err.Error()).Errorln(responses.UUIDParseError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		list, err := deps.Store.ListProcesses(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ProcessFetchError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessFetchError.Error()})
			return
		}

		logger.Infoln(responses.ProcessesFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, list)
	})
}

// This Handler will be used when we need entry inside only processes table
func createProcessHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var process db.Process
		err := json.NewDecoder(req.Body).Decode(&process)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ProcessDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ProcessDecodeError.Error()})
			return
		}

		valid, respBytes := validate(process)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.ProcessValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		var createdTemp db.Process
		createdTemp, err = deps.Store.CreateProcess(req.Context(), process)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ProcessCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessCreateError.Error()})
			return
		}

		logger.Infoln(responses.ProcessCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, createdTemp)
	})
}

func showProcessHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var latestT db.Process

		latestT, err = deps.Store.ShowProcess(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessFetchError.Error()})
			logger.WithField("err", err.Error()).Errorln(responses.ProcessFetchError)
			return
		}

		logger.Infoln(responses.ProcessFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, latestT)
	})
}

func deleteProcessHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		err = deps.Store.DeleteProcess(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessDeleteError.Error()})
			logger.WithField("err", err.Error()).Error(responses.ProcessDeleteError)
			return
		}

		logger.Infoln(responses.ProcessDeleteSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.ProcessDeleteSuccess})
	})
}

func updateProcessHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var process db.Process

		err = json.NewDecoder(req.Body).Decode(&process)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ProcessDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ProcessDecodeError.Error()})
			return
		}

		valid, respBytes := validate(process)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.ProcessValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		process.ID = id

		err = deps.Store.UpdateProcess(req.Context(), process)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessUpdateError.Error()})
			logger.WithField("err", err.Error()).Error(responses.ProcessUpdateError)
			return
		}

		logger.Infoln(responses.ProcessUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.ProcessUpdateSuccess})
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
			return
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
