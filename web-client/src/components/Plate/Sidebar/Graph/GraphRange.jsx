import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { Text } from "shared-components";
import { FormGroup, Label, Input, Button } from "core-components";
import { useFormik } from "formik";
import {
  getFormikInitialState,
  getRequestBody,
  disbleApplyBtn,
  disbleResetBtn,
} from "./helper";
import { getMaxThreshold } from "components/Plate/helpers";
import { DEFAULT_MIN_VALUE } from "appConstants";

const GraphRange = (props) => {
  const {
    className,
    handleRangeChangeBtn,
    handleResetBtn,
    headerData,
    data,
    targets,
    isExpanded,
    options,
  } = props;

  const { progressStatus } = headerData;
  const nCycles = data?.labels?.length; // number of cycles

  const formik = useFormik({
    initialValues: getFormikInitialState(nCycles),
    enableReinitialize: true,
  });

  useEffect(() => {
    const {
      scales: { xAxes, yAxes },
    } = options;

    // pre-fill values to show in the input fields after update/reset
    formik.setFieldValue("xMin.value", xAxes[0].ticks.min);
    formik.setFieldValue("xMax.value", xAxes[0].ticks.max);
    formik.setFieldValue("yMin.value", yAxes[0].ticks.min);
    formik.setFieldValue("yMax.value", yAxes[0].ticks.max);
  }, [options]);

  const handleBlurChange = ({ name, value }) => {
    const { min, max } = formik.values[`${name}`];

    if (value > max || value < min) {
      formik.setFieldValue(`${name}.isInvalid`, true);
    }
  };

  const applyBtnHandler = () => {
    const requestBody = getRequestBody(formik.values);
    const maxThreshold = getMaxThreshold(targets.toJS());
    handleRangeChangeBtn(requestBody, nCycles, maxThreshold);
  };

  const resetBtnHandler = () => {
    const maxThreshold = getMaxThreshold(targets.toJS());
    handleResetBtn(nCycles, maxThreshold);

    const { xMax, xMin, yMax, yMin } = formik.values;
    const { yAxisMin, xAxisMin } = DEFAULT_MIN_VALUE;

    // if values are empty, pre-fill input fields
    if (
      xMax.value === null ||
      xMax.value === "" ||
      xMin.value === null ||
      xMin.value === "" ||
      yMax.value === null ||
      yMax.value === "" ||
      yMin.value === null ||
      yMin.value === ""
    ) {
      formik.setFieldValue("xMin.value", xAxisMin);
      formik.setFieldValue("xMax.value", nCycles);
      formik.setFieldValue("yMin.value", yAxisMin);
      formik.setFieldValue("yMax.value", maxThreshold);
    }
  };

  return (
    <div className={`graph-range d-flex ${className}`}>
      <Text Tag="h4" size={19} className="flex-10 title mb-0 pr-3">
        Range
      </Text>
      <div className="d-flex align-items-center flex-wrap flex-90">
        <FormGroup className="d-flex align-items-center flex-35 px-2">
          <Label className="flex-20 text-right mb-0 p-1">X Axis</Label>

          <Input
            name="xMin"
            type="number"
            className="px-2 py-1 ml-2"
            placeholder="Min value"
            value={formik.values.xMin.value}
            onChange={(e) => formik.setFieldValue("xMin.value", e.target.value)}
            onBlur={(event) => handleBlurChange(event.target)}
            onFocus={() => formik.setFieldValue(`xMin.isInvalid`, false)}
          />

          <Input
            name="xMax"
            type="number"
            className="px-2 py-1 ml-2"
            placeholder="Max value"
            value={formik.values.xMax.value}
            onChange={(e) => formik.setFieldValue("xMax.value", e.target.value)}
            onBlur={(event) => handleBlurChange(event.target)}
            onFocus={() => formik.setFieldValue(`xMax.isInvalid`, false)}
          />
        </FormGroup>
        <FormGroup className="d-flex align-items-center flex-35 px-2">
          <Label className="flex-20 text-right mb-0 p-1">Y Axis</Label>
          <Input
            name="yMin"
            type="number"
            step="0.1"
            className="px-2 py-1 ml-2"
            placeholder="Min value"
            value={formik.values.yMin.value}
            onChange={(e) => formik.setFieldValue("yMin.value", e.target.value)}
            onBlur={(event) => handleBlurChange(event.target)}
            onFocus={() => formik.setFieldValue(`yMin.isInvalid`, false)}
          />
          <Input
            name="yMax"
            type="number"
            step="0.1"
            className="px-2 py-1 ml-2"
            placeholder="Max value"
            value={formik.values.yMax.value}
            onChange={(e) => formik.setFieldValue("yMax.value", e.target.value)}
            onBlur={(event) => handleBlurChange(event.target)}
            onFocus={() => formik.setFieldValue(`yMax.isInvalid`, false)}
          />
        </FormGroup>
        <Button
          color="primary"
          size="sm"
          className="mb-3 ml-3"
          onClick={applyBtnHandler}
          disabled={disbleApplyBtn(formik.values, progressStatus, isExpanded)}
        >
          Apply
        </Button>
        <Button
          color="secondary"
          size="sm"
          outline={true}
          className="mb-3 ml-3 border-2 border-gray "
          onClick={resetBtnHandler}
          disabled={disbleResetBtn(progressStatus, isExpanded)}
        >
          Reset
        </Button>
      </div>
    </div>
  );
};

GraphRange.propTypes = {
  className: PropTypes.string,
};

GraphRange.defaultProps = {
  className: "",
};

export default GraphRange;
