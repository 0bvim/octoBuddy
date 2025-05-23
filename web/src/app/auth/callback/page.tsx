'use client'

import { useEffect } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { authService } from '@/lib/auth';

export default function AuthCallback() {
  const router = useRouter();
  const searchParams = useSearchParams();
  
  useEffect(() => {
    const tokens = {
      access_token: searchParams.get('token') || '',
      refresh_token: searchParams.get('refresh_token') || '',
    };

    if (tokens.access_token && tokens.refresh_token) {
      authService.setTokens(tokens);
      router.push('/dashboard');
    } else {
      router.push('/');
    }
  }, [router, searchParams]);

  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="text-center">
        <h1 className="text-2xl font-bold mb-4">Authenticating...</h1>
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900 mx-auto"></div>
      </div>
    </div>
  );
}