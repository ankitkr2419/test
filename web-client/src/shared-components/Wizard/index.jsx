import React, { useMemo } from "react";
import { Step, StepItem, StepLink } from "shared-components/StepBar";
import PropTypes from "prop-types";

/**
 * Wizard is used to show list with wizard items
 * @param {*} props
 */

const Wizard = (props) => {
  const { list, onClickHandler, showFinishBtn, finishBtnHandler } = props;

  const widgetList = useMemo(
    () => (
      <>
        {list.map((step) => (
          <StepItem key={step.get("id")} disabled={step.get("isDisable")}>
            <StepLink
              onClick={() => {
                onClickHandler(step.get("id"));
              }}
            >
              {step.get("name")}
            </StepLink>
          </StepItem>
        ))}
        {showFinishBtn && (
          <li className="ml-auto">
            <button
              className="btn btn-primary"
              onClick={finishBtnHandler}
              style={{ color: "#707070", backgroundColor: "#fff" }}
              //TODO move styles to class
            >
              Finish
            </button>
          </li>
        )}
      </>
    ),
    [list, onClickHandler]
  );

  return <Step>{widgetList}</Step>;
};

Wizard.propTypes = {
  list: PropTypes.shape({
    id: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
    isDisable: PropTypes.bool,
    name: PropTypes.string,
  }),
  onClickHandler: PropTypes.func.isRequired,
};

export default React.memo(Wizard);
