'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

interface GitHubUser {
  login: string;
  id: number;
  node_id: string;
  avatar_url: string;
  gravatar_id: string;
  url: string;
  html_url: string;
  followers_url: string;
  following_url: string;
  gists_url: string;
  starred_url: string;
  subscriptions_url: string;
  organizations_url: string;
  repos_url: string;
  events_url: string;
  received_events_url: string;
  type: string;
  site_admin: boolean;
  name: string;
  company: string;
  blog: string;
  location: string;
  email: string;
  hireable: boolean;
  bio: string;
  twitter_username: string;
  public_repos: number;
  public_gists: number;
  followers: number;
  following: number;
  created_at: string; // ISO date string
  updated_at: string; // ISO date string
}

export default function HomePage() {
  const [user, setUser] = useState<GitHubUser | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const router = useRouter();

  useEffect(() => {
    fetch('http://localhost:3000/callback', {
      credentials: 'include',
      headers: {
        'Accept': 'application/json',
      }
    })
      .then(async (response) => {
        if (!response.ok) {
          if (response.status === 401) {
            router.push('/');
            throw new Error('Please login first');
          }
          const errorText = await response.text();
          throw new Error(errorText || 'Failed to fetch user data');
        }
        return response.json();
      })
      .then((data: GitHubUser) => {
        setUser(data);
        setLoading(false);
      })
      .catch((err: Error) => {
        console.error('Error fetching user data:', err);
        setError(err.message);
        setLoading(false);
      });
  }, [router]);

  if (loading) {
    return (
      <div className="container text-center mt-5">
        <p>Loading user data...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="container text-center mt-5">
        <p>Error: {error}</p>
      </div>
    );
  }

  if (!user) {
    return null;
  }

  return (
    <div className="container mt-5">
      <div className="text-center">
        <img
          src={user.avatar_url}
          alt="GitHub Profile Picture"
          style={{ width: '150px', height: '150px', borderRadius: '50%' }}
        />
        <h1 className="mt-3">{user.name || user.login}</h1>
        {user.bio && <p>{user.bio}</p>}
        <div className="d-flex justify-content-center mt-3">
          <div className="mx-3">
            <strong>Followers:</strong> {user.followers}
          </div>
          <div className="mx-3">
            <strong>Following:</strong> {user.following}
          </div>
          <div className="mx-3">
            <strong>Public Repos:</strong> {user.public_repos}
          </div>
          <div className="mx-3">
            <strong>Public Gists:</strong> {user.public_gists}
          </div>
        </div>
      </div>
    </div>
  );
}
