package service

import (
	"fmt"
	"mylab/cpagent/responses"
	"strings"
	"os"

	"encoding/base64"
	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"mime/multipart"
	"mylab/cpagent/config"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const(
	pdf              = "pdf"
	xlsx             = "xlsx"
)

func emailReport(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			logger.Errorln("Invalid Experiment ID: ", expID)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.InvalidExperimentID.Error()})
			return
		}

		err = emailTheReport(expID)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "Email couldn't be sent: " + err.Error()})
			return
		}
		responseCodeAndMsg(rw, http.StatusAccepted, MsgObj{Msg: "Email sent success"})
	})
}

func uploadReport(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// Parse input, multipart/form-data
		err := req.ParseMultipartForm(15 << 20) // 15 MB Max File Size
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while parsing the Report form")
			Message := "Invalid Form Data! Error while parsing the Report form"
			responseCodeAndMsg(rw, http.StatusBadRequest, Message)
			return
		}

		vars := mux.Vars(req)
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			logger.Errorln("Invalid Experiment ID: ", expID)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.InvalidExperimentID.Error()})
			return
		}

		formdata := req.MultipartForm
		report := formdata.File["report"]
		if report == nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ReportAbsent.Error()})
			return
		}

		err = uploadTheReport(expID, report)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "Report couldn't be uploaded: " + err.Error()})
			return
		}
		responseCodeAndMsg(rw, http.StatusAccepted, MsgObj{Msg: "Report uploaded successfully"})
	})
}

func uploadTheReport(expID uuid.UUID, report []*multipart.FileHeader) error {

	reportPDF, err := report[0].Open()
	defer reportPDF.Close()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while decoding report Data, probably invalid report")
		return err
	}

	extension := strings.Split(report[0].Filename, ".")
	if extension[len(extension)-1] != pdf {
		err = fmt.Errorf("Incorrect extension of file!")
		logger.WithField("err", err.Error()).Error("Error while getting report Extension. Re-check the report file extension!")

		return err
	}

	tempFile, err := ioutil.TempFile(os.ExpandEnv(reportOutputPath), expID.String()+"."+pdf)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while Creating a Temporary File")
		return err
	}
	defer tempFile.Close()

	imageBytes, err := ioutil.ReadAll(reportPDF)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while reading report File")
		return err
	}
	tempFile.Write(imageBytes)
	logger.Infoln("Filename : ", tempFile.Name())

	err = os.Chmod(tempFile.Name(), 0766)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while changing File permission")
		return err
	}

	os.Rename(tempFile.Name(),os.ExpandEnv(reportOutputPath) + "/"+ expID.String()+"."+pdf )

	return nil
}

func emailTheReport(experimentID uuid.UUID) (err error) {

	var dat []byte
	var response *rest.Response
	m := mail.NewV3Mail()

	from := mail.NewEmail("MyLab Software", "mylab.software5656@gmail.com")
	plainTextContent := mail.NewContent("text/plain", "Test mail")
	to := mail.NewEmail(config.GetReceiverName(), config.GetReceiverEmail())
	subject := "Sending sample Report Attachment"

	m.SetFrom(from)
	m.AddContent(plainTextContent)

	// create new Personalization
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.Subject = subject

	m.AddPersonalizations(personalization)

	// attach PDF

	a_pdf := mail.NewAttachment()
	dat, err = ioutil.ReadFile(fmt.Sprintf("%v/%v.%v", os.ExpandEnv(reportOutputPath), experimentID.String(), pdf))
	if err != nil {
		logger.Errorln(err)
		return
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(dat))
	a_pdf.SetContent(encoded)
	a_pdf.SetType("application/pdf")
	a_pdf.SetFilename(fmt.Sprintf("Experiment_%v.%v", experimentID.String(), pdf))
	a_pdf.SetDisposition("attachment")

	a_xlsx := mail.NewAttachment()


	fileName := fmt.Sprintf("%v/output_%v.%v", os.ExpandEnv(expOutputPath), experimentID.String(), xlsx)
	dat, err = ioutil.ReadFile(fmt.Sprintf(fileName))
	if err != nil {
		logger.Errorln(err)
		return
	}

	encoded = base64.StdEncoding.EncodeToString([]byte(dat))
	a_xlsx.SetContent(encoded)
	a_xlsx.SetType("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	a_xlsx.SetFilename(fmt.Sprintf("Experiment_%v.xlsx", experimentID))
	a_xlsx.SetDisposition("attachment")

	m.AddAttachment(a_pdf)
	m.AddAttachment(a_xlsx)

	request := sendgrid.GetRequest(config.GetSendGridAPIKey(), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err = sendgrid.API(request)
	if err != nil {
		logger.Errorln(err)
		return
	} else {
		logger.WithFields(logger.Fields{
			"Status Code:":      response.StatusCode,
			"Response Body:":    response.Body,
			"Response Headers:": response.Headers,
		}).Infoln("Email sent")
	}
	return
}
