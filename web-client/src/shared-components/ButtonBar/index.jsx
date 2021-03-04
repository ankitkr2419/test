import React from 'react';

import styled from 'styled-components';
import PropTypes from 'prop-types';
import { Text, Icon, Link } from 'shared-components';
import {
	Button, 
} from 'core-components';


const ButtonBarBox = styled.div`
  width: 100%;
  height: 3.25rem;
  background-color:#fff;
  border-radius:32px 0 0 32px;
  padding:8px 79px 8px 38px;
  box-shadow:0px 3px 16px rgba(0,0,0,0.6);
`;
const PrevBtn = styled.div`
min-width:inherit;
border:0;
box-shadow:none;
color:#F38220;
`;

const ButtonBar = (props) => {
	return (
		<ButtonBarBox className="d-flex justify-content-start align-items-center mt-5">
      <PrevBtn><Icon name='angle-left' size={30} /></PrevBtn>
        <Button
            color="outline-primary"
            className="ml-auto text-dark"
						size="md"
        >	<Icon size={20} name='plus-3' className="mr-1 text-primary"/>Add Process       
        </Button>
        <Button
            color="primary"
            className="ml-4"
            size="md"
        >	Finish       
      </Button>
			
		</ButtonBarBox>
	);
};

ButtonBar.propTypes = {
	isUserLoggedIn: PropTypes.bool,
};

ButtonBar.defaultProps = {
	isUserLoggedIn: false,
};

export default ButtonBar;
