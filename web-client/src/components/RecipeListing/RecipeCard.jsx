import React from 'react';
import { Text } from 'shared-components';
import styled from 'styled-components';

const RecipeCardStyle = styled.div`
    padding: 0.8rem 0.5rem;
    border: 1px solid #E3E3E3;
    border-radius: 0.5rem;
    margin-bottom:0.688rem;
    box-shadow: 0px 3px 16px rgba(0,0,0,0.04);
    // width:27.5rem;
    // height: 5.563rem;
    .recipe-heading{
        padding-bottom:0.5rem;
    }
    .recipe-card-body{
        padding-top:0.25rem;
        border-top: 1px solid #d9d9d9;
        
        .recipe-name{
            font-size:0.875rem;
            line-height:1rem;
        }
        .recipe-value{
            font-size:1.125rem;
            line-height:1.313rem;
        }
        .recipe-action{
            button {
                width:33px !important;
                height:33px !important;
                border:1px solid #696969 !important;
                &:not(:first-child){
                    margin-left:12px;
                }
            }
        }
    }
    &:focus, &:hover{
        background-color:rgba(243, 130, 32, 0.30);
    }
`;

const RecipeCard = (props) => {
    // const [fadeIn, setFadeIn] = useState(true);
    // const toggle = () => setFadeIn(!fadeIn);

    const {
        recipeId,
        recipeName,
        processCount,
        toggle
    } = props;

    return(
        <div onClick={() => toggle(recipeId, recipeName, processCount)}>
            <RecipeCardStyle>
                <div className="font-weight-bold recipe-heading">{recipeName}</div>
                <div className="recipe-card-body">
                <Text Tag="span" className="recipe-name">Total Processes -</Text>
                <Text Tag="span" className="text-primary font-weight-bold recipe-value ml-2">{processCount} </Text>
                {/* <Fade in={fadeIn} tag="h5" className="m-0 d-none">
                        <div className="recipe-action d-flex justify-content-between align-items-center">
                            <div className="d-flex justify-content-between align-items-center">
                                <ButtonIcon
                                    size={14}
                                    name='play'
                                    className="border-gray text-primary"
                                    //onClick={toggleExportDataModal}
                                />
                                <ButtonIcon
                                    size={14}
                                    name='edit-pencil'
                                    className="border-gray text-primary"
                                    //onClick={toggleExportDataModal}
                                />
                                <ButtonIcon
                                    size={14}
                                    name='upload'
                                    className="border-gray text-primary"
                                    //onClick={toggleExportDataModal}
                                />
                            </div>
                            <ButtonIcon
                                size={20}
                                name='minus-1'
                                className="border-gray text-primary"
                                //onClick={toggleExportDataModal}
                            />
                        </div>
                    </Fade> */}
                </div> 
            </RecipeCardStyle>
        </div>
    )
}

export default RecipeCard;
