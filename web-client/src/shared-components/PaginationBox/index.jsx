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
  return (
    <StyledPagination>
      <div className="pagination-box d-flex justify-content-between align-items-center">
        <Text>
          <Text Tag="span" className="font-weight-bold">
            1-24
          </Text>
          <Text Tag="span"> of </Text>
          <Text Tag="span">24</Text>
        </Text>
        <Text className="navigation-arrows">
          <Icon name="angle-left" size={30} />
          <Icon name="angle-right" size={30} className="ml-3" />
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
