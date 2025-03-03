import React, { useEffect, useState } from 'react';
import './Requests.css';

import axios from 'axios';
import './Requests.css'
import MatchCard from '../MatchCard/MatchCard.jsx';


    const authToken = localStorage.getItem('token');

    const Requests = () => {

        const [data, setData] = useState([])
        const [matches, setMatches] = useState([])

        useEffect(() => {
            const fetchData = async () => {
                try {
                    console.log('Try to get the requests');
                    const response = await axios.get('/requests', {
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
            fetchData();
        }, [])

        return (
            <>
            <h1>Requests</h1>
            <div className="body-div">
                <div className="body-sides"></div>
                <div className="body-content">
                    {matches && matches.length > 0 ? (
                        matches.map((item, index) => (
                            <p key={index}>
                                <MatchCard userProfile={item} key={index} />
                            </p>
                        ))
                    ) : (
                        <p>No matches found. Try updating your preferences or check back later!</p>
                    )}
                    </div>
                <div className="body-sides"></div>
            </div>
            </>
        );
    }
    
export default Requests;