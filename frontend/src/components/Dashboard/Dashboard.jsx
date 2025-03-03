import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Dashboard.css';
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';
import { Link } from 'react-router-dom';
import Matches from '../Matches/Matches';

/* 
This page is not used, as it did not provide any additional functionality to the project.
*/
const Dashboard = () => {
    const [userData, setUserData] = useState(null); // Store user data
    const [loading, setLoading] = useState(true); // Track loading state for user data
    const [error, setError] = useState(null); // Track errors
    const [profilePic, setProfilePic] = useState(null);

    useEffect(() => {
        const fetchUserData = async () => {
            try {

                const token = localStorage.getItem('token');
                if (!token) {
                    setError('No token found'); // Handle missing token
                    setLoading(false);
                    return;
                }

                const response = await axios.get('/me/profile', {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                const data = response.data; // Store the user data

                setUserData(data);

                // Handle profile picture (default or from backend)
                if (data.profile_picture) {
                    setProfilePic(`/uploads/${data.profile_picture}`); // Assuming it's a URL
                } else {
                    setProfilePic(defaultProfilePic); // Use default profile picture
                }

            } catch (err) {
                setError(err.response ? err.response.data : 'An error occurred');
            } finally {
                setLoading(false); // Stop loading for user data
            }
        };

        fetchUserData();
    }, []);

    const handleEditProfile = () => {
        // Redirect to the edit profile page
        window.location.href = '/profile';
    }

    if (loading) {
        // Show loading for user data
        return <div>Loading user data...</div>;
    }

    if (error) {
        // Show an error message if the request fails
        return <div>Error: {error}</div>;
    }

    return (
        <>
            <div className="dashboard-container">
                <div className="dashboard-card">
                    <div className="dashboard-profile-pic">
                        {profilePic ? (
                            <img src={profilePic} alt="Profile" />
                        ) : (
                            <img src={defaultProfilePic} alt="Default Profile" />
                        )}
                    </div>
                    <div className="dashboard-text-data">
                        <h2>{userData?.username}</h2>
                        <p>{userData?.email}</p>
                        <p>{userData?.age}</p>
                        <p>{`${userData?.user_nation}, ${userData?.user_region}, ${userData?.user_city}`}</p>
                        <p>{userData?.about_me}</p>
                        <button onClick={handleEditProfile} >Edit profile</button>
                    </div>
                </div>
            </div>
        </>
    );
};

export default Dashboard;
