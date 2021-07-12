import React, { useState } from "react";
import { Redirect } from "react-router";

import LoginComponent from "components/Login";
import { useDispatch, useSelector } from "react-redux";
import {
    loginAsOperator as loginAsOperatorAction,
    login,
} from "action-creators/loginActionCreators";
import { toast } from "react-toastify";

const LoginContainer = () => {
    const dispatch = useDispatch();

    const loginReducer = useSelector((state) => state.loginReducer);
    const loginReducerData = loginReducer.toJS();
    let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
    const { isLoggedIn, error, name } = activeDeckObj; //name refers to deckName

    // local state to handle admin form visibility
    const [isAdminFormVisible, setIsAdminFormVisibility] = useState(false);

    // local state to handle operator login modal visibility
    const [operatorLoginModalOpen, setOperatorLoginModalOpen] = useState(false);

    const loginAsAdmin = (data) => {
        let { username, password } = data;
        let role = username
            ? username.split("@").pop().split(".")[0]
            : undefined;
        let email = username.split("@")[0];

        if (role !== "admin") {
            toast.error("Incorrect admin credentials");
            return;
        }

        dispatch(login({ email, password, deckName: name, role }));
    };

    const toggleOperatorLoginModal = () => {
        setOperatorLoginModalOpen(!operatorLoginModalOpen);
    };

    // redirection to admin once logged in
    if (isLoggedIn === true) {
        // if (isLoggedIn === true && isSocketConnected === true) {
        return <Redirect to="/templates" />;
    }

    return (
        <LoginComponent
            isAdminFormVisible={isAdminFormVisible}
            setIsAdminFormVisibility={setIsAdminFormVisibility}
            operatorLoginModalOpen={operatorLoginModalOpen}
            toggleOperatorLoginModal={toggleOperatorLoginModal}
            deckName={name}
            loginAsAdmin={loginAsAdmin}
            isLoginError={error}
        />
    );
};

export default LoginContainer;
