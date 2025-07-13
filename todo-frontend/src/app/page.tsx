import TodoList from '@/components/TodoList';
import { CheckSquare } from 'lucide-react';

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50">
      <div className="container mx-auto px-4 py-8 max-w-4xl">
        <header className="text-center mb-10">
          <div className="flex items-center justify-center gap-3 mb-4">
            <CheckSquare size={40} className="text-blue-500" />
            <h1 className="text-4xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
              Todo Manager
            </h1>
          </div>
          <p className="text-blue-700">Organize your tasks efficiently</p>
        </header>
        
        <main className="bg-white/80 backdrop-blur-sm rounded-2xl shadow-xl p-6 md:p-8">
          <TodoList />
        </main>
        
        <footer className="text-center mt-8 text-sm text-blue-500">
          <p>Built with Next.js & Go</p>
        </footer>
      </div>
    </div>
  );
}