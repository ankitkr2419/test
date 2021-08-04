package service

import (
	"fmt"
	"mylab/cpagent/responses"
	"strings"

	"encoding/base64"
	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"mylab/cpagent/config"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	expOutputPath = "./utils/output"
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

		err = EmailReport(expID)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "Email couldn't be sent: " + err.Error()})
		}
		responseCodeAndMsg(rw, http.StatusAccepted, MsgObj{Msg: "Email sent success"})
	})
}

func EmailReport(experimentID uuid.UUID) (err error) {

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

	// TODO: Uncomment below code once @ankush's pdf is put in report folder
	// a_pdf := mail.NewAttachment()
	// dat, err := ioutil.ReadFile("/home/josh/Desktop/Programs/GO/PDF/hello.pdf")
	// if err != nil{
	// 	logger.Errorln(err)
	// 	return
	// }

	// encoded := base64.StdEncoding.EncodeToString([]byte(dat))
	// a_pdf.SetContent(encoded)
	// a_pdf.SetType("application/pdf")
	// a_pdf.SetFilename("Experiment_Report.pdf")
	// a_pdf.SetDisposition("attachment")

	// attach a xlsx report
	a_xlsx := mail.NewAttachment()

	files, err := ioutil.ReadDir(expOutputPath)
	if err != nil {
		logger.Errorln("Failed to read files from ", expOutputPath, err)
	}

	var fileName string

	for _, f := range files {
		fName := strings.SplitN(f.Name(), "_", 3)
		if fName[1] == experimentID.String() {
			fileName = f.Name()
			goto fileFound
		}
	}
	return responses.ExperimentFetchError

fileFound:
	dat, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", expOutputPath, fileName))
	if err != nil {
		logger.Errorln(err)
		return
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(dat))
	a_xlsx.SetContent(encoded)
	a_xlsx.SetType("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	a_xlsx.SetFilename(fmt.Sprintf("Experiment_%v.xlsx", experimentID))
	a_xlsx.SetDisposition("attachment")

	// m.AddAttachment(a_pdf)
	m.AddAttachment(a_xlsx)

	request := sendgrid.GetRequest(config.GetSendGridAPIKey(), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
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
