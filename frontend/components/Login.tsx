import React, { useState } from 'react';
const axios = require('axios')

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (event: { preventDefault: () => void; }) => {
    event.preventDefault();
    try {
      const response = await axios.post('api/v1/login', { email, password });
      // save the authentication token in local storage or cookies
      localStorage.setItem('token', response.data.token);
      // redirect to dashboard or home page
      window.location.href = '/dashboard';
    } catch (error) {
      alert('error');
    }
  };

  return (
  <div>
    <h2 style={{ color: "#333", marginBottom: "20px" }}>Login</h2>
    <form onSubmit={handleSubmit} style={{ display: "flex", flexDirection: "column" }}>
      <div style={{ marginBottom: "10px" }}>
        <label style={{ display: "block", marginBottom: "5px" }}>Email:</label>
        <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} style={{ padding: "5px", borderRadius: "5px", border: "1px solid #ccc", color: "#333" }} placeholder="Enter your email" />
      </div>
      <div style={{ marginBottom: "10px" }}>
        <label style={{ display: "block", marginBottom: "5px" }}>Password:</label>
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} style={{ padding: "5px", borderRadius: "5px", border: "1px solid #ccc", color: "#333" }} placeholder="Enter your password" />
      </div>
      <button type="submit" style={{ backgroundColor: "#4CAF50", color: "white", padding: "10px", borderRadius: "5px", border: "none" }}>Login</button>
    </form>
  </div>
  

  );
}

export default Login;
