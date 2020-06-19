import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import styled, { css } from 'styled-components';

const StyledLink = styled(Link)`
  display: ${props => (props.icon === 'true' ? 'flex' : 'inline-block')};
  width: ${props => (props.icon === 'true' ? '40px' : '202px')};
  height: 40px;
  font-size: 16px;
  line-height: 19px;
  font-weight: bold;
  text-align: center;
  vertical-align: middle;
  padding: ${props => (props.icon === 'true' ? '4px' : '10px 20px')};
  border-radius: ${props => (props.icon === 'true' ? '50%' : '27px')};
  border-width: ${props => (props.icon === 'true' ? '' : '1px')};
  border-style: ${props => (props.icon === 'true' ? '' : 'solid')};
  box-shadow: ${props => (props.icon === 'true' ? '' : '0 2px 6px #00000029')};
  user-select: none;
  transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out,
    border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;

  &:focus {
    outline: none;
    box-shadow: ${props => (props.icon === 'true' ? '' : '0 2px 6px #00000029')};
  }

  &:hover {
    text-decoration: none;
  }

  ${props => props.icon
    && css`
      align-items: center;
      justify-content: center;

      i {
        color: #707070;
      }
    `};
`;

const CustomLink = (props) => {
	const { icon, ...rest } = props;
	return (
		<StyledLink icon={icon.toString()} {...rest}>
			{props.children}
		</StyledLink>
	);
};

CustomLink.propTypes = {
	icon: PropTypes.bool,
};

CustomLink.defaultProps = {
	icon: false,
};

export default CustomLink;
