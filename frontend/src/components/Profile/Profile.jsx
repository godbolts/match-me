import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Profile.css';
import defaultProfilePic from '../Assets/ProfilePictures/default_profile_pic.png';
import { CitySelect, CountrySelect, StateSelect } from 'react-country-state-city';
import 'react-country-state-city/dist/react-country-state-city.css';
import InterestSection from '../InterestSection/InterestSection';
import InterestPresenter from '../InterestPresenter/InterestPresenter';
import Modal from '../Modal/Modal.jsx';

const Profile = () => {
    const [countryId, setCountryId] = useState(0);
    const [stateId, setStateId] = useState(0);
    const [username, setUsername] = useState('');
    const [rawbirthdate, setrawBirthdate] = useState('');
    const [birthdate, setBirthdate] = useState('');
    const [about, setAboutMe] = useState('');
    const [usernameText, setUsernameText] = useState('');
    const [profilePic, setProfilePic] = useState(null);
    const [previewPic, setPreviewPic] = useState(null);
    const [isEditingLocation, setIsEditingLocation] = useState(false);
    const [formData, setFormData] = useState({
        country: '',
        state: '',
        city: '',
        latitude: null,
        longitude: null,
    });

// Modal 
    const [isModalOpen, setModalOpen] = useState(false);
    const handleEditUserInterests = () => {
        setModalOpen(true);
    };
    const handleCloseModal = () => {
        setModalOpen(false);
    };

    const authToken = localStorage.getItem('token');

    const formatDate = (date) => {
        if (!date) return '';
        const parsedDate = new Date(date);
        if (isNaN(parsedDate)) return '';
        return parsedDate.toISOString();
    };

    const formattedBirthdate = formatDate(birthdate);

    const formatDateForInput = (dateString) => {
        if (!dateString) return '';
        const parsedDate = new Date(dateString);
        if (isNaN(parsedDate)) return '';
        return parsedDate.toISOString().split('T')[0];
    };

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('/me/profile', {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                });
                const data = response.data;

                setUsername(data.username || '');
                setAboutMe(data.about_me || '');
                setrawBirthdate(data.birthdate || '');
                setCountryId(null);
                setStateId(null);
                setFormData({
                    country: data.user_nation || '',
                    state: data.user_region || '',
                    city: data.user_city || '',
                });

                if (data.profile_picture) {
                    setPreviewPic(`/uploads/${data.profile_picture}`);
                } else {
                    setPreviewPic(defaultProfilePic);
                }

                if (data.username && data.about_me && data.birthdate && data.user_nation) {
                    if (!localStorage.getItem('profileExists')) {
                        localStorage.setItem('profileExists', true);
                        console.log('localstorage item created!')
                        window.location.reload();
                    }
                }
            } catch (error) {
                console.error('Error fetching profile data:', error);
                localStorage.removeItem('profileExists');
            }
        };

        fetchData();
    }, [authToken]);

    const handleSubmitUsername = async () => {
        try {
            await axios.post(
                '/username',
                { username },
                {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                }
            );
            // alert('Username updated successfully!');
        } catch (error) {
            console.error('Error updating username:', error);
            alert('Failed to update username.');
        }
        window.location.reload();
    };

    const handleSubmitAboutMe = async () => {
        try {
            await axios.post(
                '/about',
                { about },
                {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                }
            );
            // alert('About Me updated successfully!');
        } catch (error) {
            console.error('Error updating About Me:', error);
            alert('Failed to update About Me.');
        }
    };

    const handleSubmitBirthdate = async () => {
        if (!birthdate) {
            alert('Please enter a valid birthdate.');
            return;
        }
        try {
            await axios.post(
                '/birthdate',
                { birthdate: formattedBirthdate },
                {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                }
            );
            // alert('Birthdate updated successfully!');
        } catch (error) {
            console.error('Error updating birthdate:', error);
            alert('Failed to update birthdate.');
        }
    };

    const handleSubmitProfilePic = async () => {
        if (!profilePic) {
            alert('Please select a profile picture.');
            console.log('No profile picture selected');
            return;
        }

        const picData = new FormData();
        picData.append('profilePic', profilePic);

        try {
            await axios.post('/picture', picData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                    Authorization: `Bearer ${authToken}`,
                },
            });
            // alert('Profile picture updated successfully!');
        } catch (error) {
            console.error('Error updating profile picture:', error);
            alert('Failed to update profile picture.');
        }
    };

    const handleSubmitRemovePicture = async () => {
        try {
            await axios.post(
                '/picture/remove',
                {},
                {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                }
            );
            window.location.reload();
        } catch (error) {
            console.error('Error removing profile picture:', error);
            alert('Failed to remove profile picture.');
        }
    };

    const handleSubmitLocation = async () => {
        if (!countryId || !stateId || !formData.city) {
            alert('Please select a valid country, state, and city.');
            return;
        }

        const payload = {
            ...formData,
            countryId,
            stateId,
        };

        try {
            await axios.post(
                '/city',
                payload,
                {
                    headers: {
                        Authorization: `Bearer ${authToken}`,
                    },
                }
            );
            // alert('Location updated successfully!');
            setIsEditingLocation(false);
        } catch (error) {
            console.error('Error updating location:', error);
            alert('Failed to update location.');
        }
    };

    const onCountryChange = (country) => {
        if (country?.id && country?.name) {
            setCountryId(country.id);
            setFormData((prevData) => ({
                ...prevData,
                country: country.name,
                state: '',
                city: '',
            }));
        }
    };

    const onStateChange = (state) => {
        if (state?.id && state?.name) {
            setStateId(state.id);
            setFormData((prevData) => ({
                ...prevData,
                state: state.name,
                city: '',
            }));
        }
    };

    const handleCitySelect = (city) => {
        if (city?.name) {
            setFormData((prevData) => ({
                ...prevData,
                city: city.name,
                latitude: city.latitude || null,
                longitude: city.longitude || null,
            }));
        }
    };

    const handleImageChange = (e) => {
        const file = e.target.files[0];
        if (file) {
            setProfilePic(file);
        }

        const reader = new FileReader();
        reader.onloadend = () => {
            setPreviewPic(reader.result);
        };
        reader.readAsDataURL(file);
        handleSubmitProfilePic();
    };

    useEffect(() => {
        setUsernameText("Username");
        const formattedDate = formatDateForInput(rawbirthdate);
        setBirthdate(formattedDate);
    }, [rawbirthdate]);

    return (
        <body className="profile-body">
        <main className="profile-main">
            <section className="profile-left" style={{Align: 'centre'}}>
                {/* Profile Picture Section */}
                        <div className='input-profile-pic'>
                            
                            <label htmlFor="file-input" className="profile-pic-label">
                                {previewPic ? (
                                    <img src={previewPic} alt="Preview" />
                                ) : (
                                    <img src={defaultProfilePic} alt="Default Profile" />
                                )}
                            </label>
                            <input
                                id="file-input"
                                type="file"
                                accept="image/*"
                                onChange={handleImageChange}
                                style={{ display: 'none' }}
                            />
                        </div>
                        <p className='suggestion'>click on the image to select a new profile image</p>
                        <div className='submit-container'>
                            <button
                                className='submit'
                                onClick={handleSubmitProfilePic}
                            >
                                Update Picture
                            </button>
                            <button
                                className='submit'
                                onClick={handleSubmitRemovePicture}
                            >
                                Remove Picture
                            </button>
                            
                        </div>
                {/* Username Section */}
                        <div className='profile-text'>{usernameText}</div>
                       <div className='input'>
                           <input
                               type='text'
                               placeholder='Username'
                               maxLength="20"
                               value={username}
                               onChange={(e) => setUsername(e.target.value)}
                               required
                           />
                       </div>
                       <div className='submit-container'>
                           <button
                               className='submit'
                               onClick={handleSubmitUsername}
                           >
                            Update Username
                           </button>
                       </div>
                       {/* About Me Section */}
                        <div className='profile-text'>About Me</div>
                        <div className='input-textarea'>
                            <textarea
                                placeholder='About me'
                                maxLength="500"
                                value={about}
                                onChange={(e) => setAboutMe(e.target.value)}
                                required
                            />
                        </div>
                        <div className='submit-container'>
                            <button
                                className='submit'
                                onClick={handleSubmitAboutMe}
                            >
                                Update About Me
                            </button>
                        </div>
            </section>
            <section className="profile-right">
                {/* Interests Modal */}
                        <div className='interest-display'>
                        <div className='profile-text'>Interests</div>
                        <InterestPresenter/>
                            <div className="submit-container">
                                <button className='submit' onClick={handleEditUserInterests}>
                                    Edit Interests
                                </button>
                            </div>
    
                        <Modal isOpen={isModalOpen} onClose={handleCloseModal}>
                            <InterestSection />
                        </Modal>
                        </div>
                        {/* Birthdate Section */}
                        <div className='interest-display'>
                        <div className='profile-text'>Birthday</div>
                        <div className='input'>
                            <input
                                type="date"
                                value={birthdate === '0001-01-01' ? '2000-01-01' : birthdate}
                                onChange={(e) => setBirthdate(e.target.value)}
                                required
                            />
                        </div>
                        <div className='submit-container'>
                            <button
                                className='submit'
                                onClick={handleSubmitBirthdate}
                            >
                                Update Birthdate
                            </button>
                        </div>
                        </div>
                        {/* Location Section */}
                        <div className='profile-text'>Location</div>
                        <div className="interest-display">
                            {!isEditingLocation ? (
                                <>
                                    <p><strong className='location-text'>Country:</strong> {formData.country || 'Not Set'}</p>
                                    <p><strong className='location-text'>State:</strong> {formData.state || 'Not Set'}</p>
                                    <p><strong className='location-text'>City:</strong> {formData.city || 'Not Set'}</p>
                                    <div className='submit-container'>
                                        <button
                                            className='submit'
                                            onClick={() => setIsEditingLocation(true)}
                                        >
                                            Edit Location
                                        </button>
                                    </div>
                                </>
                            ) : (
                                <div className="inputGroupLocation">
                                    <div className="inputField">
                                        <h6>Country:</h6>
                                        <CountrySelect
                                            onChange={onCountryChange}
                                            placeHolder={formData.country || "Select Country"}
                                        />
                                    </div>
                                    <div className="inputField">
                                        <h6>State:</h6>
                                        <StateSelect
                                            countryid={countryId}
                                            onChange={onStateChange}
                                            placeHolder={formData.state || "Select State"}
                                            disabled={!countryId}
                                        />
                                    </div>
                                    <div className="inputField">
                                        <h6>City:</h6>
                                        <CitySelect
                                            countryid={countryId}
                                            stateid={stateId}
                                            onChange={handleCitySelect}
                                            placeHolder={formData.city || "Select City"}
                                            disabled={!countryId || !stateId}
                                        />
                                    </div>
                                    <div className='submit-container'>
                                        <button
                                            className='submit'
                                            onClick={() => {
                                                handleSubmitLocation();
                                                setIsEditingLocation(false);
                                            }}
                                        >
                                            Update Location
                                        </button>
                                        <button
                                            className='submit'
                                            onClick={() => setIsEditingLocation(false)}
                                        >
                                            Cancel
                                        </button>
                                    </div>
                                </div>
                            )}
                        </div>
            </section>
        </main>
    </body>
    );

};

export default Profile;