import React, { useEffect, useState } from "react";
import { Modal, ModalBody, Button } from "core-components";
import { Text, ButtonIcon } from "shared-components";
import { ExperimentGraphContainer } from "containers/ExperimentGraphContainer";
import Header from "components/Plate/Header";
import html2canvas from "html2canvas";
import { jsPDF } from "jspdf";
import { useSelector, useDispatch } from "react-redux";
import { saveReportInitiated } from "action-creators/reportActionCreators";

const PreviewReportModal = (props) => {
  const dispatch = useDispatch();
  const {
    isOpen,
    toggleModal,
    experimentStatus,
    isSidebarOpen,
    setIsSidebarOpen,
    resetSelectedWells,
    isMultiSelectionOptionOn,
    data,
    experimentTemplate,
    experimentDetails,
    experimentId,
    temperatureData,
  } = props;

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  //get activity logs from reducer
  const reportReducer = useSelector((state) => state.reportReducer);
  const reportReducerData = reportReducer.toJS();
  const { isLoading } = reportReducerData;

  //local state for maintaining loading for generate report
  const [loading, setLoading] = useState(false);
  //also synce local loading with save report api isLoading when report is generated
  const [isReportGenerated, setIsReportGenerated] = useState(false);


  //sync local loading with api isLoading when report generated
  useEffect(() => {
    if (isReportGenerated === true) {
      setLoading(isLoading);
    }
  }, [isLoading, isReportGenerated]);

  const onDownloadConfirmed = () => {
    setLoading(true);
    generateReport();
  };

  const generateReport = () => {
    //access html element we want to export contents from
    let page1 = document.getElementById("page-1");
    let page2 = document.getElementById("page-2");
    // let page3 = document.getElementById("page-3");

    //convert html element into canvas
    html2canvas(page1).then((canvas1) => {
      html2canvas(page2).then((canvas2) => {
        // html2canvas(page3).then((canvas3) => {
        const img1 = canvas1.toDataURL("image/png");
        const img2 = canvas2.toDataURL("image/png");
        // const img3 = canvas3.toDataURL("image/png");
        const doc = new jsPDF("l", "pt", [1024, 700]); //create custom pdf instance with its properties [width, height] in pt units //NOTE: (1 pt = 1.3281472327365 px)
        //add contents into pdf
        doc.addImage(
          img1,
          "png",
          30, //page1.offsetLeft,
          100, //page1.offsetTop,
          page1.clientWidth,
          page1.clientHeight
        );
        doc.addPage();
        doc.addImage(
          img2,
          "png",
          30, //page2.offsetLeft,//left
          100, //page2.offsetTop//top
          page2.clientWidth,
          page2.clientHeight
        );
        // doc.save("chart.pdf"); //save file locally with default filename //kept for testing
        setIsReportGenerated(true);
        let pdfInBlobFormat = doc.output("blob"); //generate blob file

        //save to server
        sendReportToServer(pdfInBlobFormat);

        // }); //page 3 finish
      }); //page 2 finish
    }); //page 1 finish
  };

  const sendReportToServer = (pdfInBlobFormat) => {
    //create file from blob
    var file = new File([pdfInBlobFormat], "report.pdf", {
      lastModified: new Date().getTime(),
    });
    //send report file to api
    var data = new FormData();
    data.append("report", file);
    dispatch(saveReportInitiated(token, experimentId, data));
  };

  const showPageOne = () => {
    return (
      <div id="page-1">
        {/** amplification graph */}
        <Text className="font-weight-bold text-center mb-4">Amplification</Text>
        <ExperimentGraphContainer
          showTempGraph={false}
          experimentStatus={experimentStatus}
          isSidebarOpen={isSidebarOpen}
          setIsSidebarOpen={setIsSidebarOpen}
          resetSelectedWells={resetSelectedWells}
          isMultiSelectionOptionOn={isMultiSelectionOptionOn}
        />
      </div>
    );
  };

  const showPageTwo = () => {
    return (
      <div id="page-2">
        {/** temp graph */}
        <Text className="font-weight-bold text-center mb-4">Temperature</Text>
        <ExperimentGraphContainer
          showTempGraph={true}
          experimentStatus={experimentStatus}
          isSidebarOpen={isSidebarOpen}
          setIsSidebarOpen={setIsSidebarOpen}
          resetSelectedWells={resetSelectedWells}
          isMultiSelectionOptionOn={isMultiSelectionOptionOn}
        />
      </div>
    );
  };

  {
    /*TODO remove if not needed*/
  }
  // const showPageThree = () => {
  //   return (
  //     <div id="page-3">
  //       <Header
  //         data={data}
  //         experimentTemplate={experimentTemplate}
  //         experimentStatus={experimentStatus}
  //         experimentDetails={experimentDetails}
  //         experimentId={experimentId}
  //         temperatureData={temperatureData}
  //       />
  //     </div>
  //   );
  // };

  return (
    <Modal isOpen={isOpen} centered size="lg">
      <ModalBody>
        <div className="d-flex mt-5">
          <Text
            Tag="h4"
            size={24}
            className="text-center text-primary my-2 mr-auto"
          >
            Report
          </Text>
          <Button
            onClick={onDownloadConfirmed}
            color="primary"
            className={"ml-auto"}
            size="md"
          >
            Download
          </Button>
        </div>

        {loading === false && (
          <ButtonIcon
            position="absolute"
            placement="right"
            top={4}
            right={8}
            size={36}
            name="cross"
            onClick={toggleModal}
            className="border-0"
          />
        )}

        {/** View Report */}
        <div
          className="d-flex flex-column align-items-center export-contents"
          //adjust scale to fit graphs on modal width
          //& disable user interactions
          style={{ transform: "scale(0.8)", pointerEvents: "none" }}
        >
          {showPageOne()}
          {showPageTwo()}
          {/* {showPageThree()} */}
        </div>
      </ModalBody>
    </Modal>
  );
};

export default React.memo(PreviewReportModal);
