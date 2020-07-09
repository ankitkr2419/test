import React from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import styled from 'styled-components';

const Well = ({
	id,
	className,
	status,
	isRunning,
	isSelected,
	taskInitials,
	onClickHandler,
	isDisabled,
}) => {
	const wellClassnames = classNames(className, {
		running: isRunning,
		selected: isSelected,
	});

	return (
		<StyledWell
			id={id}
			isRunning={isRunning}
			isSelected={isSelected}
			status={status}
			className={wellClassnames}
			onClick={onClickHandler}
			isDisabled={isDisabled}
		>
			{taskInitials}
		</StyledWell>
	);
};

const getBackgroundColor = ({ isSelected, isRunning, isDisabled }) => {
	if (isSelected && !isRunning) {
		return '#aedbd5';
	}

	if (isDisabled) {
		return 'gray';
	}

	return '#ffffff';
};

const StyledWell = styled.div`
  background-color: ${props => getBackgroundColor(props)};
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 57px;
  height: 54px;
  font-size: 20px;
  line-height: 24px;
  color: #666666;
  border: ${props => (props.isSelected && props.isRunning
		? '2px solid #707070'
		: '1px solid #aeaeae')};
  border-radius: 8px;
  margin: 0 16px 16px 0;
  padding: 20px 4px 4px;
  box-shadow: ${props => (props.isSelected && props.isRunning ? '0 3px 6px #00000029' : '')};
	opacity: ${props => (props.isDisabled ? '0.2' : '1')};
	pointer-events: ${props => (props.isDisabled ? 'none' : 'auto')};
	
  &.selected {
    &:active,
    &:active:focus {
      background-color: #aedbd5;
    }
  }

  &.running {
    &:active,
    &:active:focus {
      background-color: #ffffff;
      border: 2px solid #707070;
      box-shadow: 0 3px 6px #00000029;
    }
  }

  &:active,
  &:active:focus {
    background-color: #aedbd5;
  }

  &:focus {
    outline: none;
  }

  &::before {
    content: "";
    position: absolute;
    top: 0;
    right: 0;
    left: 0;
    height: 16px;
    border-radius: 6px 6px 0 0;
    background-color: ${props => props.status}
  }
`;

Well.propTypes = {
	id: PropTypes.string,
	className: PropTypes.string,
	status: PropTypes.string,
	taskInitials: PropTypes.string,
	isSelected: PropTypes.bool,
	isRunning: PropTypes.bool,
	onClickHandler: PropTypes.func.isRequired,
	isDisabled: PropTypes.bool,
};

Well.defaultProps = {
	id: '',
	className: '',
	status: '',
	taskInitials: '',
	isSelected: false,
	isRunning: false,
	isDisabled: false,
};

export default Well;