import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route, IndexRoute, browserHistory } from 'react-router';

import { Wrapper } from './Wrapper';
import { Home } from './pages/Home/Home';
import { Soundboard } from './pages/Soundboard/Soundboard';
import { NotFound } from './pages/NotFound/NotFound';
import { Downloader } from './pages/Downloader/Downloader';
import { Pubg } from './pages/Pubg/Pubg';
import { Clips } from './pages/Clips';
import { Oauth } from './pages/oauth/oauth';
import 'babel-polyfill';

ReactDOM.render(
    <Router history={browserHistory}>
        <Route path="/" component={Wrapper}>
            <IndexRoute component={Home}/>
            <Route path="/soundboard" component={Soundboard}/>
            <Route path="/downloader" component={Downloader}/>
            <Route path="/pubg" component={Pubg}/>
            <Route path="/clips" component={Clips}/>
            <Route path="/oauth" component={Oauth}/>
            <Route path="*" component={NotFound}/>
        </Route>
    </Router>
, document.getElementById('app'));
