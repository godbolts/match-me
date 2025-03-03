import React from "react";
import { useState, useEffect } from "react";
import "./MatchCard.css";
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';
import axios from 'axios';
import Modal from '../Modal/Modal.jsx';
import BuddyCard from "../BuddyCard/BuddyCard.jsx";


const authToken = localStorage.getItem('token');
const onlineURL = "/images/OnlineIconPNG.png"
const offlineURL = "/images/OfflineIconPNG.png"

const MatchCard = ({ userProfile, onUpdate }) => {

    console.log("User profile from requests: ", userProfile);

    const { match_id,
        match_score,
        status,
        is_online,
        requester,
        matched_user_id,
        matched_user_name,
        matched_user_picture,
        matched_user_description,
        matched_user_location } = userProfile;


    console.log("buddy profile: " + userProfile.matched_user_id)
    console.log("buddy profile: " + userProfile.requester)

    const basePictureURL = "http://localhost:4000/uploads/";
    const onlineURL = "/images/OnlineIconPNG.png"
    const offlineURL = "/images/OfflineIconPNG.png"
    // Set the default profile picture if no picture is provided
    let userProfilePic = matched_user_picture ? matched_user_picture : defaultProfilePic;

    if (userProfilePic !== defaultProfilePic) {
        userProfilePic = basePictureURL + userProfilePic;
    }

    const [isModalOpen, setModalOpen] = useState(false);
    const [userInterests, setUserInterests] = useState([]);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        // Fetch the interests of the matched user
        const fetchInterests = async () => {
            try {
                const response = await axios.get(`/interests/${matched_user_id}`, {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                });
                console.log('Interests:', response.data);
                const values = Object.values(response.data || {}).flat();
                setUserInterests(values);

            } catch (error) {
                console.error('Error fetching interests:', error);
            } finally {
                setIsLoading(false);
            }
        };
        fetchInterests();
    }, [matched_user_id]);

    const handleViewMatchedProfile = () => {
        setModalOpen(true);
    };

    const handleCloseModal = () => {
        setModalOpen(false);
    };

    const handleRemoveMatch = async () => {
        try {
            const response = await axios.put('/matches/remove', { match_id }, {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log('Match Remove:', response.data);
            onUpdate(match_id);
        } catch (error) {
            console.error('Error removing the match:', error);
        }
        window.location.reload()
    };


    const handleConnectMatch = async () => {
        try {
            const response = await axios.put('/matches/connect', { match_id }, {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log('Match accepted:', response.data);
            onUpdate(match_id);
        } catch (error) {
            console.error('Error connecting to user:', error);
        }
        window.location.reload()
    };

    const handleRequestMatch = async () => {
        try {
            const response = await axios.put('/matches/request', { match_id }, {
                headers: {
                    Authorization: `Bearer ${authToken}`,
                },
            });
            console.log('Requested to match:', response.data);
            onUpdate(match_id);
            // You can implement additional logic like updating the UI or showing a success message
        } catch (error) {
            console.error('Error requesting to connect:', error);
        }
        window.location.reload()
    };

    const isRequester = requester === "true";

    const renderButtons = () => {
        switch (status) {
            case 'new':
                return (
                    <div className="match-card-button-section">
                        <button onClick={handleRequestMatch} className="match-card-button">
                            Request
                        </button>
                        <button onClick={handleRemoveMatch} className="match-card-button">
                            Dismiss match
                        </button>
                    </div>
                );
            case 'requested':
                return (
                    <>
                    <div className="match-card-button-section">
                        {!isRequester && (
                            <div className="match-card-buttons">
                                <button onClick={handleConnectMatch} className="">
                                    Connect
                                </button>
                            </div>
                        )}
                        {isRequester && (
                            <>
                                <div className="match-card-buttons">
                                    <h3 className="message" >Your Buddy Request is pending</h3>
                                </div>
                            </>
                        )}
                        <div className="match-card-buttons">
                            <button onClick={handleRemoveMatch} className="match-card-button">
                                Cancel Request
                            </button>
                        </div>
                    </div>
                    </>
                );
            case 'blocked':
                return (
                    <div className="match-card-button-section">
                        <p>
                            You are not authorized to contact this user.
                        </p>
                    </div>
                );
            default:
                return (
                    <div className="match-card-button-section">
                        <button onClick={handleRequestMatch} className="match-card-button">
                            Request
                        </button>

                        <button onClick={handleRemoveMatch} className="match-card-button">
                            Dismiss match
                        </button>
                    </div>
                );
        }
    };

    return (
        <>
            <div className="match-card">
                <div className="match-card-status">
                    <div className="user-name" ><h3>{matched_user_name}</h3></div>
                    {is_online ? <img src={onlineURL} alt="User online" className="status-icon"></img>
                        :
                        <img src={offlineURL} alt="User offline" className="status-icon"></img>
                    }
                </div>

                <div className="match-card-info">
                    <img className="match-card-image" src={userProfilePic} alt="User"></img>

                    <h2>{matched_user_location}</h2>
                    <h3>MatchScore: {match_score}</h3>


                </div>
                {renderButtons()}
                <h3>Interest:</h3>
                <div>
                    {userInterests.map((interest) => (
                        <button className="interest-button"
                            key={interest.id}
                        >
                            {interest}
                        </button>
                    ))}
                </div>
            </div>
            {/* <Modal isOpen={isModalOpen} onClose={handleCloseModal}>
                <BuddyCard buddyProfile={userProfile} />
            </Modal> */}
        </>
    )
}

export default MatchCard