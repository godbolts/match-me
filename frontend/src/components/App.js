import React from 'react';
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';
import { WebSocketProvider } from './WebSocketContext';
import { useState, useEffect } from 'react';


import Header from './Header/Header';
import Landing from './Landing/Landing';
import LoginSignup from './LoginSignup/LoginSignup';
import Dashboard from './Dashboard/Dashboard';
import Profile from './Profile/Profile';
import Matches from './Matches/Matches';
import BuddiesSection from './BuddiesSection/BuddiesSection.jsx';
import Chat from './Chat/Chat';
import Requests from './Requests/Requests.jsx';


function App() {

  const isAuthenticated = !!localStorage.getItem('token');
  const profileExists = localStorage.getItem('profileExists');

  return (
    <WebSocketProvider>
      <BrowserRouter>
        <div>
          <Header />
          <Routes>
            <Route exact
              path='/'
              element={isAuthenticated && !profileExists ? <Navigate to="/profile" /> : <LoginSignup />} />
            <Route
              path='/login'
              element={!isAuthenticated ? <LoginSignup /> : <Navigate to="/login" />} />
            {/* <Route
              path='/dashboard'
              element={isAuthenticated && profileExists ? <Dashboard /> : <Navigate to="/profile" />} /> */}
            <Route
              path='/profile'
              element={isAuthenticated ? <Profile /> : <Navigate to="/login" />} />
            <Route
              path='/matches'
              element={isAuthenticated && profileExists ? <Matches /> : <Navigate to="/login" />} />
            <Route
              path='/connections'
              element={isAuthenticated && profileExists ? <BuddiesSection /> : <Navigate to="/login" />} />
            <Route
              path='/chat'
              element={isAuthenticated && profileExists ? <Chat /> : <Navigate to="/login" />} />
            <Route
                path='/Requests'
                element={isAuthenticated && profileExists ? <Requests/> : <Navigate to="/login" />} />
            </Routes>
        </div>
      </BrowserRouter >
    </WebSocketProvider>
  );
}

export default App;
