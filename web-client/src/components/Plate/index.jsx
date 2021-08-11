import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import { useSelector } from "react-redux";
import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";

import { ExperimentGraphContainer } from "containers/ExperimentGraphContainer";
import { getRunExperimentReducer } from "selectors/runExperimentSelector";
import SampleSideBarContainer from "containers/SampleSideBarContainer";
import { EXPERIMENT_STATUS } from "appConstants";
import Header from "./Header";

import GridWrapper from "./Grid/GridWrapper";
import GridComponent from "./Grid";
import WellGridHeader from "./Grid/WellGridHeader";

import "./Plate.scss";
import SelectAllGridHeader from "./Grid/SelectAllGridHeader";
import { Button } from "core-components";
import { ButtonIcon, Text } from "shared-components";
import PreviewReportModal from "components/modals/PreviewReportModal";

const Plate = (props) => {
  const {
    wells,
    setSelectedWell,
    setMultiSelectedWell,
    experimentTargetsList,
    positions,
    experimentId,
    isMultiSelectionOptionOn,
    isAllWellsSelected,
    toggleMultiSelectOption,
    toggleAllWellSelectedOption,
    activeWells,
    experimentTemplate,
    resetSelectedWells,
    headerData,
    temperatureData,
    mailBtnHandler,
    token,
  } = props;

  // getExperimentStatus will return us current experiment status
  const runExperimentDetails = useSelector(getRunExperimentReducer);
  const experimentStatus = runExperimentDetails.get("experimentStatus");
  const experimentDetails = runExperimentDetails.get("data");

  // local state to maintain well data which is selected for updation
  const [updateWell, setUpdateWell] = useState(null);

  // local state to toggle between emission graph and temperature graph
  const [showTempGraph, setShowTempGraph] = useState(false);

  // local state to manage toggling of graphSidebar
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  // local state to manage previewReport modal
  const [previewReportModal, setPreviewReportModal] = useState(false);

  useEffect(() => {
    if (
      experimentStatus === EXPERIMENT_STATUS.success &&
      activeTab !== "graph"
    ) {
      setActiveTab("graph");
    }
  }, [experimentStatus]);

  /**
   *
   * @param {*} well => selected well details
   * @param {*} index => selected well index
   *
   */
  const onWellClickHandler = (well, index) => {
    const { isSelected, isWellFilled, isMultiSelected } = well.toJS();
    /**
     * if well is not filled and if multi selection option is not checked
     * 				then we can make well selected
     */
    if (isMultiSelectionOptionOn === false && isWellFilled === false) {
      setSelectedWell(index, !isSelected);
    }

    /**
     * if multi-select checkbox is checked, will allow to select filled wells
     */
    if (isMultiSelectionOptionOn === true) {
      // if (isWellFilled === true) {
      setMultiSelectedWell(index, !isMultiSelected);
      // }
    }
  };

  const onWellUpdateClickHandler = (selectedWell) => {
    // update local state with selected well which is selected for updation
    setUpdateWell(selectedWell.toJS());
  };

  // hleper function to open sidebar and show graph of selected well
  const showGraphOfWell = (index, show) => {
    // set selected well index
    setSelectedWell(index, show);
    // setIsSidebarOpen(true);
  };

  const [activeTab, setActiveTab] = useState("wells");

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

  // helper function to toggle the graphs
  const toggleTempGraphSwitch = (graphType) => {
    setShowTempGraph(graphType === "temperature");
  };

  const togglePreviewReportModal = () => {
    setPreviewReportModal(!previewReportModal);
  };

  const downloadClickHandler = (e) => {
    togglePreviewReportModal();
  };

  return (
    <div className="plate-content d-flex flex-column h-100 position-relative scroll-y">
      {previewReportModal && (
        <PreviewReportModal
          isOpen={previewReportModal}
          toggleModal={togglePreviewReportModal}
          experimentStatus={experimentStatus}
          isSidebarOpen={isSidebarOpen}
          setIsSidebarOpen={setIsSidebarOpen}
          resetSelectedWells={resetSelectedWells}
          isMultiSelectionOptionOn={isMultiSelectionOptionOn}
          data={headerData}
          experimentTemplate={experimentTemplate}
          experimentDetails={experimentDetails}
          experimentId={experimentId}
          temperatureData={temperatureData}
        />
      )}
      <Header
        data={headerData}
        experimentTemplate={experimentTemplate}
        experimentStatus={experimentStatus}
        experimentDetails={experimentDetails}
        experimentId={experimentId}
        temperatureData={temperatureData}
      />
      <GridWrapper className="plate-body flex-100 scroll-y">
        <Nav className="plate-nav-tabs border-0" tabs>
          <NavItem>
            <NavLink
              className={classnames({ active: activeTab === "wells" })}
              onClick={() => {
                toggle("wells");
              }}
            >
              Wells
            </NavLink>
          </NavItem>
          <NavItem>
            <NavLink
              className={classnames({ active: activeTab === "graph" })}
              onClick={() => {
                toggle("graph");
              }}
              disabled={
                !(
                  experimentStatus === EXPERIMENT_STATUS.success ||
                  experimentStatus === EXPERIMENT_STATUS.running ||
                  experimentStatus === EXPERIMENT_STATUS.stopped
                )
              }
            >
              Graph
            </NavLink>
          </NavItem>
        </Nav>
        <TabContent
          className="plate-tab-content d-flex scroll-y"
          activeTab={activeTab}
        >
          <TabPane className="tab-pane-wells flex-100 scroll-y" tabId="wells">
            <div className="d-flex flex-100">
              <div className="sample-wrapper d-flex flex-column scroll-y">
                <Text className="wrapper-title font-weight-bold text-center mb-4">
                  Add Samples
                </Text>
                <SampleSideBarContainer
                  experimentId={experimentId}
                  positions={positions}
                  experimentStatus={experimentStatus}
                  experimentTargetsList={experimentTargetsList}
                  updateWell={updateWell}
                />
              </div>
              <div className="wells-wrapper flex-100 scroll-y">
                <div className="d-flex align-items-center mb-4">
                  <WellGridHeader
                    className="mr-4"
                    wells={wells}
                    isGroupSelectionOn={isMultiSelectionOptionOn}
                    toggleMultiSelectOption={toggleMultiSelectOption}
                    experimentStatus={experimentStatus}
                  />
                  <SelectAllGridHeader
                    isAllWellsSelected={isAllWellsSelected}
                    toggleAllWellSelectedOption={toggleAllWellSelectedOption}
                    experimentStatus={experimentStatus}
                  />
                </div>
                <GridComponent
                  activeWells={activeWells}
                  wells={wells}
                  isGroupSelectionOn={isMultiSelectionOptionOn}
                  isAllWellsSelected={isAllWellsSelected}
                  onWellClickHandler={onWellClickHandler}
                  onWellUpdateClickHandler={onWellUpdateClickHandler}
                  showGraphOfWell={showGraphOfWell}
                  experimentStatus={experimentStatus}
                />
              </div>
            </div>
          </TabPane>
          <TabPane className="tab-pane-graph flex-100 scroll-y" tabId="graph">
            <div className="d-flex flex-100">
              <div className="graph-wrapper flex-100 scroll-y">
                <div className="d-flex align-items-center mb-3">
                  <Button
                    outline={showTempGraph}
                    color={!showTempGraph ? "primary" : "secondary"}
                    className="mr-3 Amplification"
                    onClick={() => toggleTempGraphSwitch("amplification")}
                  >
                    Amplification
                  </Button>
                  <Button
                    outline={!showTempGraph}
                    color={showTempGraph ? "primary" : "secondary"}
                    className="Temperature"
                    onClick={() => toggleTempGraphSwitch("temperature")}
                  >
                    Temperature
                  </Button>
                  <ButtonIcon
                    name="published"
                    size={28}
                    className="bg-white border-secondary ml-auto"
                    onClick={mailBtnHandler}
                  />
                  <ButtonIcon
                    name="download-1"
                    size={28}
                    className="bg-white border-secondary ml-3 downloadButton"
                    onClick={downloadClickHandler}
                  />
                </div>
                <ExperimentGraphContainer
                  token={token}
                  experimentId={experimentId}
                  headerData={headerData}
                  showTempGraph={showTempGraph}
                  experimentStatus={experimentStatus}
                  isSidebarOpen={isSidebarOpen}
                  setIsSidebarOpen={setIsSidebarOpen}
                  resetSelectedWells={resetSelectedWells}
                  isMultiSelectionOptionOn={isMultiSelectionOptionOn}
                />
              </div>
            </div>
          </TabPane>
        </TabContent>
      </GridWrapper>
      {/* <SampleSideBarContainer
        experimentId={experimentId}
        positions={positions}
        experimentTargetsList={experimentTargetsList}
        updateWell={updateWell}
      />
      <ExperimentGraphContainer
        experimentStatus={experimentStatus}
        isSidebarOpen={isSidebarOpen}
        setIsSidebarOpen={setIsSidebarOpen}
        resetSelectedWells={resetSelectedWells}
        isMultiSelectionOptionOn={isMultiSelectionOptionOn}
      /> */}
    </div>
  );
};

Plate.propTypes = {
  wells: PropTypes.object.isRequired,
  setSelectedWell: PropTypes.func.isRequired,
  setMultiSelectedWell: PropTypes.func.isRequired,
  // experimentTargetsList contains targets for selected experiment
  experimentTargetsList: PropTypes.object.isRequired,
  // array of selected wells with index
  positions: PropTypes.object.isRequired,
  experimentId: PropTypes.string.isRequired,
  isMultiSelectionOptionOn: PropTypes.bool.isRequired,
  toggleMultiSelectOption: PropTypes.func.isRequired,
  activeWells: PropTypes.object.isRequired,
  experimentTemplate: PropTypes.object.isRequired,
};

export default Plate;
