import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Home() {
  const [posts, setPosts] = useState([]);
  const [description, setDescription] = useState('');

  useEffect(() => {
    fetchPosts();
  }, []);

  const fetchPosts = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/v1/posts');
      setPosts(response.data.data);
    } catch (error) {
      console.error('Error fetching posts:', error);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    window.location.href = '/';
  };

  const handleCreatePost = async () => {
    const token = localStorage.getItem('token');    
  
    try {
      await axios.post(
        'http://localhost:8080/api/v1/posts',
        {description},
        {
          headers: { Authorization: `Bearer ${token}` },
        },
      );
      setDescription('');
      fetchPosts();
    } catch (error) {
      console.error('Error creating post:', error);
    }
  };
  
  return (
    <div>
      <h1>Home</h1>
      <button onClick={handleLogout}>Logout</button>
      <div>
        <h3>Create a new post:</h3>
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          style={{ width: '100%', marginBottom: '10px' }}
        />
        <button
          onClick={handleCreatePost}
          style={{
            backgroundColor: '#4CAF50',
            color: 'white',
            padding: '10px',
            borderRadius: '5px',
            border: 'none',
            marginBottom: '20px',
          }}
        >
          Create Post
        </button>
      </div>
      <div>
        <h2>All Posts:</h2>
        <ul>
          {posts.map((post) => (
            <li key={post.id}>
              <p>Posted by: {post.userName}</p>
              <p>{post.description}</p>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default Home;
