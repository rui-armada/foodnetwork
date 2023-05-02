import React from 'react';

function Home() {
  const handleLogout = () => {
    localStorage.removeItem('token');
    window.location.href = '/';
  };

  return (
    <div>
      <h1 style={{ color: "#333", marginBottom: "20px" }}>Welcome to the Home Page!</h1>
      <p style={{ marginBottom: "20px" }}>
        You have successfully logged in. This is a sample home page.
      </p>
      <button
        onClick={handleLogout}
        style={{
          backgroundColor: "#f44336",
          color: "white",
          padding: "10px",
          borderRadius: "5px",
          border: "none",
        }}
      >
        Logout
      </button>
    </div>
  );
}

export default Home;
