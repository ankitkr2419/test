import React from 'react';
import Header from './Header';

import GridWrapper from './Grid/GridWrapper';
import GridComponent from './Grid';
import WellGridHeader from './Grid/WellGridHeader';

import SidebarSample from './Sidebar/AddSample/SidebarSample';
import SidebarGraph from './Sidebar/Graph/SidebarGraph';
import './Plate.scss';

const Plate = (props) => {
	const onWellClickHandler = (well) => {
		console.log('well clicked', well);
		// if well is blank
		// then just select it

		// if well is filled with data
		// then open pop-over

		// if multi selection flag is on
		// then allow filled well to select
	};

	return (
		<div className="plate-content d-flex flex-column h-100 position-relative">
			<Header />
			<GridWrapper className="plate-body flex-100">
				<WellGridHeader />
				<GridComponent onWellClickHandler={onWellClickHandler}/>
			</GridWrapper>
			<SidebarSample />
			<SidebarGraph />
		</div>
	);
};

Plate.propTypes = {};

export default Plate;
