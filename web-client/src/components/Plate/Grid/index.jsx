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
const { Wells } = PlateData;

const GridComponent = ({ onWellClickHandler }) => (
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
				{Wells.map((well, i) => (
					<div key={i}>
						<Well
							isRunning={well.isRunning}
							isSelected={well.isSelected}
							status={well.status}
							taskInitials={well.text}
							id={`PopoverWell${i}`}
							onClickHandler={(event) => { onWellClickHandler(well, i, event); }}
						/>
						{well.isRunning && <WellPopover status={well.status} text={well.text} index={i} />}
					</div>
				))}
			</WellGrid>
		</div>
	</div>
);

GridComponent.propTypes = {
	onWellClickHandler: PropTypes.func.isRequired,
};

export default GridComponent;
