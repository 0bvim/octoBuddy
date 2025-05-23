import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { authService, user } from '@/lib/auth';

export function useAuth() {
  const router = useRouter();
  const [loading, setLoading] = useState(true);
  const [user, setUser] = useState<user | null>(null);

  useEffect(() => {
    const checkAuth = async () => {
      try {
        const response = await authService.fetchWithAuth('http://localhost:8080/api/user');
        if (!response.ok) throw new Error('Failed to fetch user');
        const userData = await response.json();
        setUser(userData);
      } catch (error) {
        authService.clearTokens();
        router.push('/');
      } finally {
        setLoading(false);
      }
    };

    if (authService.getTokens()) {
      checkAuth();
    } else {
      setLoading(false);
    }
  }, [router]);

  return { user, loading };
}