import React from 'react';
import { Button, Select } from "core-components";
import Sidebar from "../Sidebar";
import SampleTargetList from "./SampleTargetList";

const SidebarSample = props => {
  return (
		<Sidebar className="sample" handleIcon="plus-1" handleIconSize={36}>
			<Select placeholder="Select Sample" className="mb-3" size="lg" />
			<SampleTargetList className="disabled" />
			<Select placeholder="Task - Unknown" className="mb-3" size="md" />
			<Button color="primary">Add</Button>
		</Sidebar>
	);
};

SidebarSample.propTypes = {
  
};

export default SidebarSample;