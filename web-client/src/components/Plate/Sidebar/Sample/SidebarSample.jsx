import React from 'react';
import PropTypes from 'prop-types';
import { Button, Select, CreatableSelect } from 'core-components';
import Sidebar from 'components/Sidebar';
import SampleTargetList from './SampleTargetList';

const SidebarSample = (props) => {
	const {
		sampleState,
		updateCreateSampleWrapper,
		fetchSamples,
		sampleOptions,
		addNewLocalSample,
		isSampleListLoading,
		taskOptions,
		onCrossClickHandler,
		addButtonClickHandler,
		isSampleStateValid, // form state valid
		isDisabled,
		resetLocalState,
	} = props;

	const {
		isSideBarOpen,
		sample,
		task,
		isEdit,
	} = sampleState.toJS();

	const toggleSideBar = () => {
		// if user close sidebar without editing then reset local state
		if (isSideBarOpen === true && isEdit === true) {
			resetLocalState();
		}
		updateCreateSampleWrapper('isSideBarOpen', !isSideBarOpen);
	};

	const handleSampleCreate = (inputValue) => {
		const newOption = {
			label: inputValue,
			value: inputValue,
		};
		// update local state
		updateCreateSampleWrapper('sample', newOption);
		// add new sample to sample's reducer to merge it with original list
		addNewLocalSample(newOption);
	};

	const handleSampleInputChange = (text) => {
		if (text.length >= 3) {
			fetchSamples(text);
		}
	};

	const handleSampleChange = (value) => {
		updateCreateSampleWrapper('sample', value);
	};

	const handleTaskChange = (value) => {
		updateCreateSampleWrapper('task', value);
	};

	return (
		<Sidebar
			isOpen={isSideBarOpen}
			toggleSideBar={toggleSideBar}
			className="sample"
			handleIcon="plus-1"
			handleIconSize={36}
			isDisabled={isDisabled}
		>
			<CreatableSelect
				isClearable
				isDisabled={isSampleListLoading}
				isLoading={isSampleListLoading}
				onChange={handleSampleChange}
				onCreateOption={handleSampleCreate}
				onInputChange={handleSampleInputChange}
				options={sampleOptions}
				value={sample}
				placeholder="Select Sample"
				className="mb-3"
				size="lg"
			/>
			<SampleTargetList
				list={sampleState.get('targets')}
				onCrossClickHandler={onCrossClickHandler}
			/>
			<Select
				placeholder="Task - Unknown"
				className="mb-3"
				size="md"
				options={taskOptions}
				value={task}
				onChange={handleTaskChange}
			/>
			<Button disabled={isSampleStateValid === false} onClick={addButtonClickHandler} color="primary">Add</Button>
		</Sidebar>
	);
};

SidebarSample.propTypes = {
	sampleState: PropTypes.object.isRequired,
	updateCreateSampleWrapper: PropTypes.func.isRequired,
	fetchSamples: PropTypes.func.isRequired,
	sampleOptions: PropTypes.array.isRequired,
	addNewLocalSample: PropTypes.func.isRequired,
	isSampleListLoading: PropTypes.bool.isRequired,
	taskOptions: PropTypes.array.isRequired,
	onCrossClickHandler: PropTypes.func.isRequired,
	addButtonClickHandler: PropTypes.func.isRequired,
	isSampleStateValid: PropTypes.bool.isRequired,
	isDisabled: PropTypes.bool.isRequired,
	resetLocalState: PropTypes.func.isRequired,
};

export default React.memo(SidebarSample);
