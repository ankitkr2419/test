import React from "react";
import PlateBody from "./PlateBody";
import PlateSubHeader from "./PlateSubHeader";
import Coordinate from "./Coordinate";
import CoordinateItem from "./CoordinateItem";
import WellGridHeader from "./WellGridHeader";
import WellGrid from "./WellGrid";
import Well from "./Well";
import WellPopover from "./WellPopover";
import SidebarSample from "./SidebarSample";
import SidebarGraph from "./SidebarGraph";
import "./Plate.scss";
import PlateData from "./PlateData";

const YCoordinates = PlateData.YCoordinates;
const XCoordinates = PlateData.XCoordinates;
const Wells = PlateData.Wells;

const Plate = (props) => {
	return (
		<div className="plate-content d-flex flex-column h-100 position-relative">
			<PlateSubHeader />
			<PlateBody className="plate-body flex-100">
				<WellGridHeader />
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
									/>
									<WellPopover status={well.status} text={well.text} index={i} />
								</div>
							))}
						</WellGrid>
					</div>
				</div>
			</PlateBody>
			<SidebarSample />
			<SidebarGraph />
		</div>
	);
};

Plate.propTypes = {};

export default Plate;
