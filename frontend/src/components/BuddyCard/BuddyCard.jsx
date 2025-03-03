import React, { useEffect, useState, useContext } from 'react';
import './BuddyCard.css';
import axios from 'axios';
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';
import InterestPresenter from '../InterestPresenter/InterestPresenter';




const authToken = localStorage.getItem('token');

// The buddy card take in the data from the the /matches API and get the interest from the /interests API
const BuddyCard = ({ buddyProfile, onUpdate }) => {




    const { match_id,
        match_score,
        status,
        is_online,
        requester,
        matched_user_id,
        matched_user_name, 
        matched_user_picture,
        matched_user_description, 
        matched_user_location } = buddyProfile;  

        console.log("buddy profile: " + buddyProfile.matched_user_id)
        console.log("buddy profile: " + buddyProfile.requester)


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

    
    const renderButtons = () => {
        switch (status) {
            case 'connected':
                return (
                    <>
                        <button onClick={handleRemoveMatch} className="match-card-button">
                            Dismiss match
                        </button>
                    </>
                );
            case 'blocked':
                return (
                    <>
                        <p>
                            You are not authorized to contact this user.
                        </p>
                    </>
                );
            default:
                return (
                    <>
                        <button onClick={handleRemoveMatch} className="match-card-button">
                            Un-Buddy
                        </button>
                    </>
                );
        }
    };


    
    return (
        <>
            <div className="match-card">
                    <div className="match-card-status">
                    <h3 className="user-name" >{matched_user_name}</h3>  
                    {is_online ? <img src={onlineURL} alt="User online" className="status-icon"></img>
                        :
                    <img src={offlineURL} alt="User offline" className="status-icon"></img>
                    }
                    </div>
                    <div className="match-card-info">
                        <img className="match-card-image" src={userProfilePic} alt ="User"></img>
                        <h3>Location: {matched_user_location }</h3>
                        <h3>MatchScore: {match_score}</h3>
                        <p>{matched_user_description}</p>                   
                    <div >
                        {userInterests.map((interest) => (
                        <button className="interest-button"
                        key={interest.id}
                        >
                        {interest}
                        </button>
                        ))} 
                    </div>  
                    </div>
                    <div className="match-card-buttons">
                        {renderButtons()}   
                        <button onClick= {() => window.location.href = '/chat'}   className="match-card-button">
                            Chat
                        </button>
                    </div>
                </div>                 
    </>
);

}

export default BuddyCard;  