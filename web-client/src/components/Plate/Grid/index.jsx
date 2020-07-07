import React from 'react';
import PropTypes from 'prop-types';
import { YCoordinates, XCoordinates} from 'components/Plate/plateConstant';
import Coordinate from './Coordinate';
import CoordinateItem from './CoordinateItem';
import WellGrid from './WellGrid';
import Well from './Well';
import WellPopover from './WellPopover';

const GridComponent = ({
	wells,
	onWellClickHandler,
	onWellUpdateClickHandler,
	isGroupSelectionOn,
}) => (
	<div className="d-flex flex-column flex-100">
		<Coordinate direction="horizontal">
			{YCoordinates.map((value, i) => (
				<CoordinateItem key={i} coordinateValue={value.toString()} />
			))}
		</Coordinate>
		<div className="d-flex flex-100">
			<Coordinate>
				{XCoordinates.map((value, i) => (
					<CoordinateItem key={i} coordinateValue={value} />
				))}
			</Coordinate>
			<WellGrid>
				{wells.map((well, index) => {
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
						<React.Fragment key={index}>
							<Well
								isRunning={isRunning}
								isSelected={isSelected || isMultiSelected}
								status={status}
								taskInitials={initial}
								id={`PopoverWell${index}`}
								onClickHandler={(event) => {
									onWellClickHandler(well, index, event);
								}}
								isDisabled={isWellActive === false}
							/>
							{/* popover will only visible when its filled and group selection is off */}
							{(isWellFilled && isGroupSelectionOn === false) && (
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
								/>
							)}
						</React.Fragment>
					);
				})}
			</WellGrid>
		</div>
	</div>
);

GridComponent.propTypes = {
	onWellClickHandler: PropTypes.func.isRequired,
	wells: PropTypes.object.isRequired,
	onWellUpdateClickHandler: PropTypes.func.isRequired,
	isGroupSelectionOn: PropTypes.bool.isRequired,
};

export default GridComponent;
