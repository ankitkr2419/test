import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { ExperimentGraphContainer } from 'containers/ExperimentGraphContainer';
import SampleSideBarContainer from 'containers/SampleSideBarContainer';
import Header from './Header';

import GridWrapper from './Grid/GridWrapper';
import GridComponent from './Grid';
import WellGridHeader from './Grid/WellGridHeader';

import './Plate.scss';

const Plate = (props) => {
	const {
		wells,
		setSelectedWell,
		setMultiSelectedWell,
		experimentTargetsList,
		positions,
		experimentId,
		isMultiSelectionOptionOn,
		toggleMultiSelectOption,
		activeWells,
	} = props;

	// local state to maintain well data which is selected for update
	const [updateWell, setUpdateWell] = useState(null);
	const onWellClickHandler = (well, index) => {
		const { isSelected, isWellFilled } = well.toJS();
		// if well is blank and multi-select checkbox is un-checked
		if (isMultiSelectionOptionOn === false && isWellFilled === false) {
			setSelectedWell(index, !isSelected);
		}

		// if multi-select checkbox is checked, will allow to select filled wells
		if (isMultiSelectionOptionOn === true) {
			if (isWellFilled === true) {
				setMultiSelectedWell(index, !isSelected);
			}
		}
	};

	const onWellUpdateClickHandler = (well, index) => {
		setUpdateWell(well.toJS());
	};

	return (
		<div className="plate-content d-flex flex-column h-100 position-relative">
			<Header />
			<GridWrapper className="plate-body flex-100">
				<WellGridHeader
					isGroupSelectionOn={isMultiSelectionOptionOn}
					toggleMultiSelectOption={toggleMultiSelectOption}
				/>
				<GridComponent
					activeWells={activeWells}
					wells={wells}
					isGroupSelectionOn={isMultiSelectionOptionOn}
					onWellClickHandler={onWellClickHandler}
					onWellUpdateClickHandler={onWellUpdateClickHandler}
				/>
			</GridWrapper>
			<SampleSideBarContainer
				experimentId={experimentId}
				positions={positions}
				experimentTargetsList={experimentTargetsList}
				updateWell={updateWell}
			/>
			<ExperimentGraphContainer
				experimentId={experimentId}
				experimentTargetsList={experimentTargetsList}
			/>
		</div>
	);
};

Plate.propTypes = {
	wells: PropTypes.object.isRequired,
	setSelectedWell: PropTypes.func.isRequired,
	setMultiSelectedWell: PropTypes.func.isRequired,
	// experimentTargetsList contains targets for selected experiment
	experimentTargetsList: PropTypes.object.isRequired,
	// array of selected wells with index
	positions: PropTypes.object.isRequired,
	experimentId: PropTypes.string.isRequired,
	isMultiSelectionOptionOn: PropTypes.bool.isRequired,
	toggleMultiSelectOption: PropTypes.func.isRequired,
	activeWells: PropTypes.object.isRequired,
};

export default Plate;
