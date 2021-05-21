import React from 'react';
import PropTypes from 'prop-types';
import { Input } from 'core-components';
import { Icon } from 'shared-components';
import { StyledSearchBar } from './StyledSearchBar'

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
