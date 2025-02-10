// src/app/callback/page.tsx
'use client';

import { useEffect } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';

export default function CallbackPage() {
  const router = useRouter();
  const searchParams = useSearchParams();

  useEffect(() => {
    const code = searchParams.get('code');
    const state = searchParams.get('state');

    if (!code || !state) {
      router.push('/');
      return;
    }

    // TODO: verify how to make it get callback from the backend
    // Call your backend's callback endpoint
    fetch(`http://localhost:8080/callback?code=${code}&state=${state}`, {
      method: 'GET',
      credentials: 'include',
      headers: {
        'Accept': 'application/json',
      },
    })
      .then(async (response) => {
        if (!response.ok) {
          throw new Error('Authentication failed');
        }
        return response.json();
      })
      .then(() => {
        // On successful authentication, redirect to home
        router.push('/home');
      })
      .catch((error) => {
        console.error('Authentication error:', error);
        router.push('/');
      });
  }, [router, searchParams]);

  return (
    <div className="container text-center mt-5">
      <p>Authenticating...</p>
    </div>
  );
}