import styled from "styled-components";

export const PageBody = styled.div`
  background-color: #f5f5f5;
`;
export const TipPositionBox = styled.div`
  .process-tip-position {
    &::after {
      background: url("/images/tip-position-bg.svg") no-repeat;
    }
  }
`;

export const TopContent = styled.div`
  margin-bottom: 0.5rem;
  .frame-icon {
    > button {
      > i {
        font-size: 49px;
      }
    }
  }
`;

export const TipPositionInfoBox = styled.div`
  .process-box {
    width: 90.55%;
    label {
      font-size: 1rem;
      line-height: 1.125rem;
      color: #666666;
    }
    .well-box {
      // counter-reset: section;
      .well {
        &:active {
          background-color: #ffffff;
        }
      }
      .well-no {
        .coordinate-item {
          color: #999999;
          font-size: 1.125rem;
          line-height: 1.313rem;
        }
      }
      .well {
        margin: 0 1.313rem 1.313rem 0;
        width: 2.5rem;
        height: 2.5rem;
        font-size: 0.75rem;
        line-height: 0.875rem;
      }
      .selected {
        border: 3px solid #abd5ce;
      }
    }
    .coordinate.-horizontal .coordinate-item {
      width: 2.5rem;
      font-size: 0.75rem;
      line-height: 0.875rem;
      &:not(:last-child) {
        margin-right: 1.313rem;
      }
    }
    .tip-height-input-box {
      width: 14.125rem;
      height: 2.25rem;
      position: relative;
      padding-right: 3rem;
      .tip-height-input {
        padding: 0.25rem 2.8rem 0.25rem 0.75rem;
      }
      .height-icon-btn {
        position: absolute;
        top: 0.4rem;
        right: 0.75rem;
        color: #717171;
      }
    }
  }
`;

export const DeckPositionInfoBox = styled.div`
  .process-box {
    width: 90.55%;

    .title-heading {
      margin-bottom: 2.188rem;
    }
    h5 {
      font-size: 1.125rem;
      line-height: 1.313rem;
      color: #666666;
    }
    .deck-position-options {
      > button {
        border: 1px solid #92c4bc;
        border-radius: 1rem;
        width: auto;
        min-width: 6.938rem;
        height: 1.875rem;
        font-size: 0.813rem;
        line-height: 0.5rem;
        margin-bottom: 1.375rem;
        margin-right: 1.125rem;
        &.selected-opt {
          background-color: #92c4bd;
          color: #fafafa;
        }
      }
    }
    .tip-height-input-box {
      width: 14.125rem;
      height: 2.25rem;
      position: relative;
      .tip-height-input {
        padding: 0.25rem 2.8rem 0.25rem 0.75rem;
      }
      .height-icon-btn {
        position: absolute;
        top: 0.4rem;
        right: 0.75rem;
        color: #717171;
      }
    }
  }
`;
