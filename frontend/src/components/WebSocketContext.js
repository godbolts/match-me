import React, { createContext, useState, useEffect } from 'react';
import axios from 'axios';

export const WebSocketContext = createContext();

export const WebSocketProvider = ({ children }) => {
    const [socket, setSocket] = useState(null);
    const [senderID, setSenderID] = useState('');
    const [senderUsername, setSenderUsername] = useState('');
    const authToken = localStorage.getItem('token');

    // Fetch UUID for senderID
    useEffect(() => {
        const fetchUUID = async () => {
            try {
                const response = await axios.get('/me/uuid', {
                    headers: { Authorization: `Bearer ${authToken}` },
                });
                setSenderID(response.data);
                console.log('Profile/me/uuid:', response.data);
            } catch (error) {
                console.error('Error getting sender UUID:', error);
            }
        };

        if (authToken) {
            fetchUUID();
        }
    }, [authToken]);

    // Fetch username for sender when senderID is set
    useEffect(() => {
        const fetchUsername = async () => {
            if (!senderID) return; // Wait until senderID is available
            try {
                const response = await axios.get(`/users/${senderID}/profile`, {
                    headers: { Authorization: `Bearer ${authToken}` },
                });
                if (senderUsername === "") {
                    setSenderUsername(response.data.username);
                }
                console.log('Logged in username:', response.data.username);
            } catch (error) {
                console.error('Error getting sender username:', error);
            }
        };

        if (senderID && !senderUsername) {
            fetchUsername();
        }
    }, [senderID, senderUsername, authToken]);

    // Open WebSocket connection when senderID and senderUsername are both set
    useEffect(() => {
        if (!senderID || !senderUsername) return; // Ensure both are set before connecting

        console.log('ws made for senderID:', senderID);
        const ws = new WebSocket(`ws://localhost:4000/ws?userID=${senderID}`);

        ws.onopen = () => {
            console.log('WebSocket connected');
            ws.send(JSON.stringify({ type: 'login', username: senderUsername }));
        };

        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
        // Send logout message when the socket closes
        ws.onclose = () => {
            console.log('WebSocket disconnected');
            setSenderUsername('');
        };

        setSocket(ws);

        return () => {
            ws.close();
        };
    }, [senderID, senderUsername]); // Only establish WebSocket when both are set

    return <WebSocketContext.Provider value={socket}>{children}</WebSocketContext.Provider>;
};
