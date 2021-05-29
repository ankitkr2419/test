import styled from "styled-components";

export const ActionBtn = styled.button`
    position: absolute;
    top: 0;
    right: 0;
    z-index: 2;
    display: block;
    outline: none;
    border: 0;
    background: transparent;
    .semi-circle-outter-box {
        width: 3.875rem;
        height: 2.5rem;
        background-color: #b3d9d0;
        border-bottom-left-radius: 5.5rem;
        border-bottom-right-radius: 5.5rem;
        box-shadow: 0px -3px 6px rgb(0 0 0 / 31%);
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .semi-circular-button {
        width: 2.75rem;
        height: 2.75rem;
        border-radius: 50%;
        border: 1px solid #f0801d;
        background-color: #f38220;
        z-index: 1;
        text-decoration: none;
        margin-top: -1.375rem;
        display: flex;
        justify-content: center;
        align-items: center;
        color: #fff;
        position: relative;
        .btn-label {
            position: absolute;
            bottom: -2rem;
            left: 0;
            right: 0;
            text-align: center;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 0.75rem;
            line-height: 0.875rem;
            color: #3c3c3c;
        }
        .icon-play,
        .icon-resume {
            margin-left: 0.25rem;
        }
    }
`;

// export const DeckCardBox = styled.div`
//   width: 50%;
//   // width: 32rem;
//   height: 6.625rem;
//   position: relative;
//   box-shadow: 0px -3px 6px rgba(0, 0, 0, 0.16);
//   &::before {
//     content: "";
//     position: absolute;
//     background-image: linear-gradient(
//       to right,
//       #aedbd5,
//       #a9dac5,
//       #afd7b0,
//       #bed29a,
//       #d3ca87,
//       #dcc278,
//       #e7b96c,
//       #f2ae64,
//       #f2a453,
//       #f29942,
//       #f38d31,
//       #f3811f
//     );
//     width: 100%;
//     height: 2px;
//     top: 0;
//     left: 0;
//     z-index: 1;
//   }
//   .deck-title {
//     width: 2.563rem;
//     height: 100%;
//     font-size: 1.25rem;
//     line-height: 1.688rem;
//     font-weight: bold;
//     color: #51575a;
//     // background-color: #b2dad1;
//     // border: 1px solid #ffffff;
//     border: 1px solid transparent;
//     box-shadow: 0 -3px 6px rgba(0, 0, 0, 0.16);
//     > label {
//       transform: rotate(-90deg);
//       white-space: nowrap;
//       margin-bottom: 0;
//     }
//     &.active {
//       background-color: #b2dad1;
//       border: 1px solid #ffffff;
//     }
//   }
//   .deck-content {
//     position: relative;
//     // background: #fff url("/images/deck-card-bg.svg") no-repeat;
//     > button {
//       min-width: 7.063rem;
//       height: 2.5rem;
//       line-height: 1.125rem;
//     }
//     .custom-progress-bar {
//       border-radius: 7px;
//       background-color: #b2dad131;
//       border: 2px solid #b2dad131;
//       .progress-bar {
//         //background-color:#10907A;
//         border-radius: 7px 0px 0px 7px;
//         background-color: #72b5e6;
//         animation: blink 1s linear infinite;
//       }
//     }
//     // .uv-light-button{
//     // 	position:absolute;
//     // 	right:244px;
//     // 	top:0;
//     // }
//     .resume-button {
//       position: absolute;
//       right: 123px;
//       top: 0;
//       .icon-pause {
//         font-size: 0.938rem;
//       }
//       .icon-resume {
//         font-size: 1.25rem;
//       }
//     }
//     .abort-button {
//       position: absolute;
//       right: 21px;
//       top: 0;
//       .semi-circular-button {
//         border: 1px solid transparent;
//         background-color: #ffffff;
//         color: #3c3c3c;
//       }
//       .icon-cancel {
//         font-size: 0.875rem;
//       }
//     }
//     .hour-label {
//       background-color: #f5e3d3;
//       border-radius: 4px 0 0 4px;
//       border-right: 2px solid #f38220;
//       padding: 3px 4px;
//       font-size: 0.875rem;
//       line-height: 1rem;
//     }
//     .min-label {
//       font-size: 0.875rem;
//       line-height: 1rem;
//     }
//     .process-count-label {
//       background-color: #f5e3d3;
//       border-radius: 4px;
//       padding: 3px 4px;
//       font-size: 1.125rem;
//       line-height: 1rem;
//     }
//     .process-total-count {
//       font-size: 0.875rem;
//       line-height: 1rem;
//     }
//     .process-remaining {
//       font-size: 10px;
//       line-height: 11px;
//     }
//     // add this class while login
//     &.logged-in {
//       background: #ffffff;
//     }
//   }
//   @keyframes blink {
//     0% {
//       background-color: #9d9d9d;
//     }
//     50% {
//       background-color: #72b5e6;
//     }
//     100% {
//       background-color: #9d9d9d;
//     }
//   }
//   .marquee {
//     width: 80%;
//     white-space: nowrap;
//     overflow: hidden;
//     box-sizing: border-box;
//     .recipe-name {
//       display: inline-block;
//       padding-left: 100%;
//       animation: marquee 10s linear infinite;
//     }
//     @keyframes marquee {
//       0% {
//         transform: translate(0, 0);
//       }
//       100% {
//         transform: translate(-100%, 0);
//       }
//     }
//   }
// `;


export const DeckCardBox = styled.div`
  // width: 50%;
  //   width: 32rem;
  flex: 1;
  height: 6.625rem;
  position: relative;
  box-shadow: 0px -3px 6px rgba(0, 0, 0, 0.16);
  &::before {
    content: "";
    position: absolute;
    background-image: linear-gradient(
      to right,
      #aedbd5,
      #a9dac5,
      #afd7b0,
      #bed29a,
      #d3ca87,
      #dcc278,
      #e7b96c,
      #f2ae64,
      #f2a453,
      #f29942,
      #f38d31,
      #f3811f
    );
    width: 100%;
    height: 2px;
    top: 0;
    left: 0;
    z-index: 1;
  }
  .deck-title {
    width: 2.563rem;
    height: 100%;
    font-size: 1.25rem;
    line-height: 1.688rem;
    font-weight: bold;
    color: #51575a;
    border: 1px solid transparent;
    box-shadow: 0 -3px 6px rgba(0, 0, 0, 0.16);
    > label {
      transform: rotate(-90deg);
      white-space: nowrap;
      margin-bottom: 0;
    }
    &.active {
      background-color: #b2dad1;
      border: 1px solid #ffffff;
    }
  }
  .deck-content {
    position: relative;
    background: #fff url("/images/deck-card-bg.svg") no-repeat;
    > button {
      min-width: 7.063rem;
      height: 2.5rem;
      line-height: 1.125rem;
    }
    .custom-progress-bar {
      border-radius: 7px;
      background-color: #b2dad131;
      border: 2px solid #b2dad131;
      .progress-bar {
        //background-color:#10907A;
        border-radius: 7px 0px 0px 7px;
        background-color: #72b5e6;
        animation: blink 1s linear infinite;
      }
    }
    // .uv-light-button{
    // 	position:absolute;
    // 	right:244px;
    // 	top:0;
    // }
    .resume-button {
      position: absolute;
      right: 123px;
      top: 0;
    }
    .abort-button {
      position: absolute;
      right: 21px;
      top: 0;
    }
    .hour-label {
      background-color: #f5e3d3;
      border-radius: 4px 0 0 4px;
      border-right: 2px solid #f38220;
      padding: 3px 4px;
      font-size: 0.875rem;
      line-height: 1rem;
    }
    .min-label {
      font-size: 0.875rem;
      line-height: 1rem;
    }
    // add this class while login
    &.logged-in {
      background: #ffffff;
    }
    .process-remain-label {
      background-color: #f5e3d3;
      border-radius: 4px;
      padding: 3px 4px;
      font-size: 0.875rem;
      line-height: 1rem;
    }
  }
  @keyframes blink {
    0% {
      background-color: #9d9d9d;
    }
    50% {
      background-color: #72b5e6;
    }
    100% {
      background-color: #9d9d9d;
    }
  }
`;

export const CardOverlay = styled.div`
  position: absolute;
  // display: none;
  width: 100%;
  height: 6.625rem;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.28);
  z-index: 3;
  cursor: pointer;
`;
