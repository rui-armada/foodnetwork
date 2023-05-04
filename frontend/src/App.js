import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './Login';
import Home from './Home';

function App() {
  const [loggedIn, setLoggedIn] = useState(false);

  const checkLoggedInStatus = () => {
    const token = localStorage.getItem('token');
    if (token) {
      setLoggedIn(true);
    } else {
      setLoggedIn(false);
    }
  };

  useEffect(() => {
    checkLoggedInStatus();
  }, []);

  return (
    <Router>
      <Routes>
        <Route path="/" element={loggedIn ? <Home /> : <Login onLoginSuccess={checkLoggedInStatus} />} />
        <Route path="/login" element={<Login onLoginSuccess={checkLoggedInStatus} />} />
        <Route path="/home" element={loggedIn ? <Home /> : <Login onLoginSuccess={checkLoggedInStatus} />} />
      </Routes>
    </Router>
  );
}

export default App;
