import React, { useEffect, useState, useContext } from 'react';
import axios from 'axios';
import { Link, } from 'react-router-dom';
import './Header.css'
import { WebSocketContext } from '../WebSocketContext';

function Header() {
    const socket = useContext(WebSocketContext)
    const [isOnline, setIsOnline] = useState(null);
    const [isLoggingOut, setIsLoggingOut] = useState(false);
    const [displayName, setDisplayName] = useState("")
    const isAuthenticated = !!localStorage.getItem('token');
    const authToken = localStorage.getItem('token');


    useEffect(() => {
        const fetchOnlineStatus = async () => {
            if (isAuthenticated) {
                try {
                    const response = await axios.get('/online', {
                        headers: {
                            Authorization: `Bearer ${authToken}`,
                        },
                    });
                    const data = response.data;
                    setIsOnline(data.is_online); // Assuming the backend returns true/false
                } catch (error) {
                    console.error('Error fetching online status:', error);
                }
            }
        };
        fetchOnlineStatus();
    }, [isAuthenticated, authToken]);

    useEffect(() => {
        const fetchUsername = async () => {
            if (isAuthenticated) {
                try {
                    const uuidResponse = await axios.get('/me/uuid', {
                        headers: { Authorization: `Bearer ${authToken}` },
                    });
                    console.log('Fetched currentUserID:', uuidResponse.data);

                    if (uuidResponse.data) {
                        const usernameResponse = await axios.get(`/users/${uuidResponse.data}/profile`, {
                            headers: { Authorization: `Bearer ${authToken}` },
                        });
                        if (usernameResponse.data.username) {
                            setDisplayName(usernameResponse.data.username)
                        } else {
                            console.error('Error fetching username')
                        }
                    }
                } catch (error) {
                    console.error('Error fetching online status:', error);
                }
            };
        }
        fetchUsername();
    }, [isAuthenticated, authToken]);
    const handleLogout = async () => {
        if (isLoggingOut) return;
        setIsLoggingOut(true);
        try {
            const response = await axios.get('/logout', {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log('Logout response:', response.data);
            if (response.data) {
                // Fetch UUID and username
                const uuidResponse = await axios.get('/me/uuid', {
                    headers: { Authorization: `Bearer ${authToken}` },
                });
                console.log('Fetched currentUserID:', uuidResponse.data);
                if (uuidResponse.data) {
                    const usernameResponse = await axios.get(`/users/${uuidResponse.data}/profile`, {
                        headers: { Authorization: `Bearer ${authToken}` },
                    });
                    console.log('Fetched currentUsername:', usernameResponse.data.username);
                    if (socket && socket.readyState === WebSocket.OPEN) {
                        socket.send(JSON.stringify({ type: 'logout', username: usernameResponse.data.username }));
                        console.log('Sent logout message:', usernameResponse.data.username);
                    } else {
                        console.error('Socket is not open or username is missing.');
                    }
                }
                socket.close();
                localStorage.removeItem('token');
                localStorage.removeItem('profileExists');
                window.location.href = '/login';
            } else {
                console.error('Logout failed on the backend.');
                localStorage.removeItem('token');
                localStorage.removeItem('profileExists');
            }
        } catch (error) {
            console.error('Error during logout:', error);
            localStorage.removeItem('token');
            localStorage.removeItem('profileExists');
        } finally {
            setIsLoggingOut(false);
        }
    };
    return (
        <>

            <header className="header">
                <div className="nav-left">
                    {isAuthenticated ? (
                        <>
                            <Link to="/profile" className="nav-link">
                                Dashboard
                            </Link>
                            {/* <Link to="/profile" className="nav-link">
                                    Profile
                                </Link> */}
                            <Link to="/matches" className="nav-link">
                                Recommendations
                            </Link>
                            <Link to="/Requests" className="nav-link">
                                Requests
                            </Link>
                            <Link to="/connections" className="nav-link">
                                Matches
                            </Link>
                            <Link to="/chat" className="nav-link">
                                Chat
                            </Link>
                        </>
                    ) : (
                        <Link to='/login' className="logo">
                        </Link>

                    )}
                </div>
                <div className='nav-container'></div>
                <div className="nav-right">
                    {isAuthenticated && (
                        <div className="online-status">
                            <span className="nav-link">{isOnline === true ? `${displayName} ` : isOnline === false ? `${displayName} ` : 'Loading...'}</span>
                            <span
                                className={`status-light ${isOnline === true
                                    ? 'online' // Green if online
                                    : isOnline === false
                                        ? 'offline' // Red if offline
                                        : '' // No color if status is unknown
                                    }`}
                            ></span>
                        </div>
                    )}
                    {!isAuthenticated ? (
                        <Link to="/login" className="signup">
                            Sign up/Login
                        </Link>
                    ) : (
                        <Link
                            to="/login"
                            className="signup"
                            onClick={(e) => {
                                e.preventDefault();
                                handleLogout();
                            }}
                            disabled={isLoggingOut}
                        >
                            {isLoggingOut ? 'Logging out...' : 'Logout'}
                        </Link>
                    )}
                </div>
            </header>

        </>
    );
}

export default Header;