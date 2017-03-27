import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route, IndexRoute, browserHistory } from 'react-router';

import Wrapper from './Wrapper';
import Soundboard from './pages/Soundboard/Soundboard';
import NotFound from './pages/NotFound/NotFound';

ReactDOM.render(
    <Router history={browserHistory}>
        <Route path="/" component={Wrapper}>
            <IndexRoute component={Soundboard}/>
            <Route path="*" component={NotFound}/>
        </Route>
    </Router>
, document.getElementById('app'));
