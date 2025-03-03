import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './InterestPresenter.css';

const InterestPresenter = () => {
  const [bioValues, setBioValues] = useState([]);
  const authToken = localStorage.getItem('token');

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('/me/bio', { // Need to use a different endpoint this endpoint does not ger refreshed for some reason.
          headers: {
            Authorization: `Bearer ${authToken}`,
          },
        });
        console.log('Fetched data:', response.data);

        // Extract and store only the values from the response
        const values = Object.values(response.data).flat();
        setBioValues(values);
        console.log(values)
      } catch (error) {
        console.error('Error fetching bio data:', error);
      }
    };

    fetchData();
  }, [authToken]);

  return (
    <div className="interest-sections">
      {bioValues.length > 0 ? (
        <div className="interests-container">
          {bioValues.map((value, index) => (
            <button key={index} className="selected-interest-presenter">
              {value}
            </button>
          ))}
        </div>
      ) : (
        <></>
      )}
    </div>
  );

};

export default InterestPresenter;
