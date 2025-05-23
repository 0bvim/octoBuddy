interface tokenPair {
    access_token: string;
    refresh_token: string;
  }
  
  interface user {
    id: string;
    login: string;
    name: string;
    email: string;
    avatar_url: string;
    followers: number;
    following: number;
    company: string;
    location: string;
    public_repos: number;
    followers_url: string;
    following_url: string;
    // follower: follower[];
    // followed: followed[];
  }
  
  class AuthService {
    private static instance: AuthService;
    private refreshPromise: Promise<tokenPair> | null = null;
  
    private constructor() {}
  
    static getInstance(): AuthService {
      if (!AuthService.instance) {
        AuthService.instance = new AuthService();
      }
      return AuthService.instance;
    }
  
    getTokens(): tokenPair | null {
      if (typeof window === 'undefined') return null;
      
      const tokens = localStorage.getItem('auth_tokens');
      return tokens ? JSON.parse(tokens) : null;
    }
  
    setTokens(tokens: tokenPair) {
      localStorage.setItem('auth_tokens', JSON.stringify(tokens));
    }
  
    clearTokens() {
      localStorage.removeItem('auth_tokens');
    }
  
    async refreshTokens(): Promise<tokenPair> {
      // Prevent multiple refresh calls
      if (this.refreshPromise) {
        return this.refreshPromise;
      }
  
      const tokens = this.getTokens();
      if (!tokens?.refresh_token) {
        throw new Error('No refresh token available');
      }
  
      this.refreshPromise = fetch('http://localhost:8080/auth/refresh', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ refresh_token: tokens.refresh_token }),
      }).then(async (res) => {
        if (!res.ok) {
          throw new Error('Failed to refresh tokens');
        }
        const newTokens = await res.json();
        this.setTokens(newTokens);
        return newTokens;
      }).finally(() => {
        this.refreshPromise = null;
      });
  
      return this.refreshPromise;
    }
  
    async fetchWithAuth(url: string, options: RequestInit = {}): Promise<Response> {
      const tokens = this.getTokens();
      if (!tokens) {
        throw new Error('No authentication tokens');
      }
  
      const headers = new Headers(options.headers);
      headers.set('Authorization', `Bearer ${tokens.access_token}`);
  
      try {
        const response = await fetch(url, { ...options, headers });
        
        if (response.status === 401) {
          // Token expired, try to refresh
          const newTokens = await this.refreshTokens();
          headers.set('Authorization', `Bearer ${newTokens.access_token}`);
          return fetch(url, { ...options, headers });
        }
  
        return response;
      } catch (error) {
        if (error instanceof Error && error.message === 'Failed to refresh tokens') {
          this.clearTokens();
          window.location.href = '/';
        }
        throw error;
      }
    }
  }
  
  export const authService = AuthService.getInstance();

  export type { tokenPair, user };