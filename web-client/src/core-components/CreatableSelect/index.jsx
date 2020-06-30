import React from 'react';
import PropTypes from 'prop-types';
import  Select from 'react-select/creatable';

const customStyles = {
	singleValue: (provided, state) => {
		const color = 'rgba(112, 112, 112, 1)';
		const fontSize = '18px';

		return { ...provided, color, fontSize };
	},

	placeholder: (provided, state) => {
		const color = 'rgba(112, 112, 112, 0.48)';
		const fontSize = '18px';

		return { ...provided, color, fontSize };
	},
};

const CreatableSelect = ({ className, size, ...rest }) => (
	<Select
		menuPosition='fixed'
		styles={customStyles}
		className={`${className} ${size}`}
		size={size}
		{...rest}
	/>
);

CreatableSelect.propTypes = {
	className: PropTypes.string,
	size: PropTypes.oneOf(['lg', 'md', '']),
};

CreatableSelect.defaultProps = {
	className: '',
	size: '',
};

export default CreatableSelect;
