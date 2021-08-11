import React, { useState, useEffect } from "react";
import SplashScreenComponent from "components/SplashScreenPage";
import { ROUTES, APP_TYPE } from "appConstants";
import { useSelector, useDispatch } from "react-redux";
import { appInfoInitiated } from "action-creators/appInfoActionCreators";

const SplashScreenContainer = () => {
    const dispatch = useDispatch();
    const appInfoReducer = useSelector((state) => state.appInfoReducer);
    const appInfoData = appInfoReducer.toJS();
    const app = appInfoData?.appInfo?.app;

    useEffect(() => {
        dispatch(appInfoInitiated());
    }, []);

    //appropriate page route depend on app type
    const findPath = () => {
        switch (app) {
            case APP_TYPE.EXTRACTION:
                return ROUTES.landing;
            case APP_TYPE.RTPCR:
                return ROUTES.login;
            default:
                return null;
        }
    };

    return <SplashScreenComponent redirectionPath={findPath()} />;
};

SplashScreenContainer.propTypes = {};

export default SplashScreenContainer;
