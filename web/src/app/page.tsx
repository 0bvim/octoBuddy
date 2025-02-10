// src/app/page.tsx
'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

export default function HomePage() {
  const router = useRouter();
  const [error, setError] = useState('');

  useEffect(() => {
    // Normal page load - check if already logged in
    fetch('http://localhost:8080/login', {
      credentials: 'include',
    })
      .then(async (response) => {
        if (response.ok) {
          const data = await response.json();
          // Store user data in localStorage for home page
          localStorage.setItem('userData', JSON.stringify(data));
          router.push('/home');
        } else {
          if (response.status === 401) {
            // Not logged in, show login page (default behavior)
            return;
          }
          const errorText = await response.text();
          throw new Error(errorText || 'Failed to fetch user data');
        }
      })
      .catch((err) => {
        console.error('Error:', err);
        setError(err.message);
      });
  }, [router]);

  return (
    <main className="container d-flex flex-column justify-content-center align-items-center vh-100">
      <h1 className="display-4">OctoBuddy</h1>
      <p className="lead">Welcome to OctoBuddy. Please log in to continue.</p>
      {error && <p className="text-danger mb-3">{error}</p>}
      <a href="http://localhost:8080/login" className="btn btn-dark btn-lg">
        <i className="fab fa-github me-2"></i>
        Login with GitHub
      </a>
    </main>
  );
}