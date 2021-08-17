import React, { useCallback, useEffect } from "react";
import PropTypes from "prop-types";
import { Text } from "shared-components";
import { FormGroup, Label, Input, Button } from "core-components";
import { useFormik } from "formik";
import { formikInitialState, getRequestBody, disbleApplyBtn } from "./helper";
import { EXPERIMENT_STATUS } from "appConstants";

const GraphRange = (props) => {
  const {
    className,
    handleRangeChangeBtn,
    handleResetBtn,
    headerData,
    isExpanded,
  } = props;

  const { totalCycles, progressStatus } = headerData;

  const formik = useFormik({
    initialValues: formikInitialState,
    enableReinitialize: true,
  });

  //TODO: to be confirmed
  // useEffect(() => {
  //   formik.setFieldValue("xMax.max", parseInt(totalCycles));
  // }, [totalCycles]);

  const handleBlurChange = ({ name, value }) => {
    const { min, max } = formik.values[`${name}`];

    if (value > max || value < min) {
      formik.setFieldValue(`${name}.isInvalid`, true);
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
            onChange={(e) =>
              formik.setFieldValue("xMin.value", parseInt(e.target.value))
            }
            onBlur={(event) => handleBlurChange(event.target)}
            onFocus={() => formik.setFieldValue(`xMin.isInvalid`, false)}
          />
          {/* <div className="flex-70">
            <Text Tag="p" size={14} className="text-danger">
              This should be between {32} -{21}.
            </Text>
          </div> */}

          <Input
            name="xMax"
            type="number"
            className="px-2 py-1 ml-2"
            placeholder="Max value"
            value={formik.values.xMax.value}
            onChange={(e) =>
              formik.setFieldValue("xMax.value", parseInt(e.target.value))
            }
            onBlur={(event) => handleBlurChange(event.target)}
            onFocus={() => formik.setFieldValue(`xMax.isInvalid`, false)}
          />
        </FormGroup>
        <FormGroup className="d-flex align-items-center flex-35 px-2">
          <Label className="flex-20 text-right mb-0 p-1">Y Axis</Label>
          <Input
            name="yMin"
            type="number"
            className="px-2 py-1 ml-2"
            placeholder="Min value"
            value={formik.values.yMin.value}
            onChange={(e) =>
              formik.setFieldValue("yMin.value", parseInt(e.target.value))
            }
            onBlur={(event) => handleBlurChange(event.target)}
            onFocus={() => formik.setFieldValue(`yMin.isInvalid`, false)}
          />
          <Input
            name="yMax"
            type="number"
            className="px-2 py-1 ml-2"
            placeholder="Max value"
            value={formik.values.yMax.value}
            onChange={(e) =>
              formik.setFieldValue("yMax.value", parseInt(e.target.value))
            }
            onBlur={(event) => handleBlurChange(event.target)}
            onFocus={() => formik.setFieldValue(`yMax.isInvalid`, false)}
          />
        </FormGroup>
        <Button
          color="primary"
          size="sm"
          // outline
          className="mb-3 ml-3"
          onClick={() => handleRangeChangeBtn(getRequestBody(formik.values))}
          disabled={disbleApplyBtn(formik.values, progressStatus, isExpanded)}
        >
          Apply
        </Button>
        <Button
          color="primary"
          size="sm"
          // outline
          className="mb-3 ml-3"
          onClick={handleResetBtn}
          disabled={disbleApplyBtn(formik.values, progressStatus, isExpanded)}
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
