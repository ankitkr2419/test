import React from "react";
import { Modal, ModalBody } from "core-components";
import { Text, ButtonIcon } from "shared-components";
import { ExperimentGraphContainer } from "containers/ExperimentGraphContainer";
import Header from "components/Plate/Header";
import { Button } from "core-components";

const PreviewReportModal = (props) => {
  const {
    isOpen,
    toggleModal,
    onDownloadConfirmed,
    experimentStatus,
    isSidebarOpen,
    setIsSidebarOpen,
    resetSelectedWells,
    isMultiSelectionOptionOn,
    progressStatus,
    progress,
    remainingTime,
    totalTime,
    experimentTemplate,
    experimentDetails,
    experimentId,
    temperatureData,
  } = props;

  return (
    <Modal isOpen={isOpen} toggle={toggleModal} centered size="lg">
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

        {/** View Report */}
        <div
          className="d-flex flex-column align-items-center export-contents"
          style={{ transform: "scale(0.8)" }}
        >
          <div id="page-1">
            {/** amplification graph */}
            <Text className="font-weight-bold text-center mb-4">
              Amplification
            </Text>
            <ExperimentGraphContainer
              showTempGraph={false}
              experimentStatus={experimentStatus}
              isSidebarOpen={isSidebarOpen}
              setIsSidebarOpen={setIsSidebarOpen}
              resetSelectedWells={resetSelectedWells}
              isMultiSelectionOptionOn={isMultiSelectionOptionOn}
            />
          </div>
          <div id="page-2">
            {/** temp graph */}
            <Text className="font-weight-bold text-center mb-4">
              Temperature
            </Text>
            <ExperimentGraphContainer
              showTempGraph={true}
              experimentStatus={experimentStatus}
              isSidebarOpen={isSidebarOpen}
              setIsSidebarOpen={setIsSidebarOpen}
              resetSelectedWells={resetSelectedWells}
              isMultiSelectionOptionOn={isMultiSelectionOptionOn}
            />
          </div>
          <div id="page-3">
            <Header
              progressStatus={progressStatus}
              progress={progress}
              remainingTime={remainingTime}
              totalTime={totalTime}
              experimentTemplate={experimentTemplate}
              experimentStatus={experimentStatus}
              experimentDetails={experimentDetails}
              experimentId={experimentId}
              temperatureData={temperatureData}
            />
          </div>
        </div>
      </ModalBody>
    </Modal>
  );
};

export default React.memo(PreviewReportModal);
