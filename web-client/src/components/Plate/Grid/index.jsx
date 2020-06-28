import React from 'react';
import PropTypes from 'prop-types';
import Coordinate from './Coordinate';
import CoordinateItem from './CoordinateItem';
import WellGrid from './WellGrid';
import Well from './Well';
import WellPopover from './WellPopover';
import PlateData from './PlateData.json';

const { YCoordinates } = PlateData;
const { XCoordinates } = PlateData;
// const { Wells } = PlateData;

const GridComponent = ({ onWellClickHandler, wells }) => (
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
						isRunning, isSelected, status, initial, text,
					} = well.toJS();
					return (
						<div key={index}>
							<Well
								isRunning={isRunning}
								isSelected={isSelected}
								status={status}
								taskInitials={initial}
								id={`PopoverWell${index}`}
								onClickHandler={(event) => { onWellClickHandler(well, index, event); }}
							/>
							{isRunning && <WellPopover status={status} text={text} index={index} />}
						</div>
					);
				})}
			</WellGrid>
		</div>
	</div>
);

GridComponent.propTypes = {
	onWellClickHandler: PropTypes.func.isRequired,
};

export default GridComponent;
