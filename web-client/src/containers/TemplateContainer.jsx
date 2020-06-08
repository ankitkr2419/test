import React from 'react';
// import { useDispatch, useSelector } from 'react-redux';
// import { fetchMasterTargets, fetchTargetsByTemplateID } from 'actionCreators/targetActionCreators';
// import { fetchTemplates } from 'actionCreators/templateActionCreators';


const TemplateContainer = props => {
  /** 
  const listTargetReducer = useSelector(state => state.listTargetReducer)
  const listTemplateReducer = useSelector(state => state.listTemplateReducer)
  const listTargetByTemplateIDReducer = useSelector(state => state.listTargetByTemplateIDReducer)

  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchMasterTargets());
    dispatch(fetchTemplates());
    dispatch(fetchTargetsByTemplateID("1234567890"));
  }, [dispatch])
  */
  return <div>templates</div>;
};

TemplateContainer.propTypes = {};

export default TemplateContainer;