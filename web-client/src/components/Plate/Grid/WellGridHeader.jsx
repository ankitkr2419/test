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

const WellGridHeader = ({ className }) => (
	<StyledWellGridHeader className={className}>
		<Switch
			id="selection"
			name="selection"
			label="Group Selection to view graph"
		/>
	</StyledWellGridHeader>
);

WellGridHeader.propTypes = {
	className: PropTypes.string,
};

WellGridHeader.defaultProps = {
	className: '',
};

export default WellGridHeader;
