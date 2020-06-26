import React, { useState } from 'react';
import { Button, Select } from 'core-components';
import Sidebar from 'components/Sidebar';
import SampleTargetList from './SampleTargetList';

const SidebarSample = () => {
	const [isSidebarOpen, setSideBarState] = useState(false);
	const toggleSideBar = () => {
		setSideBarState((isOpen) => !isOpen);
	};

	return (
		<Sidebar
			isOpen={isSidebarOpen}
			toggleSideBar={toggleSideBar}
			className='sample'
			handleIcon='plus-1'
			handleIconSize={36}
		>
			<Select placeholder='Select Sample' className='mb-3' size='lg' />
			<SampleTargetList className='disabled' />
			<Select placeholder='Task - Unknown' className='mb-3' size='md' />
			<Button color='primary'>Add</Button>
		</Sidebar>
	);
};

SidebarSample.propTypes = {};

export default SidebarSample;
