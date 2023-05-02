import React, { useState } from 'react';
import axios from 'axios';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [name, setName] = useState('');
  const [isRegister, setIsRegister] = useState(false); // track whether the user is registering or logging in

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const route = isRegister ? 'http://localhost:8080/api/v1/register' : 'http://localhost:8080/api/v1/login';
      const response = await axios.post(route, { email, password, name });
      // save the authentication token in local storage or cookies
      localStorage.setItem('token', response.data.token);
      // redirect to dashboard or home page
      window.location.href = '/home';
    } catch (error) {
      console.error('Error in API call:', error.response ? error.response.data : error.message);
      alert('Error: ' + (error.response ? error.response.data : error.message));
    }
  };
  

  return (
    <div>
      <h2 style={{ color: "#333", marginBottom: "20px" }}>{isRegister ? 'Register' : 'Login'}</h2>
      <form onSubmit={handleSubmit} style={{ display: "flex", flexDirection: "column" }}>
        {isRegister && (
          <div style={{ marginBottom: "10px" }}>
            <label style={{ display: "block", marginBottom: "5px" }}>Name:</label>
            <input type="text" value={name} onChange={(e) => setName(e.target.value)} style={{ padding: "5px", borderRadius: "5px", border: "1px solid #ccc", color: "#333" }} placeholder="Enter your name" />
          </div>
        )}
        <div style={{ marginBottom: "10px" }}>
          <label style={{ display: "block", marginBottom: "5px" }}>Email:</label>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} style={{ padding: "5px", borderRadius: "5px", border: "1px solid #ccc", color: "#333" }} placeholder="Enter your email" />
        </div>
        <div style={{ marginBottom: "10px" }}>
          <label style={{ display: "block", marginBottom: "5px" }}>Password:</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} style={{ padding: "5px", borderRadius: "5px", border: "1px solid #ccc", color: "#333" }} placeholder="Enter your password" />
        </div>
        <button type="submit" style={{ backgroundColor: "#4CAF50", color: "white", padding: "10px", borderRadius: "5px", border: "none" }}>{isRegister ? 'Register' : 'Login'}</button>
        <p style={{ marginTop: "10px" }}>
          {isRegister ? 'Already have an account?' : 'Don\'t have an account yet?'}
          <button type="button" onClick={() => setIsRegister(!isRegister)} style={{ marginLeft: "5px", backgroundColor: "transparent", color: "#4CAF50", border: "none", cursor: "pointer", textDecoration: "underline" }}>{isRegister ? 'Login' : 'Register'}</button>
        </p>
      </form>
    </div>
  );
}

export default Login;
