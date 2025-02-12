'use client'

import { useAuth } from '@/hooks/useAuth';
import { authService } from '@/lib/auth';
import { useRouter } from 'next/navigation';

export default function Dashboard() {
  const router = useRouter();
  const { user, loading } = useAuth();

  const handleLogout = () => {
    authService.clearTokens();
    router.push('/');
  };

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  if (!user) {
    return null;
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 to-blue-50 p-8">
      <div className="max-w-4xl mx-auto">
        {/* Header Section */}
        <div className="flex items-center justify-between mb-8 bg-white p-6 rounded-2xl shadow-lg">
          <div className="flex items-center space-x-4">
            <img 
              src={user.avatar_url} 
              alt={user.name} 
              className="w-16 h-16 rounded-full border-4 border-blue-100 shadow-sm"
            />
            <div>
              <h1 className="text-2xl font-bold text-gray-800">{user.name}</h1>
              <p className="text-gray-600 flex items-center">
                @{user.login}
                <span className="ml-2 bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded-full">
                  {user.followers} followers
                </span>
              </p>
            </div>
          </div>
          <button
            onClick={handleLogout}
            className="flex items-center bg-white text-gray-600 px-4 py-2 rounded-lg border border-gray-200 hover:border-red-500 hover:text-red-600 transition-all"
          >
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
              <path fillRule="evenodd" d="M3 3a1 1 0 00-1 1v12a1 1 0 102 0V4a1 1 0 00-1-1zm10.293 9.293a1 1 0 001.414 1.414l3-3a1 1 0 000-1.414l-3-3a1 1 0 10-1.414 1.414L14.586 9H7a1 1 0 100 2h7.586l-1.293 1.293z" clipRule="evenodd" />
            </svg>
            Logout
          </button>
        </div>
  
        {/* Stats Section */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
          <div className="bg-white p-6 rounded-xl shadow-md hover:shadow-lg transition-shadow">
            <div className="text-blue-600 font-bold text-2xl">{user.public_repos}</div>
            <div className="text-gray-500 text-sm">Public Repos</div>
          </div>
          <div className="bg-white p-6 rounded-xl shadow-md hover:shadow-lg transition-shadow">
            <div className="text-green-600 font-bold text-2xl">{user.following}</div>
            <div className="text-gray-500 text-sm">Following</div>
          </div>
          <div className="bg-white p-6 rounded-xl shadow-md hover:shadow-lg transition-shadow">
            <div className="text-purple-600 font-bold text-2xl">{user.followers}</div>
            <div className="text-gray-500 text-sm">Followers</div>
          </div>
        </div>
  
        {/* Follow/Unfollow Section */}
        <div className="bg-white p-6 rounded-2xl shadow-lg mb-8">
          <h2 className="text-xl font-bold text-gray-800 mb-6 flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 mr-2 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
            Discover Users
          </h2>
          
          {/* Search Bar */}
          <div className="mb-6 relative">
            <input
              type="text"
              placeholder="Search GitHub users..."
              className="w-full px-4 py-3 rounded-lg border border-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 absolute right-3 top-3.5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
  
          {/* Users List */}
          {/* <div className="space-y-4">
            {users.map((user) => (
              <div key={user.id} className="flex items-center justify-between p-4 hover:bg-gray-50 rounded-lg transition-colors">
                <div className="flex items-center space-x-4">
                  <img 
                    src={user.avatar_url} 
                    alt={user.login} 
                    className="w-10 h-10 rounded-full border-2 border-blue-100"
                  />
                  <div>
                    <div className="font-medium text-gray-800">{user.login}</div>
                    <div className="text-sm text-gray-500">{user.name || "GitHub User"}</div>
                  </div>
                </div>
                <button
                  onClick={() => toggleFollow(user)}
                  className={`px-4 py-2 rounded-full text-sm font-medium transition-all ${
                    user.isFollowing 
                      ? "bg-red-100 text-red-600 hover:bg-red-200"
                      : "bg-blue-100 text-blue-600 hover:bg-blue-200"
                  }`}
                >
                  {user.isFollowing ? "Unfollow" : "Follow"}
                </button>
              </div>
            ))} */}
          {/* </div> */}
        </div>
      </div>
    </div>
  );
}