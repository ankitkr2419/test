import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { useFormik } from "formik";

import { Select, Switch } from "core-components";
import { Text } from "shared-components";
import { formikInitialState, getInitialState } from "./helper";
import { FormGroup, Label, Input, Button } from "core-components";

const Filters = (props) => {
  let {
    targetOptions,
    selectedTarget,
    onTargetChanged,
    analyseDataGraphFilters,
    onFiltersChanged,
    onResetThresholdFilter,
    onResetBaselineFilter,
  } = props;

  const formik = useFormik({
    initialValues: analyseDataGraphFilters
      ? getInitialState(analyseDataGraphFilters)
      : formikInitialState,
    enableReinitialize: true,
  });

  let isAutoThreshold = formik.values.isAutoThreshold.value;
  let thresholdLabel = isAutoThreshold === true ? "Auto" : "Manual";
  let isAutoBaseline = formik.values.isAutoBaseline.value;
  let baselineLabel = isAutoBaseline === true ? "Auto" : "Manual";

  const onFormikValueChanged = (key, value) => {
    formik.setFieldValue(key, value);
  };

  const handleThresholdApplyButton = () => {
    onFiltersChanged({
      isAutoThreshold: isAutoThreshold,
      threshold: formik.values.threshold.value,
    });
  };
  const handleThresholdResetButton = () => {
    onResetThresholdFilter();
  };
  const handleBaselineApplyButton = () => {
    onFiltersChanged({
      isAutoBaseline: isAutoBaseline,
      startCycle: formik.values.startCycle.value,
      endCycle: formik.values.endCycle.value,
    });
  };
  const handleBaselineResetButton = () => {
    onResetBaselineFilter();
  };

  return (
    <>
      {/** Target selector */}
      <div className="graph-filters d-flex">
        <Text Tag="h4" size={19} className="flex-10 title mb-0 pr-3">
          Target
        </Text>
        <div style={{ width: "350px" }}>
          <Select
            placeholder="Select Target"
            className="mb-4"
            options={targetOptions}
            value={selectedTarget}
            onChange={onTargetChanged}
          />
        </div>
      </div>

      {/**threshold */}
      <div className="d-flex">
        <Text Tag="h4" size={19} className="flex-10 title mb-0 pr-3">
          Threshold
        </Text>
        <div className="d-flex align-items-center flex-wrap flex-90">
          <FormGroup className="d-flex align-items-center px-2">
            <Switch
              id="thresholdToggler"
              name="thresholdToggler"
              label={thresholdLabel}
              checked={isAutoThreshold}
              onChange={() =>
                onFormikValueChanged("isAutoThreshold.value", !isAutoThreshold)
              }
            />
            {isAutoThreshold === false && (
              <>
                <Label className="flex-40 text-right mb-0 p-1">Threshold</Label>
                <Input
                  name="thresholdValue"
                  type="number"
                  className="flex-25 px-2 py-1 ml-2"
                  placeholder="Threshold"
                  step="0.1"
                  value={formik.values.threshold.value}
                  onChange={(e) =>
                    onFormikValueChanged(
                      "threshold.value",
                      parseFloat(e.target.value)
                    )
                  }
                />
              </>
            )}
          </FormGroup>
          <div className="ml-auto">
            <Button
              color="primary"
              size="sm"
              className="mb-3 ml-3"
              onClick={handleThresholdApplyButton}
            >
              Apply
            </Button>
            <Button
              color="secondary"
              size="sm"
              outline={true}
              className="mb-3 ml-3 border-2 border-gray "
              onClick={handleThresholdResetButton}
            >
              Reset
            </Button>
          </div>
        </div>
      </div>

      {/**baseline */}
      <div className="d-flex">
        <Text Tag="h4" size={19} className="flex-10 title mb-0 pr-3">
          Baseline
        </Text>
        <div className="d-flex align-items-center flex-wrap  flex-90">
          <FormGroup className="d-flex align-items-center flex-35 px-2">
            <Switch
              id="baselineToggler"
              name="baselineToggler"
              label={baselineLabel}
              checked={isAutoBaseline}
              onChange={() =>
                onFormikValueChanged("isAutoBaseline.value", !isAutoBaseline)
              }
            />
            {isAutoBaseline === false && (
              <>
                <Label className="flex-40 text-right mb-0 p-1">
                  Start Cycle
                </Label>
                <Input
                  name="startCycle"
                  type="number"
                  className="flex-30 px-2 py-1 ml-2"
                  placeholder="Start Cycle"
                  value={formik.values.startCycle.value}
                  onChange={(e) =>
                    onFormikValueChanged(
                      "startCycle.value",
                      parseInt(e.target.value)
                    )
                  }
                />
                <Label className="flex-40 text-right mb-0 p-1">End Cycle</Label>
                <Input
                  name="endCycle"
                  type="number"
                  className="flex-30 px-2 py-1 ml-2"
                  placeholder="End Cycle"
                  value={formik.values.endCycle.value}
                  onChange={(e) =>
                    onFormikValueChanged(
                      "endCycle.value",
                      parseInt(e.target.value)
                    )
                  }
                />
              </>
            )}
          </FormGroup>
          <div className="ml-auto">
            <Button
              color="primary"
              size="sm"
              className="mb-3 ml-3"
              onClick={handleBaselineApplyButton}
            >
              Apply
            </Button>
            <Button
              color="secondary"
              size="sm"
              outline={true}
              className="mb-3 ml-3 border-2 border-gray "
              onClick={handleBaselineResetButton}
            >
              Reset
            </Button>
          </div>
        </div>
      </div>
    </>
  );
};

Filters.propTypes = {
  targetOptions: PropTypes.array.isRequired,
  selectedTarget: PropTypes.object.isRequired,
  onTargetChanged: PropTypes.func.isRequired,
  analyseDataGraphFilters: PropTypes.object.isRequired,
  onFiltersChanged: PropTypes.func.isRequired,
  onResetThresholdFilter: PropTypes.func.isRequired,
  onResetBaselineFilter: PropTypes.func.isRequired,
};

export default React.memo(Filters);
