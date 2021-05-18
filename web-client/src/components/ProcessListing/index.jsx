import React from "react";
import { StyledProcessListing } from "./Styles";
import AppFooter from "components/AppFooter";
import { useSelector } from "react-redux";
import { Redirect } from "react-router";
import { ROUTES } from "appConstants";

const ProcessListComponent = (props) => {
    let { processList } = props;

    /**
     * get active login deck data
     */
    const loginReducer = useSelector((state) => state.loginReducer);
    const loginReducerData = loginReducer.toJS();
    let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);

    /**
     * if user is not logged in, go to landing page
     */
    if (!activeDeckObj.isLoggedIn) {
        return <Redirect to={`/${ROUTES.landing}`} />;
    }

    return (
        <StyledProcessListing>
            <div className="landing-content px-2">
                {/* TODO process listing page UI*/}

            </div>
            <AppFooter />
        </StyledProcessListing>
    );
};

export default React.memo(ProcessListComponent);
