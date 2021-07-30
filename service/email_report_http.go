package service

import (
	"fmt"

	"encoding/base64"
	"io/ioutil"
	"net/http"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/config"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func emailReport(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		err := EmailReport()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "Email couldn't be sent: " + err.Error()})
		}
		responseCodeAndMsg(rw, http.StatusAccepted, MsgObj{Msg: "Email sent success"})
	})
}

func EmailReport() (err error) {

	m := mail.NewV3Mail()

	from := mail.NewEmail("MyLab Software", "mylab.software5656@gmail.com")
	plainTextContent := mail.NewContent("text/plain", "Test mail")
	to := mail.NewEmail("Santosh Kavhar", "santosh.kavhar@joshsoftware.com")
	subject := "Sending sample Report - With Report Attachment"
	
	m.SetFrom(from)
	m.AddContent(plainTextContent)

	// create new Personalization
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.Subject = subject 

	m.AddPersonalizations(personalization)

	// attach PDF

	a_pdf := mail.NewAttachment()
	dat, err := ioutil.ReadFile("/home/josh/Desktop/Programs/GO/PDF/hello.pdf")
	if err != nil{
		logger.Errorln(err)
		return
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(dat))
	a_pdf.SetContent(encoded)
	a_pdf.SetType("application/pdf")
	a_pdf.SetFilename("Experiment_Report.pdf")
	a_pdf.SetDisposition("attachment")


	// attach a xlsx report
	a_xlsx := mail.NewAttachment()
	dat, err = ioutil.ReadFile("/home/josh/Desktop/CPAgent/cpagent/utils/output/output_e83ec5c7-5fa9-4e73-954e-4e7e18c95bb1_1626186896.xlsx")
	if err != nil{
		logger.Errorln(err)
		return
	}

	encoded = base64.StdEncoding.EncodeToString([]byte(dat))
	a_xlsx.SetContent(encoded)
	a_xlsx.SetType("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	a_xlsx.SetFilename("Experiment_e83ec5c7-5fa9-4e73-954e-4e7e18c95bb1.xlsx")
	a_xlsx.SetDisposition("attachment")


	m.AddAttachment(a_pdf)
	m.AddAttachment(a_xlsx)


	request := sendgrid.GetRequest(config.GetSendGridAPIKey(), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response , err := sendgrid.API(request)
	if err != nil{
		logger.Errorln(err)
		return
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return
}
