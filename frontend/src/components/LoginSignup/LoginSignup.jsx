import React, { useState } from 'react';
import './LoginSignup.css';
import email_icon from '../Assets/email.png';
import password_icon from '../Assets/password.png';
import axios from 'axios';

const LoginSignup = () => {

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [loading, setLoading] = useState(false);
    const [errorMessage, setErrorMessage] = useState(null);
    const [login, setLogin] = useState(false);
    const [register, setRegister] = useState(false);

    const handleSubmit = async (e, isLoginState) => {
        e.preventDefault();
        setLoading(true);
        //setErrorMessage(null);

        const url = isLoginState ? '/login' : '/register';

        const payload = {
            email,
            password
        }

        try {
            const response = await axios.post(url, payload);

            if (isLoginState) {
                localStorage.setItem('token', response.data.token)
                //alert('Login successful') - probably don't need alerts, just use errorMessage to send a message and then redirect
                setErrorMessage('Login successful')
                window.location.href = '/profile';
            } else {
                //alert('Sign-up successful, please log in')
                setErrorMessage('Sign-up successful, please log in')
            }
        } catch (error) {
            setErrorMessage(null);
            const errorMessage = error.response?.data?.error || 'An error occurred, please try again.'
            setErrorMessage(errorMessage)
        } finally {
            setLoading(false);
            setLogin(false);
            setRegister(false);
        }
    };

    return (
        <div>
            <div className='login-container'>
                <div className='login-header'>
                    <div className='login-text'>{errorMessage ? errorMessage : 'Sign Up or Log In'}</div>
                </div>
                <div className='inputs'>
                    <div className='input'>
                        <img src={email_icon} alt="" />
                        <input
                            type='email'
                            placeholder='Email'
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            required />
                    </div>
                    <div className='input'>
                        <img src={password_icon} alt="" />
                        <input
                            type='password'
                            placeholder='Password'
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            required />
                    </div>
                </div>
                <div className='submit-container'>
                    <div
                        className='submit'
                        onClick={(e) => {
                            handleSubmit(e, false);
                            setRegister(true)
                        }}>
                        {loading && register ? 'Signing up...' : 'Sign up'}
                    </div>
                    <div
                        className='submit'
                        onClick={(e) => {
                            handleSubmit(e, true);
                            setLogin(true)
                        }}>
                        {loading && login ? 'Logging in...' : 'Log In'}
                    </div>
                </div>
            </div>
        </div>
    )
}

export default LoginSignup