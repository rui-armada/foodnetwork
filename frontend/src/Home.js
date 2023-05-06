import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Home.css';

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
    <div className="home">
      <header className="home-header">
        <h1>Home</h1>
        <button onClick={handleLogout} className="logout-btn">Logout</button>
      </header>
      <div className="create-post">
        <h3>Create a new post:</h3>
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="create-post-textarea"
        />
        <button onClick={handleCreatePost} className="create-post-btn">
          Create Post
        </button>
      </div>
      <div className="posts">
        <h2>All Posts:</h2>
        <div className="posts-grid">
          {posts.map((post) => (
            <div key={post.id} className="post">
              <p className="post-author">Posted by: {post.userName}</p>
              <p className="post-description">{post.description}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default Home;
