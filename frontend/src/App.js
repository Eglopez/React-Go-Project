import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';

import Login from './views/LoginView';
import Signup from './views/SignupView';

function App() {
  return (
    <>
      <Router>
        <Route path="/" exact component={Login}/>
        <Route path="/login" exact component={Login}/>
        <Route path="/signup" component={Signup}/>
      </Router>
    </>
  );
}

export default App;
