import React from "react";

import styled from "styled-components";
import { Input } from "core-components";
import { Icon } from "shared-components";
import PropTypes from "prop-types";

const StyledSearchBox = styled.div`
  .search-box {
    margin-bottom: auto;
    margin-top: auto;
    height: 42px;
    background-color: #fff;
    border-radius: 30px;
    padding: 1px;
    border: 1px solid #717171;
    .search-input {
      border: 0;
      outline: 0;
      background: none;
      width: 0;
      caret-color: transparent;
      line-height: 42px;
      transition: width 0.2s linear;
      padding: 0px 1px;
      &:focus {
        box-shadow: none;
      }
    }
    &:hover {
      .search-input {
        padding: 0 10px;
        width: 596px;
        caret-color: #000;
        transition: width 0.4s linear;
      }
      .search-icon {
        background: rgb(131, 180, 172);
        background: linear-gradient(
          -90deg,
          rgba(131, 180, 172, 1) 0%,
          rgba(178, 218, 209, 1) 100%
        );
        border-radius: 0 30px 30px 0;
        > i {
          color: #fff;
        }
      }
    }
    &:focus {
      box-shadow: none;
    }
    .search-icon {
      height: 38px;
      width: 38px;
      float: right;
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 50%;
      color: #abd5ce;
      text-decoration: none;
      &:hover {
        text-decoration: none;
      }
      > i {
        color: #abd5ce;
      }
    }
  }
`;

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
