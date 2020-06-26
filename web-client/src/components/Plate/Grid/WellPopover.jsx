import React from 'react';
import {
	Button,
	Popover,
	PopoverHeader,
	PopoverBody,
} from 'core-components';
import { Text, Center, ButtonIcon } from 'shared-components';

const WellPopover = (props) => {
	const {
		status, text, index, ...rest
	} = props;
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
				/>
			</PopoverHeader>
			<PopoverBody className="flex-100 scroll-y">
				<ul className="well-info flex-90 mx-auto mb-4 p-0">
					<li className="d-flex py-1">
						<Text className="flex-40 label mb-0">Sample</Text>
						<Text className="mb-0">ID-xx-xxx</Text>
					</li>
					<li className="d-flex py-1">
						<Text className="flex-40 label mb-0">Target</Text>
						<div className="target-list">
							<Text className={`mb-1 ${status}`}>Target 1</Text>
							<Text className={`mb-1 ${status}`}>Target 2</Text>
							<Text className={`mb-1 ${status}`}>Target 6</Text>
						</div>
					</li>
					<li className="d-flex py-1">
						<Text className="flex-40 label mb-0">Task</Text>
						<Text className="mb-0">Unknown</Text>
					</li>
					<li className="d-flex py-1">
						<Text className="flex-40 label mb-0">Comment</Text>
						<Text className="mb-0">(If any)</Text>
					</li>
				</ul>
				<Center>
					<Button className="mb-4">Show on Graph</Button>
					<Button>Edit Info</Button>
				</Center>
			</PopoverBody>
		</Popover>
	);
};

WellPopover.propTypes = {};

export default WellPopover;
