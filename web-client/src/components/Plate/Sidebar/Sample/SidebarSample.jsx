import React from "react";
import PropTypes from "prop-types";
import { Button, Select, CreatableSelect } from "core-components";
import Sidebar from "components/Sidebar";
import SampleTargetList from "./SampleTargetList";

const SidebarSample = (props) => {
  const {
    sampleState,
    updateCreateSampleWrapper,
    fetchSamples,
    sampleOptions,
    addNewLocalSample,
    isSampleListLoading,
    taskOptions,
    onTargetClickHandler,
    addButtonClickHandler,
    isSampleStateValid, // form state valid
    isDisabled,
    resetLocalState
  } = props;

  const { isSideBarOpen, sample, task, isEdit } = sampleState.toJS();

  const toggleSideBar = () => {
    // console log on sample drawer handle click
    console.info("Sample drawer handle clicked");
    // if user close sidebar without editing then reset local state
    if (isSideBarOpen === true && isEdit === true) {
      resetLocalState();
    }
    updateCreateSampleWrapper("isSideBarOpen", !isSideBarOpen);
  };

  const handleSampleCreate = (inputValue) => {
    const newOption = {
      label: inputValue,
      value: inputValue
    };
    // update local state
    updateCreateSampleWrapper("sample", newOption);
    // add new sample to sample's reducer to merge it with original list
    addNewLocalSample(newOption);
  };

  const handleSampleInputChange = (text) => {
    // fetch samples if text length is greater than zero ie.text is not empty
    if (text.length > 0) {
      fetchSamples(text);
    }
  };

  const handleSampleChange = (value) => {
    updateCreateSampleWrapper("sample", value);
  };

  const handleTaskChange = (value) => {
    updateCreateSampleWrapper("task", value);
  };

  return (
    // <Sidebar
    // 	isOpen={true}//{isSideBarOpen}
    // 	toggleSideBar={toggleSideBar}
    // 	className="sample"
    // 	handleIcon="plus-1"
    // 	handleIconSize={36}
    // 	isDisabled={isDisabled}
    // >
    <>
      <CreatableSelect
        isClearable
        isDisabled={isDisabled || isSampleListLoading}
        isLoading={isSampleListLoading}
        onChange={handleSampleChange}
        onCreateOption={handleSampleCreate}
        onInputChange={handleSampleInputChange}
        options={sampleOptions}
        value={sample}
        placeholder="Select Sample"
        className="mb-3"
      />
      <SampleTargetList
        list={sampleState.get("targets")}
        onTargetClickHandler={onTargetClickHandler}
        isDisabled={isDisabled}
      />
      <Select
        placeholder="Select Task"
        className="mb-4"
        options={taskOptions}
        value={task}
        onChange={handleTaskChange}
        isDisabled={isDisabled}
      />
      <Button
        className="mt-auto ml-2"
        disabled={isDisabled || !isSampleStateValid}
        onClick={addButtonClickHandler}
        color="primary"
      >
        Add
      </Button>
      {/* // </Sidebar> */}
    </>
  );
};

SidebarSample.propTypes = {
  sampleState: PropTypes.object.isRequired,
  updateCreateSampleWrapper: PropTypes.func.isRequired,
  fetchSamples: PropTypes.func.isRequired,
  sampleOptions: PropTypes.array.isRequired,
  addNewLocalSample: PropTypes.func.isRequired,
  isSampleListLoading: PropTypes.bool.isRequired,
  taskOptions: PropTypes.array.isRequired,
  onTargetClickHandler: PropTypes.func.isRequired,
  addButtonClickHandler: PropTypes.func.isRequired,
  isSampleStateValid: PropTypes.bool.isRequired,
  isDisabled: PropTypes.bool.isRequired,
  resetLocalState: PropTypes.func.isRequired
};

export default React.memo(SidebarSample);
