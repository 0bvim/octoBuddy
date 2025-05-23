import Link from 'next/link'

export default function Home() {
  return (
    <main className="min-h-screen flex items-center justify-center">
      <div className="text-center">
        <h1 className="text-4xl font-bold mb-8">Welcome to OctoBuddy</h1>
        <Link 
          href="http://localhost:8080/auth/github"
          className="bg-gray-900 text-white px-6 py-3 rounded-lg hover:bg-gray-800 transition-colors"
        >
          Login with GitHub
        </Link>
      </div>
    </main>
  )
}