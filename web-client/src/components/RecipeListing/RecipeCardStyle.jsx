import styled from "styled-components";

export const RecipeCardStyle = styled.div`
  padding: 0.8rem 0.5rem;
  border: 1px solid #e3e3e3;
  border-radius: 0.5rem;
  margin-bottom: 0.688rem;
  box-shadow: 0px 3px 16px rgba(0, 0, 0, 0.04);
  // width:27.5rem;
  // height: 5.563rem;
  .recipe-heading {
    padding-bottom: 0.5rem;
  }
  .recipe-card-body {
    padding-top: 0.25rem;
    border-top: 1px solid #d9d9d9;

    .recipe-name {
      font-size: 0.875rem;
      line-height: 1rem;
    }
    .recipe-value {
      font-size: 1.125rem;
      line-height: 1.313rem;
    }
    .recipe-action {
      border-top: 1px solid #d9d9d9 !important;
      button {
        width: 33px !important;
        height: 33px !important;
        border: 1px solid #696969 !important;
        &:not(:first-child) {
          margin-left: 12px;
        }
      }
    }
  }
  &:focus,
  &:hover {
    background-color: "white";
  }
  &.selected {
    background-color: rgba(243, 130, 32, 0.3);
  }

  &.admin-selected {
    position: relative;
    z-index: 4;
    background-color: #fff;
  }

  .recipe-action {
    border-top: 1px solid #d9d9d9 !important;
    button {
      width: 33px !important;
      height: 33px !important;
      border: 1px solid #696969 !important;
      &:not(:first-child) {
        margin-left: 12px;
      }
    }
  }
`;

export const CardOverlay = styled.div`
  position: absolute;
  // display: none;
  width: 100%;
  height: 505px;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #00000012;
  cursor: pointer;
  border-radius: 36px 36px 0px 0px;
`;
