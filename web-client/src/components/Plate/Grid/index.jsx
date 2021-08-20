import React from "react";
import PropTypes from "prop-types";
import WellGrid from "./WellGrid";
import Well from "./Well";
import WellPopover from "./WellPopover";
import { EXPERIMENT_STATUS } from "appConstants";

const GridComponent = ({
  wells,
  onWellClickHandler,
  onWellUpdateClickHandler,
  isGroupSelectionOn,
  showGraphOfWell,
  experimentStatus,
}) => (
  <div className="d-flex flex-column flex-100 pt-4">
    <WellGrid className="rtpcr-well-grid">
      {wells.map((well, index) => {
        if (well !== null) {
          const {
            isWellFilled,
            isRunning,
            isSelected,
            isMultiSelected,
            status,
            initial,
            text,
            sample,
            task,
            targets,
            isWellActive,
          } = well.toJS();
          return (
            <>
              <Well
                key={index}
                isRunning={isRunning}
                isSelected={isSelected || isMultiSelected}
                status={status}
                taskInitials={initial}
                id={`PopoverWell${index}`}
                onClickHandler={(event) => {
                  onWellClickHandler(well, index, event);
                }}
                isDisabled={
                  isWellFilled === false &&
                  (experimentStatus === EXPERIMENT_STATUS.success ||
                    experimentStatus === EXPERIMENT_STATUS.running ||
                    experimentStatus === EXPERIMENT_STATUS.stopped)
                }
                position={index}
              />
              {/* popover will only visible when its filled and group selection is off */}
              {isWellFilled && isGroupSelectionOn === false && (
                <WellPopover
                  sample={sample}
                  status={status}
                  text={text}
                  task={task}
                  targets={targets}
                  index={index}
                  onEditClickHandler={(event) => {
                    onWellUpdateClickHandler(well, index, event);
                  }}
                  showGraphOfWell={showGraphOfWell}
                  isEditBtnDisabled={
                    experimentStatus === EXPERIMENT_STATUS.success ||
                    experimentStatus === EXPERIMENT_STATUS.running ||
                    experimentStatus === EXPERIMENT_STATUS.stopped
                  }
                />
              )}
            </>
          );
        }
        return null;
      })}
    </WellGrid>
  </div>
);

GridComponent.propTypes = {
  onWellClickHandler: PropTypes.func.isRequired,
  wells: PropTypes.object.isRequired,
  onWellUpdateClickHandler: PropTypes.func.isRequired,
  isGroupSelectionOn: PropTypes.bool.isRequired,
  experimentStatus: PropTypes.string,
};

export default GridComponent;
