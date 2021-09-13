import React, { useEffect, useReducer, useState } from "react";
import PropTypes from "prop-types";
import { useSelector, useDispatch } from "react-redux";

import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import generatePdfIcon from "assets/images/generatePdfIcon.svg";

import { ExperimentGraphContainer } from "containers/ExperimentGraphContainer";
import { getRunExperimentReducer } from "selectors/runExperimentSelector";
import SampleSideBarContainer from "containers/SampleSideBarContainer";
import { EXPERIMENT_STATUS } from "appConstants";
import Header from "./Header";

import GridComponent from "./Grid";
import WellGridHeader from "./Grid/WellGridHeader";
import SelectAllGridHeader from "./Grid/SelectAllGridHeader";
import { Button } from "core-components";
import { ButtonIcon, ImageIcon, Text } from "shared-components";
import PreviewReportModal from "components/modals/PreviewReportModal";
import { graphs } from "./plateConstant";
import { getExperimentGraphTargets } from "selectors/experimentTargetSelector";
import { updateFilter } from "action-creators/analyseDataGraphActionCreators";
import { generateTargetOptions } from "components/AnalyseDataGraph/helper";
import { rangeActions, rangeInitialState, rangeReducer } from "./helpers";

import GridWrapper from "./Grid/GridWrapper";
import "./Plate.scss";

const initialOptions = {
  legend: {
    display: false,
  },
  animation: false,
  scales: {
    xAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "Cycles",
          fontSize: 15,
          fontStyle: "bold",
          padding: 5,
        },
        offset: true,
        ticks: {
          fontSize: 15,
          fontStyle: "bold",
          min: 0,
          max: 4,
        },
      },
    ],
    yAxes: [
      {
        scaleLabel: {
          display: true,
          labelString: "F-value",
          fontSize: 15,
          fontStyle: "bold",
          padding: 10,
        },
        ticks: {
          fontSize: 15,
          fontStyle: "bold",
          min: 0,
          max: 1,
        },
      },
    ],
  },
};

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
    isExpanded,
  } = props;

  const dispatch = useDispatch();
  // getExperimentStatus will return us current experiment status
  const runExperimentDetails = useSelector(getRunExperimentReducer);
  const createExperimentReducer = useSelector(
    (state) => state.createExperimentReducer
  );

  // get targets from experiment target reducer(graph : target filters)
  const experimentGraphTargetsList = useSelector(getExperimentGraphTargets);
  const targetsData = experimentGraphTargetsList.toJS();

  const thresholdArr = targetsData?.map((targetObj) =>
    parseInt(targetObj.threshold)
  );
  const maxThreshold = thresholdArr.length > 0 ? Math.max(...thresholdArr) : 10;

  const experimentStatus = runExperimentDetails.get("experimentStatus");

  let experimentDetails =
    isExpanded === true
      ? createExperimentReducer.toJS()
      : runExperimentDetails.get("data").toJS();

  // local state to maintain well data which is selected for updation
  const [updateWell, setUpdateWell] = useState(null);

  // local state to maintain active graph
  const [activeGraph, setActiveGraph] = useState(graphs.Amplification);

  // local state to manage toggling of graphSidebar
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  // local state to manage previewReport modal
  const [previewReportModal, setPreviewReportModal] = useState(false);

  // local state to manage min max values of X and Y
  const [rangeState, updateRangeState] = useReducer(
    rangeReducer,
    rangeInitialState
  );
  const { xMaxValue, xMinValue, yMaxValue, yMinValue } = rangeState;

  const [options, setOptions] = useState(initialOptions);
  const [isDataFromAPI, setDataFromAPI] = useState(false);

  //set first target as selected to use it in analyse data graph
  useEffect(() => {
    const targets = generateTargetOptions(targetsData);
    if (targets.length > 0) {
      const firstTarget = targets[0];
      dispatch(updateFilter({ selectedTarget: firstTarget }));
    }
  }, [dispatch, targetsData]);

  useEffect(() => {
    let newOptions = {
      ...initialOptions,
      scales: {
        ...initialOptions.scales,
        xAxes: [
          {
            ...initialOptions.scales.xAxes[0],
            ticks: {
              ...initialOptions.scales.xAxes[0].ticks,
              min: xMinValue,
              max: xMaxValue,
            },
          },
        ],
        yAxes: [
          {
            ...initialOptions.scales.yAxes[0],
            ticks: {
              ...initialOptions.scales.yAxes[0].ticks,
              min: yMinValue,
              max: yMaxValue,
            },
          },
        ],
      },
    };

    setOptions(newOptions);
  }, [xMaxValue, xMinValue, yMaxValue, yMinValue]);

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
     * 				then we can make well selected and isExpanded === false
     */
    if (
      isMultiSelectionOptionOn === false &&
      isWellFilled === false &&
      isExpanded === false
    ) {
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
    // if the fields are pre-filled, empty it
    setUpdateWell(null);
  };

  const onWellUpdateClickHandler = (selectedWell) => {
    // update local state with selected well which is selected for updation
    let wellToUpdate = null;
    if (selectedWell) {
      wellToUpdate = selectedWell.toJS();
    }
    setUpdateWell(wellToUpdate);
  };

  // hleper function to open sidebar and show graph of selected well
  const showGraphOfWell = (index, show) => {
    // set selected well index
    setSelectedWell(index, show);
  };

  const [activeTab, setActiveTab] = useState("wells");

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

  // helper function to toggle the graphs
  const onChangeActiveGraph = (graphType) => {
    if (activeGraph !== graphType) {
      setActiveGraph(graphType);
    }
  };

  const makeAmplificationGraphActive = () => {
    onChangeActiveGraph(graphs.Amplification);
  };
  const makeTemperatureGraphActive = () => {
    onChangeActiveGraph(graphs.Temperature);
  };
  const makeAnalyseDataGraphActive = () => {
    onChangeActiveGraph(graphs.AnalyseData);
  };

  const togglePreviewReportModal = () => {
    setPreviewReportModal(!previewReportModal);
  };

  const downloadClickHandler = (e) => {
    togglePreviewReportModal();
  };

  const handleRangeChangeBtn = ({ xMax, xMin, yMax, yMin }) => {
    setDataFromAPI(true);

    updateRangeState({ type: rangeActions.UPDATE_X_MAX, value: xMax });
    updateRangeState({ type: rangeActions.UPDATE_X_MIN, value: xMin });
    updateRangeState({ type: rangeActions.UPDATE_Y_MAX, value: yMax });
    updateRangeState({ type: rangeActions.UPDATE_Y_MIN, value: yMin });
  };

  const handleResetBtn = (cycleCount) => {
    setDataFromAPI(true);
    updateRangeState({ type: rangeActions.RESET_VALUES });
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
          mailBtnHandler={mailBtnHandler}
          options={options}
          isDataFromAPI={isDataFromAPI}
          experimentGraphTargetsList={experimentGraphTargetsList}
        />
      )}
      <Header
        data={headerData}
        experimentTemplate={experimentTemplate}
        experimentStatus={experimentStatus}
        experimentDetails={experimentDetails}
        experimentId={experimentId}
        temperatureData={temperatureData}
        isExpanded={isExpanded}
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
                  experimentStatus === EXPERIMENT_STATUS.stopped ||
                  isExpanded === true
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
                  isExpanded={isExpanded}
                  onWellUpdateClickHandler={onWellUpdateClickHandler}
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
                    isExpanded={isExpanded}
                  />
                  <SelectAllGridHeader
                    isAllWellsSelected={isAllWellsSelected}
                    toggleAllWellSelectedOption={toggleAllWellSelectedOption}
                    experimentStatus={experimentStatus}
                    isExpanded={isExpanded}
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
                  isExpanded={isExpanded}
                />
              </div>
            </div>
          </TabPane>
          <TabPane className="tab-pane-graph flex-100 scroll-y" tabId="graph">
            <div className="d-flex flex-100">
              <div className="graph-wrapper flex-100 scroll-y">
                <div className="d-flex align-items-center mb-3">
                  <Button
                    outline={activeGraph !== graphs.Amplification}
                    color={
                      activeGraph === graphs.Amplification
                        ? "primary"
                        : "secondary"
                    }
                    className="mr-3 Amplification"
                    onClick={makeAmplificationGraphActive}
                  >
                    Amplification
                  </Button>
                  <Button
                    outline={activeGraph !== graphs.Temperature}
                    color={
                      activeGraph === graphs.Temperature
                        ? "primary"
                        : "secondary"
                    }
                    className="mr-3 Temperature"
                    onClick={makeTemperatureGraphActive}
                  >
                    Temperature
                  </Button>
                  {(experimentStatus === EXPERIMENT_STATUS.success ||
                    experimentStatus === EXPERIMENT_STATUS.stopped ||
                    isExpanded === true) && (
                    <Button
                      outline={activeGraph !== graphs.AnalyseData}
                      color={
                        activeGraph === graphs.AnalyseData
                          ? "primary"
                          : "secondary"
                      }
                      className="mr-3 AnalyseData"
                      onClick={makeAnalyseDataGraphActive}
                    >
                      Analyse Data
                    </Button>
                  )}

                  {/* <ButtonIcon
                    name="download-1"
                    size={28}
                    className="bg-white border-secondary ml-auto downloadButton"
                    onClick={downloadClickHandler}
                  /> */}

                  <div className="ml-auto" onClick={downloadClickHandler}>
                    <ImageIcon
                      src={generatePdfIcon}
                      alt="icon not available"
                      style={{ cursor: "pointer", maxHeight: 40 }}
                    />
                  </div>
                </div>

                <ExperimentGraphContainer
                  isInsidePreviewModal={false}
                  headerData={headerData}
                  activeGraph={activeGraph}
                  experimentStatus={experimentStatus}
                  isSidebarOpen={isSidebarOpen}
                  setIsSidebarOpen={setIsSidebarOpen}
                  resetSelectedWells={resetSelectedWells}
                  isMultiSelectionOptionOn={isMultiSelectionOptionOn}
                  isExpanded={isExpanded}
                  handleRangeChangeBtn={handleRangeChangeBtn}
                  handleResetBtn={handleResetBtn}
                  options={options}
                  isDataFromAPI={isDataFromAPI}
                  experimentGraphTargetsList={experimentGraphTargetsList}
                />
              </div>
            </div>
          </TabPane>
        </TabContent>
      </GridWrapper>
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
