'use client';

import { useState, useEffect } from 'react';
import { Todo, todoApi } from '@/lib/api';
import TodoItem from './TodoItem';
import AddTodoForm from './AddTodoForm';
import { ListTodo } from 'lucide-react';

export default function TodoList() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchTodos = async () => {
    try {
      const data = await todoApi.getTodos();
      console.log('Fetched todos:', data);
      // データが配列でない場合は空配列にする
      if (Array.isArray(data)) {
        setTodos(data);
      } else {
        console.warn('Data is not an array:', data);
        setTodos([]);
      }
    } catch (error) {
      console.error('Failed to fetch todos:', error);
      setTodos([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const handleTodoAdded = () => {
    fetchTodos();
  };

  const handleTodoUpdated = () => {
    fetchTodos();
  };

  const handleTodoDeleted = () => {
    fetchTodos();
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <AddTodoForm onTodoAdded={handleTodoAdded} />
      
      {todos.length === 0 ? (
        <div className="text-center py-12">
          <ListTodo size={64} className="mx-auto text-blue-300 mb-4" />
          <p className="text-blue-500 text-lg">No todos yet. Create your first one!</p>
        </div>
      ) : (
        <div>
          <div className="mb-4 flex justify-between items-center">
            <h2 className="text-xl font-semibold text-blue-800">Your Tasks</h2>
            <span className="text-sm text-blue-600">
              {todos.filter(t => !t.completed).length} of {todos.length} remaining
            </span>
          </div>
          <div>
            {todos.map((todo, index) => (
              <TodoItem
                key={`todo-${todo.id}-${index}`}
                todo={todo}
                onUpdate={handleTodoUpdated}
                onDelete={handleTodoDeleted}
              />
            ))}
          </div>
        </div>
      )}
    </div>
  );
}