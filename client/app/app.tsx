import { Provider } from 'mobx-react';
import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import {
  Clips,
  Downloader,
  NotFound,
  Oauth,
  Soundboard,
  Stats,
  UploadHistory,
  UserEventLog,
  VideoArchive,
} from './pages';
import { Users } from './pages/users';
import './scss/index.scss';
import { rootStoreInstance } from './stores';
import { Wrapper } from './wrapper';

const App = () => {
  return (
    <BrowserRouter>
      <Provider {...rootStoreInstance}>
        <Wrapper>
          <Switch>
            <Route exact path="/" component={Soundboard} />
            <Route path="/upload-history" component={UploadHistory} />
            <Route path="/downloader" component={Downloader} />
            <Route path="/clips" component={Clips} />
            <Route path="/oauth" component={Oauth} />
            <Route path="/stats" component={Stats} />
            <Route path="/video-archive" component={VideoArchive} />
            <Route path="/user-event-log" component={UserEventLog} />
            <Route path="/users" component={Users} />
            <Route component={NotFound} />
          </Switch>
        </Wrapper>
      </Provider>
    </BrowserRouter>
  );
};

ReactDOM.render(<App />, document.getElementById('app'));
