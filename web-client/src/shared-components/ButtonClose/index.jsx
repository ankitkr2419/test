import React from 'react';
import Icon from "shared-components/Icon";
import ButtonIcon from "shared-components/ButtonIcon";

const ButtonClose = props => {
  return (
		<ButtonIcon {...props}>
			<Icon size="32" name="cross" />
		</ButtonIcon>
	);
};

ButtonClose.propTypes = {};

export default ButtonClose;