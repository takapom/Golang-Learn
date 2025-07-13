'use client';

import { useState } from 'react';
import { todoApi } from '@/lib/api';
import { Plus } from 'lucide-react';

interface AddTodoFormProps {
  onTodoAdded: () => void;
}

export default function AddTodoForm({ onTodoAdded }: AddTodoFormProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim()) return;

    setIsSubmitting(true);
    try {
      await todoApi.createTodo({
        title: title.trim(),
        description: description.trim(),
      });
      setTitle('');
      setDescription('');
      setIsOpen(false);
      onTodoAdded();
    } catch (error) {
      console.error('Failed to create todo:', error);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="mb-6">
      {!isOpen ? (
        <button
          onClick={() => setIsOpen(true)}
          className="w-full flex items-center justify-center gap-2 py-3 bg-gradient-to-r from-blue-500 to-purple-600 text-white rounded-lg hover:from-blue-600 hover:to-purple-700 transition-all duration-200 shadow-md hover:shadow-lg"
        >
          <Plus size={20} />
          Add New Task
        </button>
      ) : (
        <form onSubmit={handleSubmit} className="bg-white rounded-lg shadow-md p-6">
          <h3 className="text-lg font-semibold text-blue-800 mb-4">Create New Task</h3>
          <div className="space-y-4">
            <div>
              <label htmlFor="title" className="block text-sm font-medium text-blue-800 mb-1">
                Title *
              </label>
              <input
                id="title"
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Enter task title..."
                className="w-full px-4 py-2 border border-blue-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-blue-900 placeholder-blue-400"
                autoFocus
              />
            </div>
            <div>
              <label htmlFor="description" className="block text-sm font-medium text-blue-800 mb-1">
                Description
              </label>
              <textarea
                id="description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder="Enter task description..."
                className="w-full px-4 py-2 border border-blue-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none text-blue-900 placeholder-blue-400"
                rows={3}
              />
            </div>
            <div className="flex gap-3">
              <button
                type="submit"
                disabled={!title.trim() || isSubmitting}
                className="flex-1 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors"
              >
                {isSubmitting ? 'Creating...' : 'Create Task'}
              </button>
              <button
                type="button"
                onClick={() => {
                  setIsOpen(false);
                  setTitle('');
                  setDescription('');
                }}
                className="flex-1 py-2 bg-blue-200 text-blue-700 rounded-md hover:bg-blue-300 transition-colors"
              >
                Cancel
              </button>
            </div>
          </div>
        </form>
      )}
    </div>
  );
}