import React from 'react';

import { 
	Col
} from 'core-components';
import {
	ButtonIcon,
	Text
} from 'shared-components';

import styled from 'styled-components';

const ProcessBox = styled.div`
margin-bottom:14px;
.process-card{
    border-radius:9px;
    border:1px solid #CCCCCC;
    padding:29px 34px;
    height:108px;
    .process-name{
        font-size:18px;
        line-height:21px;
        color: #666666;
    }
}
`;


const Process = ({iconName, processName}) => {
	return (
        <Col md={4}>
            <ProcessBox>
                <div className="process-card bg-white d-flex align-items-center frame-icon">
                    <ButtonIcon
                        size={51}
                        name={iconName}
                        className="border-dark-gray text-primary"
                        //onClick={toggleExportDataModal}
                    />
                    <Text Tag="span" className="ml-2 process-name">
                        {processName}
                    </Text>
                </div>
            </ProcessBox>
        </Col>
	);
};

Process.propTypes = {};

export default Process;
