import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { FormGroup, Input } from 'core-components';
import { Icon } from 'shared-components';

const StyledSearchBar = styled(FormGroup)`
	position: relative;
	margin-bottom: -1px;

	.form-control {
		padding: 4px 51px 4px 16px;
		border-color: #cbcbcb;
		border-radius: 0;

		&:focus {
			box-shadow: none;
		}
	}

	i {
		position: absolute;
		top: 3px;
		right: 10px;
		color: #666666;
	}
`;

const SearchBar = (props) => {
	const { className, id, name, placeholder, ...rest } = props;
	return (
		<StyledSearchBar className={`search-bar ${className} mt-3`} {...rest}>
			<Input type='text' id={id} name={name} placeholder={placeholder} />
			<Icon name='search' size={32} />
		</StyledSearchBar>
	);
};

SearchBar.propTypes = {
	className: PropTypes.string,
	id: PropTypes.string,
	name: PropTypes.string,
	placeholder: PropTypes.string,
};

SearchBar.defaultProps = {
	className: '',
};

export default SearchBar;
