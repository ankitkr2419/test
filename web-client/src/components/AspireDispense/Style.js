import styled from "styled-components";

export const PageBody = styled.div`
  background-color: #f5f5f5;
`;
export const AspireDispenseBox = styled.div`
  .process-aspire-dispense {
    &::after {
      background: url("/images/aspire-dispense-bg.svg") no-repeat;
    }
    .side-bar {
      width: 10.875rem;
      flex: 0 0 10.875rem;
      height: 100vh;
      background: rgb(131, 180, 172);
      background: linear-gradient(
        -90deg,
        rgba(131, 180, 172, 1) 0%,
        rgba(178, 218, 209, 1) 100%
      );
      box-shadow: 0px 3px 16px rgba(0, 0, 0, 0.06);
      .nav-link {
        color: #666666;
        border-radius: 0;
        padding: 0.5rem 1.5rem;
        display: flex;
        justify-content: flex-start;
        align-items: center;
        height: 3.25rem;
        &.active {
          &::after {
            top: inherit;
          }
        }
      }
      .icon-upward-magnet {
        animation: 1s slideInFromTop;
        @keyframes slideInFromTop {
          0% {
            transform: translateY(100%);
          }
          100% {
            transform: translateY(0);
          }
        }
      }
      .icon-downward-magnet {
        animation: 1s slideInFromDown;
        @keyframes slideInFromDown {
          0% {
            transform: translateY(-100%);
          }
          100% {
            transform: translateY(0);
          }
        }
      }
    }
    .tab-content-top-heading {
      height: 2.875rem;
      padding: 0 2.875rem;
      font-size: 1.125rem;
      color: #9d9d9d70;
    }
    label {
      font-size: 1rem;
      line-height: 1.125rem;
      color: #666666;
    }
    .well-box {
      // counter-reset: section;
      .well {
        width: 40px;
        height: 40px;
        margin: 0 20px 0px 0;
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
      .selected {
        border: 3px solid #abd5ce;
      }
      .aspire-from {
        border: 3px solid #c9c6c6;
        position: relative;
        &::after {
          content: "Aspired from";
          position: absolute;
          left: auto;
          right: auto;
          bottom: -20px;
          font-size: 0.75rem;
          line-height: 0.875rem;
          white-space: nowrap;
          color: #717171;
        }
      }
    }
    .tab-pane {
      padding: 1.813rem 2.125rem !important;
      .label-name {
        width: 9.125rem;
      }
      .input-field {
        width: 14.125rem;
        height: 2.25rem;
        .height-icon-box {
          position: absolute;
          top: 5px;
          left: 11.596rem;
          color: gray;
        }
      }
      .cycle-input {
        width: 4rem;
        height: 2.25rem;
      }
      .aspire-input-field,
      .dispense-input-field {
        padding-right: 3rem;
      }
    }
    .coordinate.-horizontal .coordinate-item {
      width: 40px;
      font-size: 12px;
      line-height: 14px;
      &:not(:last-child) {
        margin-right: 20px;
      }
    }
  }
`;

export const TopContent = styled.div`
  margin-bottom: 0.5rem;
  .frame-icon {
    > button {
      > i {
        font-size: 1.813rem;
      }
    }
  }
`;

export const CommmonFields = styled.div`
  .label-name {
    width: 9.125rem;
  }
  .input-field {
    width: 14.125rem;
    height: 2.25rem;
    .height-icon-box {
      position: absolute;
      top: 3px;
      right: 0.75rem;
    }
  }
  .cycle-input {
    width: 4rem;
    height: 2.25rem;
  }
  .aspire-input-field,
  .dispense-input-field {
    padding-right: 3rem;
  }
`;
