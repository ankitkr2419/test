import React from "react";

import styled from "styled-components";
// import { Link } from 'react-router-dom';
import {} from "core-components";
import { Icon, Text } from "shared-components";
import PropTypes from "prop-types";

const StyledPagination = styled.div`
  .pagination-box {
    width: 174px;
    height: 19px;
    .navigation-arrows {
      > i {
        color: #bebebe;
      }
    }
  }
`;

const PaginationBox = (props) => {
  const {firstIndexOnPage, lastIndexOnPage, totalPages, handlePrev, handleNext } = props;
  return (
    <StyledPagination>
      <div className="pagination-box d-flex justify-content-between align-items-center">
        <Text>
          <Text Tag="span" className="font-weight-bold">
            {`${firstIndexOnPage || 0}-${lastIndexOnPage || 0}`}
          </Text>
          <Text Tag="span"> of </Text>
          <Text Tag="span">{totalPages || 0}</Text>
        </Text>
        <Text className="navigation-arrows">
          <span onClick={handlePrev} >
              <Icon name="angle-left" size={30} />
          </span>
          <span onClick={handleNext} >
              <Icon name="angle-right" size={30} className="ml-3" />
          </span>        
        </Text>
      </div>
    </StyledPagination>
  );
};

PaginationBox.propTypes = {
  isUserLoggedIn: PropTypes.bool,
};

PaginationBox.defaultProps = {
  isUserLoggedIn: false,
};

export default PaginationBox;
