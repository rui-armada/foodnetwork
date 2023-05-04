import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

const root = document.getElementById('root');

if (root) {
  ReactDOM.createRoot(root).render(<App />);
} else {
  console.error('Failed to find root element');
}
