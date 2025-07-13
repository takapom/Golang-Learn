'use client';

import { useState } from 'react';
import { Todo, todoApi } from '@/lib/api';
import { Check, Edit2, Trash2, X, Save } from 'lucide-react';

interface TodoItemProps {
  todo: Todo;
  onUpdate: () => void;
  onDelete: () => void;
}

export default function TodoItem({ todo, onUpdate, onDelete }: TodoItemProps) {
  const [isEditing, setIsEditing] = useState(false);
  const [title, setTitle] = useState(todo.title);
  const [description, setDescription] = useState(todo.description);

  const handleUpdate = async () => {
    try {
      await todoApi.updateTodo(todo.id, {
        title,
        description,
      });
      setIsEditing(false);
      onUpdate();
    } catch (error) {
      console.error('Failed to update todo:', error);
    }
  };

  const handleToggleComplete = async () => {
    try {
      await todoApi.updateTodo(todo.id, {
        completed: !todo.completed,
      });
      onUpdate();
    } catch (error) {
      console.error('Failed to toggle todo:', error);
    }
  };

  const handleDelete = async () => {
    try {
      await todoApi.deleteTodo(todo.id);
      onDelete();
    } catch (error) {
      console.error('Failed to delete todo:', error);
    }
  };

  return (
    <div className={`bg-white rounded-lg shadow-md p-4 mb-3 transition-all duration-200 hover:shadow-lg ${todo.completed ? 'opacity-75' : ''}`}>
      {isEditing ? (
        <div className="space-y-2">
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="w-full px-3 py-2 border border-blue-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-blue-900"
          />
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="w-full px-3 py-2 border border-blue-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none text-blue-900"
            rows={2}
          />
          <div className="flex gap-2">
            <button
              onClick={handleUpdate}
              className="flex items-center gap-1 px-3 py-1 bg-green-500 text-white rounded-md hover:bg-green-600 transition-colors"
            >
              <Save size={16} />
              Save
            </button>
            <button
              onClick={() => {
                setIsEditing(false);
                setTitle(todo.title);
                setDescription(todo.description);
              }}
              className="flex items-center gap-1 px-3 py-1 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition-colors"
            >
              <X size={16} />
              Cancel
            </button>
          </div>
        </div>
      ) : (
        <div className="flex items-center justify-between">
          <div className="flex items-start gap-3 flex-1">
            <button
              onClick={handleToggleComplete}
              className={`mt-1 w-5 h-5 rounded border-2 flex items-center justify-center transition-all ${
                todo.completed
                  ? 'bg-blue-500 border-blue-500'
                  : 'border-blue-300 hover:border-blue-500'
              }`}
            >
              {todo.Completed && <Check size={14} className="text-white" />}
            </button>
            <div className="flex-1">
              <h3 className={`font-semibold text-lg ${todo.Completed ? 'line-through text-blue-400' : 'text-blue-900'}`}>
                {todo.title}
              </h3>
              {todo.description && (
                <p className={`mt-1 text-sm ${todo.Completed ? 'text-blue-400' : 'text-blue-700'}`}>
                  {todo.description}
                </p>
              )}
            </div>
          </div>
          <div className="flex items-center gap-2">
            <button
              onClick={() => setIsEditing(true)}
              className="p-2 text-blue-500 hover:bg-blue-50 rounded-md transition-colors"
            >
              <Edit2 size={18} />
            </button>
            <button
              onClick={handleDelete}
              className="p-2 text-red-500 hover:bg-red-50 rounded-md transition-colors"
            >
              <Trash2 size={18} />
            </button>
          </div>
        </div>
      )}
    </div>
  );
}