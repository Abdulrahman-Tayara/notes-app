import React from 'react';
import './App.css';
import SignUpPage from 'features/auth/signup/SignUpPage';
import { setup as setupDI } from 'di/containers';

setupDI()

function App() {
  return (
    <SignUpPage />
  );
}

export default App;
