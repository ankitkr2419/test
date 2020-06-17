import React from 'react';
import PropTypes from 'prop-types';
import {
	StyledUl,
	StyledLi,
	CustomButton,
	ButtonGroup,
} from 'shared-components';
import CreateTemplateModal from './CreateTemplateModal';
// import { Button } from 'core-components';

const TemplateComponent = (props) => {
	const { templates } = props;

	return (
		<div className="d-flex flex-100 flex-column p-4 mt-3">
			<StyledUl>
				{templates.map(template => (
					<StyledLi key={template.get('id')}>
						<CustomButton
							title={template.get('name')}
							isEditable
							onEditClickHandler={() => {}}
							isDeletable
							onDeleteClickHandler={() => {}}
						/>
					</StyledLi>
				))}
			</StyledUl>
			<ButtonGroup className="text-center">
				{/*
          TODO Handle login flow when operator
          <Button color="primary">Next</Button>
        */}
				<CreateTemplateModal />
			</ButtonGroup>
		</div>
	);
};

TemplateComponent.propTypes = {
	templates: PropTypes.shape({}).isRequired,
};

export default React.memo(TemplateComponent);
