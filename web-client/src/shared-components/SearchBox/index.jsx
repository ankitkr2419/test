import React from "react";

import { Input } from "core-components";
import { Icon } from "shared-components";
import PropTypes from "prop-types";
import { StyledSearchBox } from './StyledSearchBox';

const SearchBox = (props) => {
  const { value, onChange} = props;

  return (
    <StyledSearchBox className="mb-3">
      <div className="d-flex h-100">
        <div className="d-flex search-box">
          <Input
            className="search-input"
            type="text"
            name=""
            value={value}
            onChange={onChange}
            placeholder="Search..."
          />
          <div className="search-icon">
            <Icon name="search" size={32} />
          </div>
        </div>
      </div>
    </StyledSearchBox>
  );
};

SearchBox.propTypes = {
  isUserLoggedIn: PropTypes.bool,
};

SearchBox.defaultProps = {
  isUserLoggedIn: false,
};

export default React.memo(SearchBox);
