/* eslint-disable no-undef */
import React from 'react';
import PropTypes from 'prop-types';
import {
	Button, Popover, PopoverHeader, PopoverBody,
} from 'core-components';
import { Text, Center, ButtonIcon } from 'shared-components';

const WellPopover = (props) => {
	const {
		status,
		text,
		index,
		sample,
		task,
		targets,
		onEditClickHandler,
		...rest
	} = props;

	const simulateOutSideClick = () => document.body.click();

	return (
		<Popover
			trigger="legacy"
			target={`PopoverWell${index}`}
			hideArrow
			placement="top-start"
			popperClassName={`popover-well ${status}`}
			{...rest}
		>
			<PopoverHeader>
				<Text Tag="span">{text}</Text>
				<ButtonIcon
					position="absolute"
					placement="right"
					top={0}
					right={0}
					size={32}
					name="cross"
					className="btn-close"
					onClick={simulateOutSideClick}
				/>
			</PopoverHeader>
			<PopoverBody className="flex-100 scroll-y">
				<ul className="well-info flex-90 mx-auto mb-4 p-0">
					<li className="d-flex py-1">
						<Text className="flex-40 label mb-0">Sample</Text>
						<Text className="mb-0">{sample}</Text>
					</li>
					<li className="d-flex py-1">
						<Text className="flex-40 label mb-0">Target</Text>
						<div className="target-list">
							{targets === null && (
								<Text className={`mb-1 ${status}`}>---</Text>
							)}
							{targets !== null
                && targets.map(ele => (
                	<Text key={ele.target_id} className={`mb-1 ${status}`}>
                		{ele.target_id}
                	</Text>
                ))}
						</div>
					</li>
					<li className="d-flex py-1">
						<Text className="flex-40 label mb-0">Task</Text>
						<Text className="mb-0">{task}</Text>
					</li>
					<li className="d-flex py-1">
						<Text className="flex-40 label mb-0">Comment</Text>
						<Text className="mb-0">(No comments)</Text>
					</li>
				</ul>
				<Center>
					<Button className="mb-4">Show on Graph</Button>
					<Button onClick={onEditClickHandler}>Edit Info</Button>
				</Center>
			</PopoverBody>
		</Popover>
	);
};

WellPopover.propTypes = {
	status: PropTypes.string,
	text: PropTypes.string,
	index: PropTypes.number,
	sample: PropTypes.string,
	task: PropTypes.string,
	targets: PropTypes.array,
	onEditClickHandler: PropTypes.func,
};

export default WellPopover;
