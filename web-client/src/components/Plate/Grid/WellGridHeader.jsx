import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { Switch } from 'core-components';

const StyledWellGridHeader = styled.header`
	display: flex;
	height: 40px;
	align-items: center;
	padding: 0 16px 0 26px;
`;

const WellGridHeader = ({ className, isGroupSelectionOn, toggleMultiSelectOption }) => (
	<StyledWellGridHeader className={className}>
		<Switch
			id="selection"
			name="selection"
			label="Group Selection to view graph"
			value={isGroupSelectionOn}
			onChange={toggleMultiSelectOption}
		/>
	</StyledWellGridHeader>
);

WellGridHeader.propTypes = {
	className: PropTypes.string,
	isGroupSelectionOn: PropTypes.bool.isRequired,
	toggleMultiSelectOption: PropTypes.func.isRequired,
};

WellGridHeader.defaultProps = {
	className: '',
};

export default WellGridHeader;
