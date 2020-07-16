import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { useSelector } from 'react-redux';
import { ExperimentGraphContainer } from 'containers/ExperimentGraphContainer';
import { getRunExperimentReducer } from 'selectors/runExperimentSelector';
import SampleSideBarContainer from 'containers/SampleSideBarContainer';
import Header from './Header';

import GridWrapper from './Grid/GridWrapper';
import GridComponent from './Grid';
import WellGridHeader from './Grid/WellGridHeader';

import './Plate.scss';
import { EXPERIMENT_STATUS } from 'appConstants';

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
		experimentTemplate,
	} = props;

	// getExperimentStatus will return us current experiment status
	const runExperimentDetails = useSelector(getRunExperimentReducer);
	const experimentStatus = runExperimentDetails.get('experimentStatus');
	const experimentDetails = runExperimentDetails.get('data');

	// local state to maintain well data which is selected for updation
	const [updateWell, setUpdateWell] = useState(null);

	/**
	 *
	 * @param {*} well => selected well details
	 * @param {*} index => selected well index
	 *
	 */
	const onWellClickHandler = (well, index) => {
		const { isSelected, isWellFilled, isMultiSelected } = well.toJS();
		/**
		 * if well is not filled and if multi selection option is not checked
		 * 				then we can make well selected
		 */
		if (isMultiSelectionOptionOn === false && isWellFilled === false) {
			setSelectedWell(index, !isSelected);
		}

		/**
		 * if multi-select checkbox is checked, will allow to select filled wells
		 */
		if (isMultiSelectionOptionOn === true) {
			// if (isWellFilled === true) {
			setMultiSelectedWell(index, !isMultiSelected);
			// }
		}
	};

	const onWellUpdateClickHandler = (selectedWell) => {
		// update local state with selected well which is selected for updation
		setUpdateWell(selectedWell.toJS());
	};

	return (
		<div className="plate-content d-flex flex-column h-100 position-relative">
			<Header
				experimentTemplate={experimentTemplate}
				isExperimentSucceeded={experimentStatus === EXPERIMENT_STATUS.success}
				experimentDetails={experimentDetails}
			/>
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
			<ExperimentGraphContainer experimentStatus={experimentStatus} />
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
	experimentTemplate: PropTypes.object.isRequired,
};

export default Plate;
