import styled from "styled-components";

export const PageBody = styled.div`
  background-color: #f5f5f5;
`;
export const LabwareBox = styled.div`
  .process-labware {
    padding: 1rem 4.5rem 0.75rem 4.5rem;
    .side-bar {
      width: 10.875rem;
      height: 100vh;
      background: rgb(131, 180, 172);
      background: linear-gradient(
        -90deg,
        rgba(131, 180, 172, 1) 0%,
        rgba(178, 218, 209, 1) 100%
      );
      box-shadow: 0px 3px 1rem rgba(0, 0, 0, 0.06);
      .nav-link {
        color: #000000;
        border-radius: 0;
        padding: 0.5rem 1.25rem;
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

    label {
      font-size: 1rem;
      line-height: 1.125rem;
      color: #666666;
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
          top: 3px;
          right: 0.75rem;
        }
      }
    }
    .bottom-btn-bar {
      > div {
        bottom: 2rem;
      }
    }
    .labware-card-box {
      height: 35.813rem;
      .icon-tick {
        transform: rotate(20deg);
      }
    }

    //Preview
    .preview-box {
      .top-heading {
        height: 3rem;
        padding: 0 6.25rem;
        background-color: #b2dad122;
        font-size: 1.125rem;
        line-height: 1.313rem;
      }
      .labware-selection-info {
        padding: 2rem 5.5rem 2rem 5.5rem;

        .selected-positions {
          color: #92c4bc;
        }
        .setting-value {
          color: #666666;
        }
      }
    }
  }
`;

export const ProcessSetting = styled.div`
  width: 22.938rem;
  height: 30.813rem;
  position: absolute;
  top: -15px;
  right: -2px;
  img {
    max-width: 100%;
  }
  .highlighted {
    background: rgba(243, 130, 32, 0.3);
    width: 80%;
    height: 1.125rem;
    display: block;
    position: absolute;
    left: 18%;
    right: 0;
    margin: 0 auto;
    border-radius: 0.25rem;
  }
  .tips-info .active {
    opacity: 1;
  }

  // .tips-info .active .active {
  //   opacity: 0.5;
  // }
  .tip-position {
    .tip-position-1 {
      top: 2.5rem;
    }

    .tip-position-2 {
      top: 3.625rem;
    }

    .tip-position-3 {
      top: 4.688rem;
    }
  }
  // Piercing
  .piercing-info {
    .piercing-position {
      .piercing-position-1 {
        top: 5.625rem;
      }
      .piercing-position-2 {
        top: 6.688rem;
      }
    }
  }

  // Deck Position 1
  .deck-position-info {
    .deck-position {
      .deck-position-1 {
        top: 8rem;
        height: 1.25rem;
      }
      .deck-position-2 {
        top: 9.375rem;
        height: 4.688rem;
      }
      .deck-position-3 {
        top: 21.25rem;
        height: 1.25rem;
      }
      .deck-position-4 {
        top: 25.625rem;
        height: 2rem;
      }
    }
  }

  //// Deck Position 1
  .cartridge-position-info {
    .cartridge-position {
      .cartridge-position-1 {
        top: 14.063rem;
        height: 6.875rem;
      }
      .cartridge-position-2 {
        top: 22.5rem;
        height: 2.5rem;
      }
    }
  }
`;
