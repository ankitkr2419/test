import React from "react";
import Select from "react-select";

const customStyles = {
	singleValue: (provided, state) => {
		const color = "rgba(112, 112, 112, 1)";
		const fontSize = "18px";

		return { ...provided, color, fontSize };
	},

	placeholder: (provided, state) => {
		const color = "rgba(112, 112, 112, 0.48)";
		const fontSize = "18px";

		return { ...provided, color, fontSize };
	},
};

const StyledSelect = (props) => {
	return (
		<Select
			styles={customStyles}
			className={`ml-select ${props.wrapperClassName}`}
			{...props}
		/>
	);
};

StyledSelect.propTypes = {};

export default StyledSelect;