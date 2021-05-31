import styled from 'styled-components';

export const PageBody = styled.div`
background-color:#f5f5f5;
`;
export const AspireDispenseBox = styled.div`
.process-aspire-dispense {
	&::after {
		background:url('/images/aspire-dispense-bg.svg') no-repeat;
	}
	.side-bar{
		width:10.875rem;
		height:100vh;
		background: rgb(131,180,172);
		background: linear-gradient(-90deg, rgba(131,180,172,1) 0%, rgba(178,218,209,1) 100%);
    box-shadow:0px 3px 16px rgba(0,0,0,0.06);
    .nav-link{
			color:#666666;
			border-radius:0;
			padding: 0.5rem 1.5rem;
			display:flex;
			justify-content:flex-start;
			align-items:center;
			height:52px;
			&.active{
				&::after{
					top:inherit;
				}
			}
		}
		.icon-upward-magnet{
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
		.icon-downward-magnet{
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
  .tab-content-top-heading{
    height:2.875rem;
    padding: 0 2.875rem;
    font-size:18px;
    color:#9D9D9D70;
  }
  label{
		font-size:1rem;
		line-height:1.125rem;
		color:#666666;
	}
	.well-box{
		// counter-reset: section;
		.well{
			&:active{
				background-color: #ffffff;
			}
		}
		.well-no{
			.coordinate-item{
			color:#999999;
			font-size:18px;
			line-height:21px;
			}
		}
		.selected{
			border:3px solid #ABD5CE;
		}
		.aspire-from{
			border:3px solid #C9C6C6;
			position:relative;
			&::after{
				content:"Aspired from";
				position:absolute;
				left: auto;
				right: auto;
				bottom: -20px;
				font-size: 12px;
				line-height: 14px;
				white-space: nowrap;
				color:#717171;
			}
		}
	}
  .tab-pane{
    padding:29px 34px !important;
		.label-name{
			width:146px;
		}
		.input-field{
			width:14.125rem;
			height:2.25rem;
			.height-icon-box{
				position:absolute;
				top:3px;
				right:12px;
			}
		}
		.cycle-input{
			width: 4rem;
			height:2.25rem;
		}
		.aspire-input-field, .dispense-input-field{
			padding-right:2rem;
		}
  }

}
`;

export const TopContent = styled.div`
	margin-bottom:0.75rem;
	.frame-icon{
		> button {
			> i{
				font-size:29px;
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
    padding-right: 2rem;
  }
`;
