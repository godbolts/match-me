import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import './Matches.css'
import '../MatchCard/MatchCard.jsx'
import MatchCard from '../MatchCard/MatchCard.jsx';

const authToken = localStorage.getItem('token');

const Matches = () => {
    const [loading, setLoading] = useState(true)
    const [data, setData] = useState([])
    const [matches, setMatches] = useState([])


    useEffect(() => {
        const fetchMatches = async () => {
            setLoading(true);
            try {
                const response = await axios.get('/matches', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                });
                console.log(response.data);
                setMatches(response.data);
            }
            catch (error) {
                console.error('Error fetching data: ', error)
            } 
        }
        fetchMatches();
    }, [])

    const handleMatchUpdate =(match_id) => {
        // setMatches((prevMatches) => {prevMatches.filter((match) => match.match_id !== match_id)});
        window.location.reload()
    }

    return (
        <>
            <h1>Matches</h1>
            <div className="body-div">
                <div className="body-sides"></div>
                <div className="body-content">
                {Array.isArray(matches) && matches.length > 0 ? (
                    matches.map((item) => (
                        <MatchCard
                            key={item.match_id}
                            userProfile={item}
                            onUpdate={handleMatchUpdate} // Pass callback to update matches
                        />
                    ))
                ) : (
                    <p>No matches found. Try updating your preferences or check back later!</p>
                )}
                {Array.isArray(matches) && matches.length === 10 && ( 
                    <button
                        className="load-more-button"
                        onClick={() => window.location.reload()} // Reload the page
                    >
                        Load More Matches
                    </button>
                )}
                    </div>
                <div className="body-sides"></div>
            </div>
        </>
    );
};

export default Matches;