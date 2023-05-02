// src/App.tsx
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from '@/components/Login';
import Home from '@/components/Home';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" Component={Home} />
        <Route path="/login" Component={Login} />
        <Route path="/home" Component={Home} />
      </Routes>
    </Router>
  );
}

export default App;
