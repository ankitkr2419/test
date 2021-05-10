import styled from "styled-components";

//For Run Recipes Type Form
export const RunRecipeTypeForm = styled.div`
    .step-run-option,
    .continuous-run-option {
        border-radius: 1rem;
        width: 22.188rem;
        height: 3.75rem;
        padding: 0 5rem;
    }
    .step-run-option {
        border: 2px dashed #dbdbdb;
        margin-bottom: 30px;
    }
    .continuous-run-option {
        border: 2px solid #dbdbdb;
        margin-bottom: 30px;
    }
    .selected {
        border-color: #b2dad1;
        box-shadow: 0px 3px 16px rgba(0, 0, 0, 0.06);
    }
    label {
        font-size: 0.813rem;
        line-height: 0.938rem;
    }
`;
