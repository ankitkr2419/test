import React, { useState, useEffect } from 'react';
import { takeEvery, put, call } from 'redux-saga/effects';
import { Button, Select, CreatableSelect } from 'core-components';
import { useDispatch } from 'react-redux';
import Sidebar from 'components/Sidebar';
import { getResponse } from 'apis/apiHelper';
import SampleTargetList from './SampleTargetList';

const taskOptions = [
	{ label : 'Unknown', value: 'Unknown' },
	{ label : 'NC', value: 'NC' },
	{ label : 'PC', value: 'PC' },
	{ label : 'NTC', value: 'NTC' },
];

const SidebarSample = ({ experimentTargetsList }) => {
	const dispatch = useDispatch();
	const [isSidebarOpen, setSideBarState] = useState(false);
	const [targetList, updateTargetList] = useState(experimentTargetsList);

	useEffect(() => {
		if (experimentTargetsList !== null && experimentTargetsList.size !== 0) {
			updateTargetList(experimentTargetsList);
		}
	}, [experimentTargetsList, updateTargetList]);

	const toggleSideBar = () => {
		setSideBarState(isOpen => !isOpen);
	};

	const getSampleOptions = () => getResponse('samples/abc');

	useEffect(() => {
		// console.log(getSampleOptions().then(response => {
		// 	console.log('response: ', response);
		// }));
	}, []);

	const onCrossClickHandler = (index) => {
		updateTargetList(targetList.delete(index));
	};

	return (
		<Sidebar
			isOpen={isSidebarOpen}
			toggleSideBar={toggleSideBar}
			className="sample"
			handleIcon="plus-1"
			handleIconSize={36}
		>
			<CreatableSelect
				// isClearable
				// cacheOptions
				// loadOptions={getSampleOptions}
				// isDisabled={isLoading}
				// isLoading={isLoading}
				// onChange={handleChange}
				// onCreateOption={handleCreate}
				// options={options}
				// value={value}
				placeholder="Select Sample"
				className="mb-3"
				size="lg"
			/>
			<SampleTargetList
				list={targetList}
				onCrossClickHandler={onCrossClickHandler}
			/>
			<Select placeholder="Task - Unknown" className="mb-3" size="md" options={taskOptions} />
			<Button color="primary">Add</Button>
		</Sidebar>
	);
};

SidebarSample.propTypes = {};

export default SidebarSample;
