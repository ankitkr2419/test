import styled from "styled-components";

export const PageBody = styled.div`
background-color:#f5f5f5;
`;

export const ProcessInnerBox = styled.div`
margin-bottom:14px;
// background: black;
.process-card{
    border-radius:9px;
    border:1px solid #CCCCCC;
    padding:29px 34px;
    height:108px;
    .process-name{
        font-size:18px;
        line-height:21px;
        color: #666666;
    }
}
`;

export const ProcessOuterBox = styled.div`
.select-process-bg {
	// background:black;
	padding: 16px 94px;
	&::after {
		background:url('/images/process-bg.svg')no-repeat;
		background-position: top right;
		margin-top:20px;
	}
	.process-content-box{
		background:none;
		border:none;
		box-shadow:none;
	}
	.process-card{
		border-radius:9px;
		border:1px solid #CCCCCC;
		padding:29px 34px;
		height:108px;
		.process-name{
			font-size:18px;
			line-height:21px;
			color: #666666;
		}
	}
	// .frame-icon{
	// 	> button {
	// 		> i{
	// 			font-size:28px;
	// 		}
	// 	}
	// }
	.row-small-gutter {
			margin-left: -7px !important;
			margin-right: -7px !important;
	}

	.row-small-gutter > * {
			padding-left: 7px !important;
			padding-right: 7px !important;
	}
	.icon-piercing{
		font-size:28px;
	}
	.icon-tip-pickup{
		font-size:21px;
	}
	.icon-aspire-dispense{
		font-size:17px;
	}
	.icon-shaking{
		font-size:26px;
	}
	.icon-heating{
		font-size:22px;
	}
	.icon-magnet{
		font-size:17px;
	}
	.icon-tip-discard{
		font-size:24px;
	}
	.icon-delay{
		font-size:19px;
	}
	.icon-tip-position{
		font-size:24px;
	}
`;

export const TopContent = styled.div`
	margin-bottom:0.75rem;
}
`;

export const HeadingTitle = styled.label`
    font-size: 1.25rem;
    line-height: 1.438rem;
`;